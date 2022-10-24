package lib

type OrderStr struct {
	Submitter string `json:submitter`
	Task      string `json:task` //任务名称
	//Receiver  string `json:receiver`
	Status string `json:status`
	Time   string `json:time`
	Owner  string `json:owner`
}

// type Res struct{
// 	//Num int //用作计数
// 	Process []string  //加工能力
// }

type Receiver struct {
	ResName string
	//Res Res
	Process []string
}

type DoMatch struct {
	Task     string              `json:task`     //需求方发布的任务
	Resource map[string][]string `json:resource` //供应方提供的制造资源
	Owner    string              `json:owner`
	OwnerRec []string            `json:ownerRec`
}

type OrderStrByRe struct {
	Owner      string   `json:"owner"`
	CreateTime string   `json:"createTime"`
	OrderStr   OrderStr `json:"orderStr"`
}

const (
	OrderStrKey     = "orderstr-key"
	ReceiverKey     = "receiver-key"
	PlusAccountKey  = "plusAccount-key"
	OrderStrByReKey = "orderStr-byRe-Key"
	DoMatchKey      = "doMatch-key"
	TaskMatchKey    = "taskMatch-key"
)

var OrderStrStatus = func() map[string]string {
	return map[string]string{
		"noStart": "未开始",
		"done":    "已完成",
	}
}

type PlusAccount struct {
	AccountId string `json:"accountId"` //账号Name
	UserName  string `json:"userName"`  //账号名
	//Balance   float64 `json:"balance"`   //余额
	ProcessName []string `json:processName`
}