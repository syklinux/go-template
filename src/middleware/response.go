package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// ResponseCode ResponseCode
type ResponseCode int

// 1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	OkCode      ResponseCode = 200
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode
	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000
	GroupallSaveFlowerror   ResponseCode = 2001
)

// ResponseWithData ResponseWithData1
type ResponseWithData struct {
	Status  ResponseCode `json:"status"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

// Response Response
type Response struct {
	Status  ResponseCode `json:"status"`
	Message string       `json:"message"`
}

// ResponseForbidden ResponseForbidden
func ResponseForbidden(c *gin.Context, status ResponseCode, err error) {
	resp := &Response{Status: status, Message: err.Error()}
	c.JSON(403, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseUnauthorized ResponseUnauthorized
func ResponseUnauthorized(c *gin.Context, status ResponseCode, err error) {
	resp := &Response{Status: status, Message: err.Error()}
	c.JSON(401, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseStatusError ResponseStatusError
func ResponseStatusError(c *gin.Context, status ResponseCode, err error) {
	resp := &Response{Status: status, Message: err.Error()}
	c.JSON(int(status), resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseError ResponseError
func ResponseError(c *gin.Context, status ResponseCode, err error) {
	resp := &Response{Status: status, Message: err.Error()}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseSuccessWithData ResponseSuccessWithData
func ResponseSuccessWithData(c *gin.Context, msg string, data interface{}) {
	resp := &ResponseWithData{Status: OkCode, Message: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseErrorWithData ResponseErrorWithData1
func ResponseErrorWithData(c *gin.Context, msg string, data interface{}) {
	resp := &ResponseWithData{Status: 500, Message: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseWithList ResponseWithList1
func ResponseWithList(c *gin.Context, data interface{}) {
	c.JSON(200, data)
	response, _ := json.Marshal(data)
	c.Set("response", string(response))
}

// ResponseSuccess ResponseSuccess
func ResponseSuccess(c *gin.Context, msg string) {
	resp := &Response{Status: OkCode, Message: msg}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
