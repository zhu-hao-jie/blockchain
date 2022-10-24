package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"testing"
	//"strings"
)

func initTest(t *testing.T) *shim.MockStub {
	scc := new(BlockChainServices)
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func byteToString(input [][]byte) []string {
	var output []string
	for _, v := range input {
		output = append(output, string(v))
	}
	return output
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) pb.Response {
	res := stub.MockInvoke("1", args)
	argsString := byteToString(args)
	if res.Status != shim.OK {
		fmt.Println("Error : Invoke Failed ", "agrgs : ", argsString, string(res.Message))
		// fmt.Println("Invoke", argsString, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

// 测试链码初始化
func TestBlockChainServices_Init(t *testing.T) {
	initTest(t)
}

//测试获取账户信息
func Test_QueryAccountList(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取多个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("5feceb66ffc8"),
			[]byte("6b86b273ff34"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取单个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("4e07408562be"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("4、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("0"),
		}).Payload)))
}

// 测试捐赠合约
// func Test_Donating(t *testing.T) {
// 	stub := initTest(t)
// 	servicesList := checkCreateServices(stub, t)

// fmt.Println(fmt.Sprintf("获取制造服务信息\n%s",
// 	string(checkInvoke(t, stub, [][]byte{
// 		[]byte("queryServicesList"),
// 	}).Payload)))
// 	//先发起
// 	fmt.Println(fmt.Sprintf("发起捐赠\n%s", string(checkInvoke(t, stub, [][]byte{
// 		[]byte("createDonating"),
// 		[]byte(servicesList[0].ServicesName),
// 		[]byte(servicesList[0].Proprietor),
// 		[]byte(servicesList[2].Proprietor),
// 	}).Payload)))

// fmt.Println(fmt.Sprintf("获取制造服务信息\n%s",
// 	string(checkInvoke(t, stub, [][]byte{
// 		[]byte("queryServicesList"),
// 	}).Payload)))

// 	fmt.Println(fmt.Sprintf("1、查询所有\n%s", string(checkInvoke(t, stub, [][]byte{
// 		[]byte("queryDonatingList"),
// 	}).Payload)))
// 	fmt.Println(fmt.Sprintf("2、查询指定%s\n%s", servicesList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
// 		[]byte("queryDonatingList"),
// 		[]byte(servicesList[2].Proprietor),
// 	}).Payload)))
// 	fmt.Println(fmt.Sprintf("3、查询指定受赠%s\n%s", servicesList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
// 		[]byte("queryDonatingListByGrantee"),
// 		[]byte(servicesList[2].Proprietor),
// 	}).Payload)))

// 	//fmt.Println(fmt.Sprintf("4、接受受赠%s\n%s", servicesList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
// 	//	[]byte("updateDonating"),
// 	//	[]byte(servicesList[0].ServicesName),
// 	//	[]byte(servicesList[0].Proprietor),
// 	//	[]byte(servicesList[2].Proprietor),
// 	//	[]byte("done"),
// 	//}).Payload)))
// 	fmt.Println(fmt.Sprintf("4、取消受赠%s\n%s", servicesList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
// 		[]byte("updateDonating"),
// 		[]byte(servicesList[0].ServicesName),
// 		[]byte(servicesList[0].Proprietor),
// 		[]byte(servicesList[2].Proprietor),
// 		[]byte("cancelled"),
// 	}).Payload)))

// 	fmt.Println(fmt.Sprintf("获取制造服务信息\n%s",
// 		string(checkInvoke(t, stub, [][]byte{
// 			[]byte("queryServicesList"),
// 		}).Payload)))
// }

//手动创建一些订单
func checkOrders(t *testing.T, stub *shim.MockStub) {

	fmt.Println("0. 创建部分新订单")
	fmt.Println("============================")
	checkInvoke(t, stub, [][]byte{
		[]byte("createOrderStr"),
		[]byte("Enterprise1"), //任务发布者
		[]byte("CheXiao"),     //任务名称
		[]byte("done"),        //状态
		[]byte("20220323"),    //时间
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createOrderStr"),
		[]byte("Enterprise2"), //操作人
		[]byte("Mo"),          //owner
		[]byte("noStart"),     //orderId
		[]byte("20001202"),    //status
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createOrderStr"),
		[]byte("Enterprise3"), //操作人
		[]byte("GuangKe"),     //owner
		[]byte("done"),        //orderId
		[]byte("20221230"),    //status
	})
}

//zzy 测试创建新的订单
func Test_OrderCURD(t *testing.T) {
	stub := initTest(t)
	//创建一些订单
	checkOrders(t, stub)

	//查询所有订单
	// fmt.Println(fmt.Sprintf("1、测试获取所有订单数据\n%s",
	// 	string(checkInvoke(t, stub, [][]byte{
	// 		[]byte("queryOrderList"),
	// 	}).Payload)))
	fmt.Println(fmt.Sprintf("1.测试获取所有订单数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryOrderStr"),
		}).Payload)))

	//查询用户1拥有的订单
	// fmt.Println(fmt.Sprintf("2、测试获取指定用户所有订单数据\n%s",
	// 	string(checkInvoke(t, stub, [][]byte{
	// 		[]byte("queryOrderList"),
	// 		[]byte("6b86b273ff34"),
	// 	}).Payload)))

	//查询指定订单Name
	// fmt.Println(fmt.Sprintf("3、测试获取指定订单数据\n%s",
	// 	string(checkInvoke(t, stub, [][]byte{
	// 		[]byte("queryOrderList"),
	// 		[]byte("6b86b273ff34"),
	// 		[]byte("nuaaLab41602"),
	// 	}).Payload)))
	fmt.Println(fmt.Sprintf("2.测试获取指定订单数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryOrderStr"),
			[]byte("Enterprise2"),
			[]byte("Mo"),
		}).Payload)))

	// fmt.Println(fmt.Sprintf("4、测试更新指定订单数据\n%s",
	// 	string(checkInvoke(t, stub, [][]byte{
	// 		[]byte("updateOrder"),
	// 		[]byte("6b86b273ff34"), //操作人
	// 		[]byte("6b86b273ff34"), //owner
	// 		[]byte("nuaaLab41601"), //orderId
	// 		[]byte("inProgress"),   //status
	// 	}).Payload)))

	// fmt.Println("5、测试获取指定订单历史数据")
	// fmt.Printf("%s\n", string(checkInvoke(t, stub, [][]byte{
	// 	[]byte("queryOrderHistory"),
	// 	[]byte("5feceb66ffc8"), //操作人
	// 	[]byte("6b86b273ff34"), //owner
	// 	[]byte("nuaaLab41601"), //orderId
	// }).Payload))

	// fmt.Println("5、测试更新指定订单数据")
	// fmt.Printf("%s\n", string(checkInvoke(t, stub, [][]byte{
	// 	[]byte("updateOrder"),
	// 	[]byte("5feceb66ffc8"), //操作人
	// 	[]byte("6b86b273ff34"), //owner
	// 	[]byte("nuaaLab41601"), //orderId
	// 	[]byte("inProgress"),   //status
	// }).Payload))
	fmt.Println("3.测试更新指定订单数据")
	fmt.Printf("%s\n", string(checkInvoke(t, stub, [][]byte{
		[]byte("updateOrderStr"),
		[]byte("Enterprise2"),
		[]byte("Mo"),
		[]byte("done"),
		[]byte("20001202"),
	}).Payload))

	//创建已有订单
	// fmt.Println("6、测试创建已有订单数据")
	// fmt.Printf("%s\n", string(checkInvoke(t, stub, [][]byte{
	// 	[]byte("createOrder"),
	// 	[]byte("5feceb66ffc8"), //操作人
	// 	[]byte("d4735e3a265e"), //owner
	// 	[]byte("nuaaLab41603"), //orderId
	// 	[]byte("done"),         //status
	// }).Payload))

	fmt.Println("4、测试创建已有订单数据")
	fmt.Printf("%s\n", string(checkInvoke(t, stub, [][]byte{
		[]byte("createOrderStr"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("d4735e3a265e"), //owner
		[]byte("nuaaLab41603"), //orderId
		[]byte("done"),         //status
	}).Payload))

}

