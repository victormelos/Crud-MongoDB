package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"github.com/victormelos/curso-youtube/src/interfaces/dto"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUserHandler(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDomain := &user.UserDomain{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Age:      userRequest.Age,
	}

	result, err := h.userService.Create(userDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Age:   result.Age,
	})
}

func (h *UserHandler) FindUserByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Age:   result.Age,
	})
}

func (h *UserHandler) FindUserByEmail(c *gin.Context) {
	email := c.Param("email")

	result, err := h.userService.FindByEmail(email)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Age:   result.Age,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var userRequest dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDomain := &user.UserDomain{
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	if err := h.userService.Update(id, userDomain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.userService.Delete(id); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusNoContent)
}
