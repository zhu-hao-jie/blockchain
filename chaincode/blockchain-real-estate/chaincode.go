package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"strings"
	"time"
)

type BlockChainServices struct {
}

// Init 链码初始化
func (t *BlockChainServices) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = []string{"管理员", "A制造商", "B制造商", "C制造商", "D制造商", "E制造商"}
	//var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	var processName = []string{"", "De,Hu,Bo,Ki", "Gr,Mi,Sa,Tu", "Sa,Br,Ta", "Fil,Rea,Tu", "Bo,Tu,Gr,Sa"}
	// servicesNames := make([]string,6)
	// servicesNames = append(servicesNames,"De,Dr,Bo,M,W","De,Sh","Dr,Bo","Bo","M,De","W,Qz")
	//append(servicesNames,strings.Join({"De","Dr","Bo","M","W"},""), strings.Join({"De","Sh"},""), strings.Join({"Dr","Bo"},""), strings.Join({"Bo"},""), strings.Join({"M","De"},""), strings.Join({"W","Qz"},""))
	//[]string {[]string{"De","Dr","Bo","M","W"}, []string{"De","Sh"}, []string{"Dr","Bo"}, []string{"Bo"}, []string{"M","De"}, []string{"W","Qz"}}
	//servicesName := [6]string{{"De","Dr","Bo","M","W"}, {"De","Sh"}, {"Dr","Bo"}, "Bo", {"M","De"}, {"W","Qz"}}
	//初始化账号数据
	var proName [6][]string
	for i := 0; i < len(processName); i++ {
		for j := 0; j < len(processName[i]); j++ {
			proName[i] = strings.Split(processName[i], ",")
		}
	}

	for i, val := range accountIds {
		account := &lib.PlusAccount{
			AccountId: val,
			UserName:  userNames[i],
			//Balance:   balances[i],
			ProcessName: proName[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, lib.PlusAccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainServices) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryAccountList":
		return routers.QueryAccountList(stub, args)
	case "createRealEstate":
		return routers.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return routers.QueryRealEstateList(stub, args)
	case "createSelling":
		return routers.CreateSelling(stub, args)
	case "createSellingByBuy":
		return routers.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return routers.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return routers.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return routers.UpdateSelling(stub, args)
	case "createDonating":
		return routers.CreateDonating(stub, args)
	case "queryDonatingList":
		return routers.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return routers.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return routers.UpdateDonating(stub, args)
	// // ========== order function test ==============
	// case "createOrder":
	// 	return routers.CreateOrder(stub, args)
	// case "queryOrderList":
	// 	return routers.QueryOrderList(stub, args)
	// case "updateOrder":
	// 	return routers.UpdateOrder(stub, args)
	// case "queryOrderHistory":
	// 	return routers.QueryOrderHistory(stub, args)
	// default:
	// 	return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))

	//======================================================
	case "createReceiver":
		return routers.CreateReceiver(stub, args)
	case "createOrderStr":
		return routers.CreateOrderStr(stub, args)
	case "queryOrderStr":
		return routers.QueryOrderStr(stub, args)
	case "queryReceiverRes":
		return routers.QueryReceiverRes(stub, args)
	case "updateOrderStr":
		return routers.UpdateOrderStr(stub, args)
	case "queryPlusAccountList":
		return routers.QueryPlusAccountList(stub, args)
	case "doMatch":
		return routers.DoMatch(stub, args)
	case "queryDoMatch":
		return routers.QueryDoMatch(stub, args)
	// case "taskMatch":
	// 	return routers.TaskMatch(stub, args)
	// case "queryTaskMatch":
	// 	return routers.QueryTaskMatch(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能：%s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainServices))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

