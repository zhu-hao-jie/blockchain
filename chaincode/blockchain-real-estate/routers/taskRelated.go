package routers

//task 是由 submitter 发布的任务（需要完成的订单任务）
import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"strconv"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	// "chaincode/blockchain-real-estate/match"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/match"
	"strings"
)

func CreateOrderStr(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("参数个数有缺")
	}
	submitter := args[0]
	task := args[1]
	//receiver := args[2]
	status := args[2]
	time := args[3]
	//owner := args[4]   //创建初订单拥有者为空: ""

	//根据 owner 获取制造商信息
	//resultsAccount,err := utils.GetStateByPartialCompositeKeys(stub,lib.PlusAccountKey,[]string)

	orderStr := &lib.OrderStr{
		Submitter: submitter,
		Task:      task,
		Status:    status,
		Time:      time,
		//Owner    : owner,
	}
	//写入账本
	if err := utils.WriteLedger(orderStr, stub, lib.OrderStrKey, []string{orderStr.Submitter, orderStr.Task}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	orderStrByte, err := json.Marshal(orderStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错：%s", err))
	}
	//成功返回
	return shim.Success(orderStrByte)
}

// func CreateOrderStrByReceiver(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	验证参数
// 	if len(args)!= 3 {
// 		return shim.Error("残花个数不满足")
// 	}
// 	task := args[0]
// 	submitter := arg[1]
// 	owner := args[2]

// 	if task =="" || submitter == "" || owner ==""{
// 		return shim.Error("参数存在空值")
// 	}
// 	//根据 submitter 和 task 获取订单信息
// 	resultOrderStr,err := utils.GetStateByPartialCompositeKeys2(stub,lib.OrderStrKey,[]string{submitter,task})
// 	if err != nil || len(resultOrderStr)!= 1 {
// 		return shim.Error(fmt.Sprintf("%s 和 %s获取订单信息失败：%s",err))
// 	}
// 	var orderStr lib.OrderStr
// 	if err = json.Unmarshal(resultOrderStr[0],&orderStr);err != nil {
// 		return shim.Error(fmt.Sprintf("CreateOrderStrByReceiver-反序列化失败：%s",err))
// 	}
// 	获取制造商信息
// 	resultAccount,err := utils.GetStateByPartialCompositeKeys(stub,lib.PlusAccountKey,[]string{owner})
// 	if err != nil ||len(resultAccount)!= 1{
// 		return shim.Error(fmt.Sprintf("制造商信息验证失败"%s,err))
// 	}
// 	var ownerAccount lib.PlusAccountKey
// 	if err = json.Unmarshal(resultAccount[0],&ownerAccount);err != nil{
// 		return shim.Error(fmt.Sprintf("查询制造商信息-反序列化出错%s",err))
// 	}
// /***********************出错处***************************************/
// 	//判断算法
// 	resource:= match.Match(orderStr.Task,ownerAccount.processName)

// 	//取得 resource 中的 key 值
// 	func getKeys(m map[string][]string) []string {
// 		// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
// 		keys := make([]string, 0, len(m))
// 		for k := range m {
// 			keys = append(keys, k)
// 		}
// 		return keys
// 	}
// 	orderStr.Owner = keys
// 	if err := utils.WriteLedger(orderStr, stub, lib.OrderStrKey, []string{OrderStr.Submitter, OrderStr.Task}); err != nil {
// 		return shim.Error(fmt.Sprintf("将owner写入订单OrderStr,修改交易状态 失败%s", err))
// 	}
// 	//将本次购买交易写入账本,可供买家查询
// 	orderStrByRe := &lib.OrderStrByRe{
// 		Owner:      owner,
// 		CreateTime: time.Now().Local().Format("2006-01-02 15:04:05"),
// 		OrderStr:   orderStr,
// 	}
// 	local, _ := time.LoadLocation("Local")
// 	createTimeUnixNano, _ := time.ParseInLocation("2006-01-02 15:04:05", sellingBuy.CreateTime, local)
// 	if err := utils.WriteLedger(orderStrByRe, stub, lib.OrderStrByReKey, []string{orderStrByRe.Owner, fmt.Sprintf("%d", createTimeUnixNano.UnixNano())}); err != nil {
// 		return shim.Error(fmt.Sprintf("将本次购买交易写入账本失败%s", err))
// 	}
// 	orderStrByReByte, err := json.Marshal(orderStrByRe)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
// 	}
// 	return shim.Success(orderStrByReByte)
// }

