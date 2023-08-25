package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "ok"})
}

type User struct {
	Name string
	Age  int
}

var UsersDB = map[string]User{"1": User{Name: "arshia", Age: 19}}

func (t *TestHandler) UsersList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, UsersDB)
}

func (t *TestHandler) GetUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	u, ok := UsersDB[id]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"Status": "User Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, u)

}

func (t *TestHandler) CreateUser(ctx *gin.Context) {
	b := ctx.GetStringMapString("id")
	fmt.Println(b)
}

type Header struct {
	UserID string
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	userID := ctx.GetHeader("UserID")
	ctx.JSON(http.StatusOK, gin.H{
		"result": userID,
		"func":   "headerBinder1",
	})
}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	head := Header{}
	ctx.BindHeader(&head)
	ctx.JSON(http.StatusOK, gin.H{
		"result": head,
		"func":   "headerBinder1",
	})
}

func (h *TestHandler) QueryBinder(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"id":   ids,
	})
}

func (h *TestHandler) UriBinder(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"id":   id,
	})
}

type Person struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required,alpha,min=4,max=10"`
	Phone    string `json:"phone" binding:"phone,required"`
	Password string `json:"password" binding:"password,required"`
}

func (h *TestHandler) BodyBinder(ctx *gin.Context) {
	p := Person{}
	// ctx.Bind(&p)
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":     p.Name,
		"id":       p.Id,
		"number":   p.Phone,
		"password": p.Password,
	})
}

func (h *TestHandler) FormBinder(ctx *gin.Context) {
	p := Person{}
	// ctx.Bind(&p)
	ctx.ShouldBind(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"name": p.Name,
		"id":   p.Id,
		"p":    p,
	})
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	ctx.SaveUploadedFile(file, "file")
	// ctx.Bind(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"name": "",
		"id":   file.Filename,
	})
}
