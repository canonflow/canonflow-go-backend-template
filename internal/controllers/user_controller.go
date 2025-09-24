package controllers

import (
	"net/http"

	"canonflow-golang-backend-template/internal/middlewares"
	"canonflow-golang-backend-template/internal/models/converter"
	"canonflow-golang-backend-template/internal/models/domain"
	"canonflow-golang-backend-template/internal/models/web"
	"canonflow-golang-backend-template/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserController struct {
	Log     *logrus.Logger
	Service *services.UserService
}

func NewUserController(service *services.UserService, log *logrus.Logger) *UserController {
	return &UserController{
		Log:     log,
		Service: service,
	}
}

func (c *UserController) SignUp(ctx *gin.Context) {
	// TODO: Validate
	var request web.UserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		ctx.Abort()
		return
	}

	// TODO: Find The User by username
	var user domain.User

	err := c.Service.UserRepository.FindByUsername(c.Service.DB, &user, request.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		// unexpected error
		ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		ctx.Abort()
		return
	}

	//! If the username is exists, then reject the request
	if err == nil {
		// user found → username taken
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "Username already taken",
		})
		ctx.Abort()
		return
	}

	// if user.ID != 0 {
	// 	ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Error:  "Username already taken",
	// 	})
	// 	ctx.Abort()
	// 	return

	// TODO: Create user
	user, err = c.Service.Create(ctx, request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "Failed to create user: " + err.Error(),
		})
		ctx.Abort()
		return
	}

	// TODO: Return the response
	ctx.JSON(http.StatusOK, web.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   converter.ToUserData(&user),
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	// TODO: Validate the Request
	var request web.UserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		ctx.Abort()
		return
	}

	// TODO: Find the user by username
	var user domain.User
	err := c.Service.UserRepository.FindByUsername(c.Service.DB, &user, request.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Error:  "Username not found",
		})
		ctx.Abort()
		return
	}

	// TODO: Login
	response, err := c.Service.Login(&user, request.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "Invalid username or password",
		})
		ctx.Abort()
		return
	}

	// TODO: Create Access Token
	tokenString, err := c.Service.CreateAccessToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		ctx.Abort()
		return
	}

	// TODO: Create HTTP Only Cookie
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(middlewares.TOKEN_COOKIE, tokenString, 3600*6, "/", "", true, true)

	// ctx.Header("Authorization", tokenString)
	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Logout(ctx *gin.Context) {
	middlewares.DeleteToken(ctx)

	ctx.JSON(http.StatusOK, web.SuccessResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   "Logged out successfully",
	})
}
