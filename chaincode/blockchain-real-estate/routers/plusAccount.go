package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

// QueryAccountList 查询账户列表
func QueryPlusAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var plusAccountList []lib.PlusAccount
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.PlusAccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var plusAccount lib.PlusAccount
			err := json.Unmarshal(v, &plusAccount)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryPlusAccountList-反序列化出错: %s", err))
			}
			plusAccountList = append(plusAccountList, plusAccount)
		}
	}
	plusAccountListByte, err := json.Marshal(plusAccountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryPlusAccountList-序列化出错: %s", err))
	}
	return shim.Success(plusAccountListByte)
}
