package handlers

import (
	"fmt"
	"net/http"

	"github.com/MalshanPerera/go-expense-tracker/modals"
	"github.com/MalshanPerera/go-expense-tracker/services"
	"github.com/MalshanPerera/go-expense-tracker/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {

	var user modals.LoginUserParams
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	loggedInUser, err := h.AuthService.Login(c.Request().Context(), user)
	if loggedInUser == nil {
		return c.JSON(http.StatusFound, err)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, loggedInUser)
}

func (h *AuthHandler) RegisterHandler(c echo.Context) error {

	var user modals.CreateUserParams
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	registeredUser, err := h.AuthService.Register(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, registeredUser)

}