// type Res struct {
// 	//num     int //用作计数
// 	process []string
// }

// func Match(arr []string, m map[string]Res) { //arr 是 task , m 是 resource
// 	//key := &
// 	for key, val := range m { //for i := 0; i < len(arr); i++ { //for i, v := range arr {
// 		j := 0
// 		num:=0
// 		t := make([]string, len(m))

// 		for i := 0; i < len(arr); i++ {
// 			//v := &val
// 			for n := 0; n < len(val.process); n++ {
// 				//strings.Compare(val.process[n], arr[i])
// 				if val.process[n] == arr[i] {
// 					num++
// 					t[j] = val.process[n]
// 					j++

// 				}
// 			}
// 		}
// 		if len(arr) == num {
// 			fmt.Printf("%v 企业能接受 task: %v\n", key, arr)
// 		} else if num > 0 && num < len(arr) {
// 			fmt.Printf("%v 企业部分接受 task:%v\n", key, t)
// 		} else {
// 			fmt.Printf("%v 企业不能加工 task\n ", key)
// 		}
// 	}
// }
// func Test_DoMatch(t *testing.T){
// 	res1 := Res{ []string{"De", "Dr", "Bo", "Mi", "Sa"}}
// 	res2 := Res{ []string{"Sa", "Bo", "W"}}
// 	res3 := Res{ []string{"Bo", "Mi"}}
// 	res4 := Res{ []string{"De", "Sa", "Dr"}}
// 	res5 := Res{ []string{"Tu"}}

