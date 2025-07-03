package service

import (
	"main/dto"
	"main/model"
	"main/repository"
)

type OrderService interface {
	CreateOrder(user_id int, o dto.OrderReq) (model.Orders, int, error)
	CreatePayments(user_id int, order_id int) (model.Payments, error)
	PaymentMidtrans(user_id int, m dto.MidtransReq) (model.Users, model.Payments, error)
	UpdatePaymentMidtrans(payment_id int, message string) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo}
}

func (r *orderService) CreateOrder(user_id int, o dto.OrderReq) (model.Orders, int, error) {
	return r.repo.CreateOrder(user_id, o)
}

func (r *orderService) CreatePayments(user_id int, order_id int) (model.Payments, error) {
	return r.repo.CreatePayments(user_id, order_id)
}

func (r *orderService) PaymentMidtrans(user_id int, m dto.MidtransReq) (model.Users, model.Payments, error) {
	return r.repo.PaymentMidtrans(user_id, m)
}

func (r *orderService) UpdatePaymentMidtrans(payment_id int, message string) error {
	return r.repo.UpdatePaymentMidtrans(payment_id, message)
}
