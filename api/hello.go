package api

import (
	"github.com/gin-gonic/gin"
)

// hello godoc
// @Summary Say Hello
// @Description Say Hello
// @Tags api
// @Acception json
// @Produce json
// @Router /test [get]
//另apifox中的测试环境可以添加前置的端口号，如果使用/swagger的url访问则可以直接用localhost
//要辩识上面的注释，不能有空行
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})

}

//swag init命令要在main.go所在的目录下执行，如果你的主函数是其他名字例如code.go,指令为swag init -g code.go
//swagger会搜索从main.go所在目录下的所有包的所有函数是否有swagger形式的注释并形成文档
//使用自定义的结构json返回的时候要在结构前加一个{object}或者object
//
