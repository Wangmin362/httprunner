package main

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"gitcdteam.skyguardmis.com/bigdt/gokit/pkg/cryptox"
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
	token := cryptox.Sha256Str(xTimestamp + popCode + tenantId + popId)
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
	token := cryptox.Sha256Str(xTimestamp + popCode + popId)
	basicAuthStr := strings.Join([]string{xTimestamp, token, popId}, ":")
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuthStr))
	return authorization
}
func GetTenantId() string {
	return "1000001"
}
