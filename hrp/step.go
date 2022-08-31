package hrp

type StepType string

const (
	stepTypeRequest     StepType = "request"
	stepTypeAPI         StepType = "api"
	stepTypeTestCase    StepType = "testcase"
	stepTypeTransaction StepType = "transaction"
	stepTypeRendezvous  StepType = "rendezvous"
	stepTypeThinkTime   StepType = "thinktime"
	stepTypeWebSocket   StepType = "websocket"
)

type StepResult struct {
	Name              string                 `json:"name" yaml:"name"`                                                   // step name
	StepType          StepType               `json:"step_type" yaml:"step_type"`                                         // step type, testcase/request/transaction/rendezvous
	Success           bool                   `json:"success" yaml:"success"`                                             // step execution result
	Elapsed           int64                  `json:"elapsed_ms" yaml:"elapsed_ms"`                                       // step execution time in millisecond(ms)
	HttpStat          map[string]int64       `json:"httpstat,omitempty" yaml:"httpstat,omitempty"`                       // httpstat in millisecond(ms)
	DisableReportBody bool                   `json:"disable_report_body,omitempty" yaml:"disable_report_body,omitempty"` // 是否在html中打印body
	Data              interface{}            `json:"data,omitempty" yaml:"data,omitempty"`                               // session data or slice of step data
	ContentSize       int64                  `json:"content_size" yaml:"content_size"`                                   // response body length
	ExportVars        map[string]interface{} `json:"export_vars,omitempty" yaml:"export_vars,omitempty"`                 // extract variables
	Attachment        string                 `json:"attachment,omitempty" yaml:"attachment,omitempty"`                   // step error information
}

// TStep represents teststep data structure.
// Each step maybe three different types: make one request or reference another api/testcase.
type TStep struct {
	Name              string                 `json:"name" yaml:"name"` // required
	Request           *Request               `json:"request,omitempty" yaml:"request,omitempty"`
	API               interface{}            `json:"api,omitempty" yaml:"api,omitempty"`           // *APIPath or *API api和request有啥区别，api主要用在什么场景当中？
	TestCase          interface{}            `json:"testcase,omitempty" yaml:"testcase,omitempty"` // *TestCasePath or *TestCase 这个应该是用于跑测试用例的，有点类似于golang的单元测试
	Transaction       *Transaction           `json:"transaction,omitempty" yaml:"transaction,omitempty"`
	Rendezvous        *Rendezvous            `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty"`
	ThinkTime         *ThinkTime             `json:"think_time,omitempty" yaml:"think_time,omitempty"` // 模拟人的思考时间，实际上就是加延时
	WebSocket         *WebSocketAction       `json:"websocket,omitempty" yaml:"websocket,omitempty"`
	Variables         map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`                     // 变量设置
	DisableReportBody bool                   `json:"disable_report_body,omitempty" yaml:"disable_report_body,omitempty"` // 是否在html中打印body
	SetupHooks        []string               `json:"setup_hooks,omitempty" yaml:"setup_hooks,omitempty"`
	TeardownHooks     []string               `json:"teardown_hooks,omitempty" yaml:"teardown_hooks,omitempty"`
	Extract           map[string]string      `json:"extract,omitempty" yaml:"extract,omitempty"`
	Validators        []interface{}          `json:"validate,omitempty" yaml:"validate,omitempty"`
	Export            []string               `json:"export,omitempty" yaml:"export,omitempty"`
}

// IStep represents interface for all types for teststeps, includes:
// StepRequest, StepRequestWithOptionalArgs, StepRequestValidation, StepRequestExtraction,
// StepTestCaseWithOptionalArgs,
// StepTransaction, StepRendezvous, StepWebSocket.
type IStep interface {
	Name() string
	Type() StepType
	Struct() *TStep
	Run(*SessionRunner) (*StepResult, error)
}
