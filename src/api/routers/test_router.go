package routers

import (
	"wapi/src/api/handler"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handler.NewTestHandler()

	r.GET("/", handler.Test)
	r.GET("/users", handler.UsersList)
	r.GET("/user/:id", handler.GetUser)
	// r.GET("/user-username/:username", handler.Test)
	r.POST("/user/create", handler.CreateUser)
	r.POST("/header1/", handler.HeaderBinder1)
	r.POST("/header2/", handler.HeaderBinder2)
	r.POST("/qq/", handler.QueryBinder)
	r.POST("/uri/:id/:name/", handler.UriBinder)
	r.POST("/body/", handler.BodyBinder)
	r.POST("/form/", handler.FormBinder)
	r.POST("/file/", handler.FileBinder)

}