// 	resource := make(map[string]Res)

// 	resource["E1"] = res1
// 	resource["E2"] = res2
// 	resource["E3"] = res3
// 	resource["E4"] = res4
// 	resource["E5"] = res5
// 	var task1 = []string{"Bo", "Dr","Mi","Sa"} //task 由任务发布者发布
// 	// var task2 = []string{"De", "M", "Dr"}
// 	// var task3 = []string{"De", "Bo", "M", "W"}
// 	fmt.Println("24.测试匹配过程")
// 	Match(task1, resource)
// }

func Test_TaskMatch(t *testing.T) {
	stub := initTest(t)
	checkInvoke(t, stub, [][]byte{
		[]byte("taskMatch"),
		[]byte("De,Bt"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("taskMatch"),
		[]byte("Tr,Bt,Mt"),
	})
	fmt.Println("19.测试只传入任务匹配算法")
}

func Test_DoMatch(t *testing.T) {
	stub := initTest(t)
	checkInvoke(t, stub, [][]byte{
		[]byte("doMatch"),
		[]byte("De,Bt"),
		[]byte("Drilling,De,Fr,Bt"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("doMatch"),
		[]byte("Tr,Bt,Mt"),
		[]byte("Turning,De,Fr,Bt"),
	})
	fmt.Println("20.测试匹配算法")
}

// func checkDoMatch(stub *shim.MockStub,t *testing.T)[]lib.DoMatch{
// 	var testMatch []lib.DoMatch
// 	var tMatch lib.DoMatch

// 	r1 := checkInvoke(t,stub,[][]byte{
// 		[]byte("doMatch"),
// 		[]byte("De,Bt"),
// 		[]byte("Turn,De,Fr,Bt"),
// 	})
// 	r2 := checkInvoke(t,stub,[][]byte{
// 		[]byte("doMatch"),
// 		[]byte("De,Ty"),
// 		[]byte("Saw,De,Fr,Bt"),
// 	})
// 	json.Unmarshal(bytes.NewBuffer(r1.Payload).Bytes(), &tMatch)
// 	testMatch = append(testMatch, tMatch)
// 	json.Unmarshal(bytes.NewBuffer(r2.Payload).Bytes(), &tMatch)
// 	testMatch = append(testMatch, tMatch)
// 	return testMatch
// }
func checkCreateReceiver(stub *shim.MockStub, t *testing.T) []lib.Receiver {
	var receiverRes []lib.Receiver
	var recRes lib.Receiver

	r1 := checkInvoke(t, stub, [][]byte{
		[]byte("createReceiver"),
		[]byte("E2"),
		[]byte("De,Fr,Ty"),
	})
	r2 := checkInvoke(t, stub, [][]byte{
		[]byte("createReceiver"),
		[]byte("E3"),
		[]byte("De,Bt,Mo"),
	})
	r3 := checkInvoke(t, stub, [][]byte{
		[]byte("createReceiver"),
		[]byte("E4"),
		[]byte("M,Ty,Bt"),
	})
	json.Unmarshal(bytes.NewBuffer(r1.Payload).Bytes(), &recRes)
	receiverRes = append(receiverRes, recRes)
	json.Unmarshal(bytes.NewBuffer(r2.Payload).Bytes(), &recRes)
	receiverRes = append(receiverRes, recRes)
	json.Unmarshal(bytes.NewBuffer(r3.Payload).Bytes(), &recRes)
	receiverRes = append(receiverRes, recRes)
	return receiverRes
}

func Test_QueryReceiverRes(t *testing.T) {
	stub := initTest(t)
	receiverRes := checkCreateReceiver(stub, t)

	fmt.Println(fmt.Sprintf("11、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryReceiverRes"),
		}).Payload)))

	fmt.Println(fmt.Sprintf("12、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryReceiverRes"),
			[]byte(receiverRes[0].ResName),
		}).Payload)))

	fmt.Println(fmt.Sprintf("13、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryReceiverRes"),
			[]byte("0"),
		}).Payload)))

}

