package main


import "github.com/gin-gonic/gin"

// User 参数绑定结构体
type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Hobby []string `form:"hobby"`
	Gender string `form:"gender"`
	City string `form:"city"`
}

func RegisterUser(c *gin.Context) {
	var user User
	// 查询的数据直接绑定到User内，不再需要一个一个拿
	c.ShouldBind(&user)
	// c.ShouldBindJson(&user)
	c.String(200, "User: %s", user)
}

func GoRegisterUser(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("../html/*")
	e.GET("/register", GoRegisterUser)
	e.POST("/register", RegisterUser)
	e.Run()
}

