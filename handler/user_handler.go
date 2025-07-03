package handler

import (
	"main/dto"
	"main/helper"
	"main/model"
	"main/service"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	serv service.UserService
}

func NewUserHandler(serv service.UserService) *UserHandler {
	return &UserHandler{serv}
}

// POST /api/users/register
// @Summary      Register User
// @Description  Register user baru app
// @Tags         User
// @Accept     	 json
// @Produce      json
// @Param      	 data body model.Users true "User Data"
// @Success      201   {object} helper.RegisterResp
// @Failure      401   {object} dto.ErrorResponse
// @Router       /api/users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var u model.Users

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(hashed)
	user, err := h.serv.Register(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := helper.RespHelper("register user success", user)
	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var u dto.LoginReq
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.serv.Login(u)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Email atau password salah")
	}

	token := map[string]string{"token": resp}

	res := helper.RespHelper("login success", token)
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) TopUp(c echo.Context) error {
	var t dto.TopUpReq

	user_id := c.Get("id").(int)

	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	top, err := h.serv.TopUp(user_id, t)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := helper.RespHelper("top up berhasil", top)

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetBook(c echo.Context) error {

	book, err := h.serv.GetBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := helper.RespHelper("get book sukses", book)

	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetInterBooks(c echo.Context) error {
	resp, err := utils.GetInterNationalBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