//调用匹配算法
func DoMatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("参数个数不满足")
	}
	task := args[0]
	res := args[1] //是字符串

	var tempR []string
	var task1 []string
	var r []string
	for i := 0; i < len(res); i++ {
		tempR = strings.Split(res[0:], ",")
	}
	r = tempR[1:]
	for i := 0; i < len(task); i++ {
		task1 = strings.Split(task, ",")
	}

	resource := make(map[string][]string)
	resource[tempR[0]] = r

	// var key string
	// var value []string

	// doMatch := &lib.DoMatch{
	// 	Owner :  owner,
	// 	OwnerRec  :  ownerRec,
	// }

	var doMatch lib.DoMatch
	retResource := match.Match(task1, resource)
	for k, v := range retResource {
		doMatch.Owner = k
		doMatch.OwnerRec = v
	}
	doMatch.Task = task
	doMatch.Resource = resource
	//OwnerReceiver := strings.Join(doMatch.OwnerRec,",")

	//写入账本
	if err := utils.WriteLedger(doMatch, stub, lib.DoMatchKey, []string{doMatch.Owner, doMatch.Task}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回,此处更改了序列化
	doMatchByte, err := json.Marshal(doMatch)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错：%s", err))

	}
	//成功返回
	return shim.Success(doMatchByte)

}

//输入参数为任务，输入后执行与资源匹配
// func TaskMatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	// if len(args) != 1 {
// 	// 	return shim.Error("参数个数不满足")
// 	// }
// 	task := args[0]
// 	var tasks []string
// 	for i := 0; i < len(task); i++ {
// 		tasks = strings.Split(task, ",") //将字符串转为数组
// 	}
// 	/*********** 取得 receiver 中的 Res************************/
// 	/****************此处数据没有读进来*********************/
// 	// var receiverRes []lib.Receiver
// 	// result,err := utils.GetStateByPartialCompositeKeys2(stub,lib.ReceiverKey,args)
// 	// if err != nil{
// 	// 	return shim.Error(fmt.Sprintf("%s",err))
// 	// }

// 	// for _,v := range result {
// 	// 	if v != nil{
// 	// 		var recRes lib.Receiver
// 	// 		err := json.Unmarshal(v,&recRes)
// 	// 		if err != nil {
// 	// 			return shim.Error(fmt.Sprintf("QueryReceiverRes-反序列化出错: %s", err))
// 	// 		}
// 	// 		receiverRes = append(receiverRes,recRes)
// 	// 	}
// 	// }

// 	receiverRes := []lib.Receiver{
// 		{"E1", []string{"Dr", "Bt"}},
// 		{"E2", []string{"Tr,Mo,Yo"}},
// 		{"E3", []string{"Dr,Bt", "Mo"}},
// 	}
// 	resource := make(map[string][]string)

// 	for i := 0; i < len(receiverRes); i++ {
// 		resource[receiverRes[i].ResName] = receiverRes[i].Process
// 	}

// 	var taskMatch []lib.DoMatch
// 	returnResource := match.Match(tasks, resource)
// 	i := 0
// 	for k, v := range returnResource {
// 		taskMatch[i].Owner = k
// 		taskMatch[i].OwnerRec = v
// 		taskMatch[i].Task = task
// 		taskMatch[i].Resource = returnResource
// 		i++
// 	}