func checkTaskMatch(stub *shim.MockStub, t *testing.T) []lib.DoMatch {
	var testMatch []lib.DoMatch
	var tMatch lib.DoMatch

	r1 := checkInvoke(t, stub, [][]byte{
		[]byte("taskMatch"),
		[]byte("De,Bt"),
		//[]byte("Turn,De,Fr,Bt"),
	})
	r2 := checkInvoke(t, stub, [][]byte{
		[]byte("taskMatch"),
		[]byte("De,Ty"),
		//[]byte("Saw,De,Fr,Bt"),
	})
	json.Unmarshal(bytes.NewBuffer(r1.Payload).Bytes(), &tMatch)
	testMatch = append(testMatch, tMatch)
	json.Unmarshal(bytes.NewBuffer(r2.Payload).Bytes(), &tMatch)
	testMatch = append(testMatch, tMatch)
	return testMatch
}

func Test_QueryTaskMatch(t *testing.T) {
	stub := initTest(t)
	taskMatch := checkTaskMatch(stub, t)
	fmt.Println(fmt.Sprintf("21.测试匹配\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTaskMatch"),
		}).Payload)))

	fmt.Println(fmt.Sprintf("22.测试获取部分匹配\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTaskMatch"),
			[]byte(taskMatch[0].Owner),
		}).Payload)))

	fmt.Println(fmt.Sprintf("23.测试获取指定匹配\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryTaskMatch"),
			[]byte("E2"),
		}).Payload)))

}

// func Test_QueryDoMatch(t *testing.T){
// 	stub := initTest(t)
// 	doMatch := checkDoMatch(stub,t)
// 	fmt.Println(fmt.Sprintf("21.测试匹配\n%s",
// 		string(checkInvoke(t,stub,[][]byte{
// 			[]byte("queryDoMatch"),
// 		}).Payload)))

// 	fmt.Println(fmt.Sprintf("22.测试获取部分匹配\n%s",
// 		string(checkInvoke(t,stub,[][]byte{
// 			[]byte("queryDoMatch"),
// 			[]byte(doMatch[0].Owner),
// 		}).Payload)))

// 		fmt.Println(fmt.Sprintf("23.测试获取指定匹配\n%s",
// 		string(checkInvoke(t,stub,[][]byte{
// 			[]byte("queryDoMatch"),
// 			[]byte("Turn"),
// 		}).Payload)))

// }

func Test_CreateReceiver(t *testing.T) {
	//name := "E1"
	//resource := "E1,3,De,Fr,Bt,Ty"

	stub := initTest(t)

	checkInvoke(t, stub, [][]byte{
		[]byte("createReceiver"),
		[]byte("E1"),             //操作人
		[]byte("De,Fr,Bt,Ty,By"), //receiver
		//[]byte("De"),           //制造服务

	})
	fmt.Println("25.测试创建制造商")
}

//手动创建一些制造商

//***************************************************************//
//测试账户查询 queryPlusAccountList
func Test_QueryPlusAccountList(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试获取所有账户数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryPlusAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取多个账户数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryPlusAccountList"),
			[]byte("5feceb66ffc8"),
			[]byte("6b86b273ff34"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取单个账户数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryPlusAccountList"),
			[]byte("d4735e3a265e"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("4、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("0"),
		}).Payload)))
}
