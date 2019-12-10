package api

import (
	//"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/pkg/logging"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)


	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH_TOKEN
		}

	} else {
		for _, err := range valid.Errors {
			//log.Printf("err.key: %s, err.message: %s\n", err.Key, err.Message)
			// 使用自己的logger写日志
			logging.Info(err.Key, err.Message)
		}
	}

	// 使用自己的logger写日志, 此处仅仅为测试
	logging.Info("code:", code , "msg:", e.GetMsg(code), "data:", data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}