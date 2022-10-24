package routers

//制造企业，有哪些加工能力，能加工哪些任务
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"

	//"strconv"
	//"time"
	"strings"
)

// func ArrayToString(arr []string) string{
// 	var result string
// 	for _,i := range arr {
// 		result += i
// 	}
// 	return result
// }

func ArrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result += i
	}
	return result
}

func CreateReceiver(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("参数个数不满足")
	}

	resName := args[0]

	process := args[1]

	//var r lib.Res
	var r []string
	for i := 0; i < len(process); i++ {
		//r.Num = int(res[0])-48
		r = strings.Split(process[0:], ",") //将字符串转换成字符串数组
	}

	if resName == "" || process == "" {
		return shim.Error("参数存在空值")
	}

	receiver := &lib.Receiver{
		ResName: resName,
		Process: r,
	}

	//process := ArrayToString(receiver.Res.Process)
	// 写入账本
	if err := utils.WriteLedger(receiver, stub, lib.ReceiverKey, []string{receiver.ResName}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	receiverByte, err := json.Marshal(receiver)
	if err != nil {
		return shim.Error(fmt.Sprintf("receiver序列化成功创建的信息出错: %s", err))
	}

	//fmt.Printf("name is : %v; \n 制造服务是 is : %v ", resName, r.Process)
	// 成功返回
	return shim.Success(receiverByte)

}

//查询制造商能够提供的制造能力,由 ResName 搜寻
func QueryReceiverRes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var receiverRes []lib.Receiver
	result, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ReceiverKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range result {
		if v != nil {
			var recRes lib.Receiver
			err := json.Unmarshal(v, &recRes)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryReceiverRes-反序列化出错: %s", err))
			}
			receiverRes = append(receiverRes, recRes)
		}
	}

	receiverResByte, err := json.Marshal(receiverRes)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryReceiverRes-序列化出错: %s", err))
	}
	return shim.Success(receiverResByte)

}
