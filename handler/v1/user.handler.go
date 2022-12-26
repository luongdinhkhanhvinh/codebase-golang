package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/luongdinhkhanhvinh/golang_heroku/common/obj"
	"github.com/luongdinhkhanhvinh/golang_heroku/common/response"
	"github.com/luongdinhkhanhvinh/golang_heroku/dto"
	"github.com/luongdinhkhanhvinh/golang_heroku/service"
)

type User struct {
}

type UserHandler interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserHandler(
	userService service.UserService,
	jwtService service.JWTService,
) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userHandler) getUserIDByHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

// CreateTodo godoc
// @Summary Update profile user
// @Description Update profile user
// @Tags Update user
// @Accept json
// @Produce json
// @Param user body entity.User true "Value"
// @Success 201 {object} entity.User
// @Router /api/user/profile [put]
func (c *userHandler) Update(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserRequest

	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := c.getUserIDByHeader(ctx)

	if id == "" {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_id, _ := strconv.ParseInt(id, 0, 64)
	updateUserRequest.ID = _id
	res, err := c.userService.UpdateUser(updateUserRequest)

	if err != nil {
		response := response.BuildErrorResponse("Error", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := response.BuildResponse(true, "OK", res)
	ctx.JSON(http.StatusOK, response)

}

// @Summary		Get profile User
// @Description	Get profile User
// @Tags Get User
// @Accept			json
// @Produce		json
// @Param			file	formData	file			true	"this is a test file"
// @Success		200		{string}	string			"ok"
// @Router			/api/user/profile [get]
func (c *userHandler) Profile(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.FindUserByID(id)

	if err != nil {
		response := response.BuildErrorResponse("Error", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	res := response.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
