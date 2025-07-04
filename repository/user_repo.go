package repository

import (
	"main/dto"
	"main/helper"
	"main/model"
	"main/utils"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(model model.Users) (dto.RegisterResp, error)
	Login(model dto.LoginReq) (string, error)
	TopUp(id int, balance dto.TopUpReq) (dto.TopUpResp, error)
	GetBook() ([]model.Books, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user model.Users) (dto.RegisterResp, error) {
	var resp dto.RegisterResp
	err := r.db.Create(&user).Error
	if err != nil {
		return resp, err
	}
	resp.Email = user.Email
	resp.Full_name = user.Name
	return resp, nil
}

func (r *userRepository) Login(user dto.LoginReq) (string, error) {
	var u model.Users

	err := r.db.Where("email=?", user.Email).First(&u).Error

	if err != nil {
		return "", utils.ErrUserNotFound
	}

	err = helper.CheckPassword(u.Password, user.Password)

	if err != nil {
		return "", utils.ErrUnauthorized
	}

	data := helper.Login{
		Id:   int(u.User_id),
		Name: u.Name,
	}

	token, err := helper.CreateJwt(data)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userRepository) TopUp(id int, balance dto.TopUpReq) (dto.TopUpResp, error) {
	var t model.Topup
	var resp dto.TopUpResp
	var u model.Users

	today := time.Now()
	t.User_id = id
	t.Balance = balance.Balance
	t.TopupDate = today
	err := r.db.Create(&t).Error

	if err != nil {
		return dto.TopUpResp{}, err
	}

	err = r.db.Model(&model.Users{}).
		Where("user_id=?", id).
		Update("total_balance", gorm.Expr("total_balance + ?", balance.Balance)).Error

	if err != nil {
		return dto.TopUpResp{}, err
	}

	err = r.db.Where("user_id=?", id).First(&u).Error

	if err != nil {
		return dto.TopUpResp{}, utils.ErrUserNotFound
	}

	resp.Balance = balance.Balance
	resp.Total_balance = u.TotalBalance
	resp.Topup_date = today.String()

	return resp, err
}

func (r *userRepository) GetBook() ([]model.Books, error) {
	var b []model.Books

	err := r.db.Find(&b).Error

	if err != nil {
		return []model.Books{}, err
	}

	return b, err
}
