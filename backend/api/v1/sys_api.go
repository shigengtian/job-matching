package v1

import (
	"github.com/gin-gonic/gin"
	"job-maching/model/request"
	"job-maching/model/response"
)

func CreateApi(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func DeleteApi(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}
func GetApiList(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func GetApiById(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func UpdateApi(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func DeleteApisByIds(c *gin.Context) {
	var api request.SearchApiReq
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}
