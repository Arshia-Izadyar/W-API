package handler

import (
	"errors"
	"fmt"
	"net/http"
	"wapi/src/api/helper"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("ok", true, 0))
}

type User struct {
	Name string
	Age  int
}

var UsersDB = map[string]User{"1": {Name: "arshia", Age: 19}}

// UsersList godoc
// @Summary Gs
// @Description Get s ID
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Router /api/v1/test/users [get]
func (t *TestHandler) UsersList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(UsersDB, true, 0))

}

// GetUser godoc
// @Summary Get user br ID
// @Description Get user br ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 200 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Router /api/v1/test/user/{id} [get]
func (t *TestHandler) GetUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	u, ok := UsersDB[id]
	if !ok {
		ctx.JSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError("user not found", false, 0, errors.New("cant find user")))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(u, true, 0))

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
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": userID,
		"func":   "headerBinder1",
	}, true, 0))

}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	head := Header{}
	ctx.BindHeader(&head)
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": head,
		"func":   "headerBinder1",
	}, true, 0))
}

func (h *TestHandler) QueryBinder(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"name": name,
		"id":   ids,
	}, true, 0))

}

func (h *TestHandler) UriBinder(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"name": name,
		"id":   id,
	}, true, 0))
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
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
	err := ctx.ShouldBind(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(404, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"name": p.Name,
		"id":   p.Id,
		"p":    p,
	}, true, 0))
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	ctx.SaveUploadedFile(file, "file")
	// ctx.Bind(&p)
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"name": "",
		"id":   file.Filename,
	}, true, 0))
}
