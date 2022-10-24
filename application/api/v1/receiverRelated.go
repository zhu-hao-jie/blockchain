package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"
	"github.com/gin-gonic/gin"
	"strings"
)

type ReceiverBody struct {
	ResName string   `json:"resName"`
	Process []string `json:"process"`
}

type ReceiverQueryBody struct {
	ResName string `json:resName`
}

func QueryReceiverRes(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ReceiverQueryBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.ResName != "" {
		bodyBytes = append(bodyBytes, []byte(body.ResName))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryReceiverRes", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func CreateReceiver(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ReceiverBody)
	//解析 Body 参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	BProcess := strings.Join(body.Process, ",") //将字符串数组转换为 字符串

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ResName))
	bodyBytes = append(bodyBytes, []byte(BProcess))

	//调用智能合约
	resp, err := bc.ChannelExecute("createReceiver", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
