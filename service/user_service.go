package service

import (
	"main/dto"
	"main/model"
	"main/repository"
)

type UserService interface {
	Register(model model.Users) (dto.RegisterResp, error)
	Login(model dto.LoginReq) (string, error)
	TopUp(id int, balance dto.TopUpReq) (dto.TopUpResp, error)
	GetBook() ([]model.Books, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (r *userService) Register(u model.Users) (dto.RegisterResp, error) {
	return r.repo.Register(u)
}

func (r *userService) Login(u dto.LoginReq) (string, error) {
	return r.repo.Login(u)
}
func (r *userService) TopUp(id int, balance dto.TopUpReq) (dto.TopUpResp, error) {
	return r.repo.TopUp(id, balance)
}
func (r *userService) GetBook() ([]model.Books, error) {
	return r.repo.GetBook()
}
