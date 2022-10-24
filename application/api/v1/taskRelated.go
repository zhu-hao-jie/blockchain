package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

type OrderStrBody struct {
	Submitter string `json:submitter` //任务发布者
	Task      string `json:task`      //任务名称
	Status    string `json:status`
	Time      string `json:time`
}

type OrderStrQueryBody struct {
	Submitter string `json:submitter`
}

type UpdateOrderStrBody struct {
	Submitter string `json:submitter`
	Task      string `json:task` //任务名称
	Status    string `json:status`
	Time      string `json:time`
}

type QueryDoMatchBody struct {
	Owner string `json:owner`
}

type QueryMatchBody struct {
	Owner string `json:owner`
}

func CreateOrderStr(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderStrBody)
	//解析智能合约
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Submitter))
	bodyBytes = append(bodyBytes, []byte(body.Task))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.Time))
	//调用智能合约
	resp, err := bc.ChannelExecute("createOrderStr", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{} //此处 []map[string]interface{} map前的“[]”号待明确
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryOrderStr(c *gin.Context) { //以 submitter 进行查询
	appG := app.Gin{C: c}
	body := new(OrderStrQueryBody)
	//解析 Body 参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Submitter != "" {
		bodyBytes = append(bodyBytes, []byte(body.Submitter))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryOrderStr", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{} //此处 []map[string]interface{} map前的“[]”号待明确
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func UpdateOrderStr(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateOrderStrBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Submitter == "" || body.Task == "" || body.Status == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Submitter))
	bodyBytes = append(bodyBytes, []byte(body.Task))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.Time))

	//调用智能合约
	resp, err := bc.ChannelExecute("updateOrderStr", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data []map[string]interface{} //此处 []map[string]interface{} map前的“[]”号待明确
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//查询匹配后的结果
func QueryDoMatch(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(QueryDoMatchBody)
	//解析 Body 参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Owner != "" {
		bodyBytes = append(bodyBytes, []byte(body.Owner))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryDoMatch", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	//反序列化
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//查询匹配
func QueryMatch(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(QueryMatchBody)
	//解析Body 参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Owner != "" {
		bodyBytes = append(bodyBytes, []byte(body.Owner))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("doMatch", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	//反序列化 json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}