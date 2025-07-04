package repository

import (
	"fmt"
	"main/dto"
	"main/model"
	"main/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(user_id int, o dto.OrderReq) (model.Orders, int, error)
	CreatePayments(user_id int, order_id int) (model.Payments, error)
	PaymentMidtrans(user_id int, m dto.MidtransReq) (model.Users, model.Payments, error)
	UpdatePaymentMidtrans(payment_id int, message string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(user_id int, o dto.OrderReq) (model.Orders, int, error) {
	var order model.Orders
	var book model.Books
	today := time.Now()
	order.User_id = user_id
	order.Book_id = o.Book_id
	order.TotalDays = o.Total_days
	order.Rent_date = today

	r.db.Where("book_id=?", o.Book_id).First(&book)

	err := r.db.Create(&order).Error
	if err != nil {
		return model.Orders{}, 0, err
	}
	amount := book.Price_per_day * o.Total_days
	return order, amount, nil
}

func (r *orderRepository) CreatePayments(user_id int, order_id int) (model.Payments, error) {
	var u model.Users
	var o model.Orders
	var b model.Books
	var p model.Payments

	err := r.db.Where("order_id=?", order_id).First(&o).Error

	if err != nil {
		return model.Payments{}, err
	}

	if o.User_id != user_id {
		return model.Payments{}, utils.ErrUnauthorized
	}

	r.db.Where("book_id", o.Book_id).First(&b)
	r.db.Where("user_id=?", user_id).First(&u)

	rent_cost := b.Price_per_day * o.TotalDays

	if u.TotalBalance < rent_cost {
		return model.Payments{}, utils.ErrBadReq
	}

	p.Order_id = order_id
	p.Amount = rent_cost
	p.Status = "paid"

	err = r.db.Create(&p).Error

	if err != nil {
		return model.Payments{}, err
	}

	return model.Payments{}, nil
}

func (r *orderRepository) PaymentMidtrans(user_id int, m dto.MidtransReq) (model.Users, model.Payments, error) {
	var u model.Users
	var p model.Payments
	p.Order_id = m.Order_id
	p.Amount = m.Amount
	p.Status = "pending"

	err := r.db.Where("user_id=?", user_id).First(&u).Error
	if err != nil {
		return model.Users{}, model.Payments{}, utils.ErrUnauthorized
	}

	err = r.db.Create(&p).Error

	return u, p, err
}

func (r *orderRepository) UpdatePaymentMidtrans(payment_id int, message string) error {
	var p model.Payments

	status_msg := strings.ToLower(message)
	fmt.Println(status_msg)
	if !strings.Contains(status_msg, "success") {
		return utils.ErrBadReq
	}

	err := r.db.Model(&p).Where("payment_id=?", payment_id).Update("status", "paid").Error
	if err != nil {
		return err
	}
	return nil
}
