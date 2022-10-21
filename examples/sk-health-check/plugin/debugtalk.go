//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetCurrentTimestamp() string {
	xTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return xTimestamp
}

// 3.1 pop集群、租户容器调用ucss接口添加header参数
// x-timestamp: xxxx
// x-tenant-id: 1xxxx
// Authorization: Basic base64.encode(timestamp:token:tenantid:popid)
// 其中，token=SHA-256(timestamp+popCode+tenantId+popId)；
// popid和popCode在

func GetAuth(tenantId, popCode, popId, xTimestamp string) (authorization string) {
	pCode, _ := base64.StdEncoding.DecodeString(popCode)
	pCode = pCode[4:28]
	sha256Bytes := sha256.Sum256([]byte(xTimestamp + string(pCode) + tenantId + popId))
	token := hex.EncodeToString(sha256Bytes[:])
	basicAuthStr := strings.Join([]string{xTimestamp, token, tenantId, popId}, ":")
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuthStr))
	return authorization
}

// 3.2 pop集群接口添加header参数
// x-timestamp: xxxx
// Authorization: Basic base64.encode(timestamp:token:popid)
// 其中，token=SHA-256(timestamp+popCode+popId)；
// popid和popCode在

func GetPopAuth(popCode, popId, xTimestamp string) (authorization string) {
	sha256Bytes := sha256.Sum256([]byte(xTimestamp + popCode + popId))
	token := hex.EncodeToString(sha256Bytes[:])
	basicAuthStr := strings.Join([]string{xTimestamp, token, popId}, ":")
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuthStr))
	return authorization
}

func GetUcwiAuth(accessKey, secretKey, timestamp string) string {
	tokenSource := secretKey + timestamp
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(tokenSource))
	tokenResult := hex.EncodeToString(hash.Sum(nil))

	auth := fmt.Sprintf("SKG %s:%s", accessKey, tokenResult)
	return auth
}

func GetUUID() string {
	return uuid.New().String()
}

func GetUcwiDlpChannelMetadata(uuid string) string {
	metadatas := map[string]interface{}{
		"user":     "cloudtest\\clouduser",
		"filename": "file1",
		"queryID":  uuid,
		"redaction": map[string]string{
			"sendBackType": "response",
		},
	}

	metadataStr, _ := json.Marshal(metadatas)
	return string(metadataStr)
}

func GetRandomFile() []interface{} {
	return []interface{}{"examples/sk-health-check/testdata/file1.txt", "examples/sk-health-check/testdata/file2.txt", "examples/sk-health-check/testdata/file3.txt"}
}

func GetRandomData(num, length string) string {
	n, _ := strconv.ParseInt(num, 10, 64)
	l, _ := strconv.ParseInt(length, 10, 64)
	data := make([]string, n)
	for i := int64(0); i < n; i++ {
		d := make([]rune, l)
		for j := int64(0); j < l; j++ {
			d[j] = rune(RandInt())
		}
		data[i] = string(d)
	}

	randomData := "[" + strings.Join(data, ",") + "]"
	return randomData
}

func RandInt() int64 {
	rand.Seed(time.Now().UnixNano())
	return 19968 + rand.Int63n(40869-19968)
}
