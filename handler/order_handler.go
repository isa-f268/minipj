package handler

import (
	"main/dto"
	"main/helper"
	"main/model"
	"main/service"
	"main/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
)

type OrderHandler struct {
	serv service.OrderService
}

func NewOrderHandler(serv service.OrderService) *OrderHandler {
	return &OrderHandler{serv}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var o dto.OrderReq
	user_id := c.Get("id").(int)

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, amount, err := h.serv.CreateOrder(user_id, o)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	orderResp := struct {
		Order  model.Orders
		Amount int `json:"amount"`
	}{order, amount}

	resp := helper.RespHelper("create order success", orderResp)
	return c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) CreatePayment(c echo.Context) error {
	var p dto.PaymentReq
	user_id := c.Get("id").(int)

	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	payment, err := h.serv.CreatePayments(user_id, p.Order_id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := helper.RespHelper("create payment success", payment)

	return c.JSON(http.StatusCreated, res)

}

func (h *OrderHandler) PayMidtrans(c echo.Context) error {
	var req dto.MidtransReq
	user_id := c.Get("id").(int)

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, payment, err := h.serv.PaymentMidtrans(user_id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	strId := strconv.Itoa(int(payment.Payment_id))

	tokenUrl := utils.MidtransPayment(strId, req.Amount, user.Name, user.Email)

	res := struct {
		TokenUrl   midtrans.SnapResponse `json:"token_url"`
		Payment_id int                   `json:"payment_id"`
	}{
		TokenUrl:   tokenUrl,
		Payment_id: int(payment.Payment_id),
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *OrderHandler) UpdateStatusPayment(c echo.Context) error {
	type statusMid struct {
		Payment_id int `json:"Payment_id"`
	}

	var o statusMid

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := strconv.Itoa(o.Payment_id)

	res := utils.GetStatus(id)
	err := h.serv.UpdatePaymentMidtrans(o.Payment_id, res.StatusMessage)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	msg := map[string]string{"message": "your payment finish"}
	return c.JSON(http.StatusOK, msg)
}