// 	//写入账本
// 	if err := utils.WriteLedger(taskMatch, stub, lib.TaskMatchKey, []string{taskMatch.Owner, taskMatch.Task}); err != nil {
// 		return shim.Error(fmt.Sprintf("%s", err))
// 	}
// 	//将成功创建的信息返回,此处更改了序列化
// 	taskMatchByte, err := json.Marshal(taskMatch)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错：%s", err))

// 	}
// 	//成功返回
// 	return shim.Success(taskMatchByte)

// }

// /*************** 查询单一输入 **************************************/
// func QueryTaskMatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var taskMatchList []lib.DoMatch
// 	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.TaskMatchKey, args)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("%s", err))
// 	}
// 	for _, v := range result {
// 		if v != nil {
// 			var match lib.DoMatch
// 			err := json.Unmarshal(v, &match)
// 			if err != nil {
// 				return shim.Error(fmt.Sprintf("QueryTaskMatch-反序列化出错：%s", err))
// 			}
// 			taskMatchList = append(taskMatchList, match)
// 		}
// 	}
// 	taskMatchListByte, err := json.Marshal(taskMatchList)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("QueryTaskMatch-序列化出错：%s,err"))
// 	}
// 	return shim.Success(taskMatchListByte)
// }

/*****************************************************/
func QueryDoMatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var doMatchList []lib.DoMatch
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.DoMatchKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range result {
		if v != nil {
			var match lib.DoMatch
			err := json.Unmarshal(v, &match)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDoMatch-反序列化出错：%s", err))
			}
			doMatchList = append(doMatchList, match)
		}
	}
	doMatchListByte, err := json.Marshal(doMatchList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDoMatch-序列化出错：%s,err"))
	}
	return shim.Success(doMatchListByte)
}

// type Resou struct {
// 	Rprocess []string
// }
func QueryOrderStr(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//var arr []string
	//var doMatch lib.DoMatch
	var orderStrList []lib.OrderStr
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderStrKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range result {
		if v != nil {
			var orderStr lib.OrderStr
			/********         如何将 OWNER 加入到 订单状态中           *********************/
			// orderStr.Owner = "tURN"//doMatch.Owner
			// //写入账本
			// if err := utils.WriteLedger(orderStr,stub,lib.OrderStrKey,[]string{orderStr.Owner});err != nil {
			// 	return shim.Error(fmt.Sprintf("%s",err))
			// }

			err := json.Unmarshal(v, &orderStr)
			if err != nil {
				return shim.Error(fmt.Sprintf("查询订单反序列化错误：%s", err))
			}
			orderStrList = append(orderStrList, orderStr)
		}
	}

	orderStrListByte, err := json.Marshal(orderStrList)
	if err != nil {
		return shim.Error(fmt.Sprintf("查询订单序列化出错：%s", err))
	}
	return shim.Success(orderStrListByte)
}

//更新参数为 状态值信息
func UpdateOrderStr(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//验证参数
	if len(args) != 4 {
		return shim.Error("参数个数不满足")
	}
	submitter := args[0]
	task := args[1]

	statusMap := lib.OrderStrStatus()
	status := statusMap[args[2]]
	time := args[3]

	if submitter == "" || task == "" || status == "" || time == "" {
		return shim.Error("参数存在空值")
	}

	//判断订单数据是否存在
	orderStrCheck, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderStrKey, []string{submitter, task})
	if err != nil /*|| len(orderStrCheck) != 1*/ {
		return shim.Error(fmt.Sprintf("task 不存在，%s", err))
	}
	//获取目标订单
	var targetOrderStr lib.OrderStr
	err = json.Unmarshal(orderStrCheck[0], &targetOrderStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("反序列化出错，%s", err))
	}

	//更新订单信息
	targetOrderStr.Status = status
	//写入账本
	if err := utils.WriteLedger(targetOrderStr, stub, lib.OrderStrKey, []string{targetOrderStr.Submitter, targetOrderStr.Task}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	orderStrByte, err := json.Marshal(targetOrderStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错：%s", err))
	}
	//成功返回
	return shim.Success(orderStrByte)
}
