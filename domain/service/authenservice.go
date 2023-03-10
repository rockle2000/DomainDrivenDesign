package service

import (
	"context"

	"DDD_Project/domain/model/entity"
	"DDD_Project/domain/repository"
	"DDD_Project/infrastructure/config"
	"DDD_Project/infrastructure/util"
)

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginRes struct {
		AccessToken string `json:"access_token"`
	}

	RegisterReq struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
)

type AuthenticationService interface {
	Login(req LoginReq) (*LoginRes, error)
	Register(req RegisterReq) error
}

type authenticationService struct {
	repo   repository.AuthenticationRepo
	config *config.AppConfig
}

func NewAuthenticationService(repo repository.AuthenticationRepo, config *config.AppConfig) AuthenticationService {
	return &authenticationService{
		repo:   repo,
		config: config,
	}
}

func (a *authenticationService) Login(req LoginReq) (*LoginRes, error) {
	ctx := context.Background()
	customerEntity := &entity.Customer{
		Email:    req.Email,
		Password: req.Password,
	}
	customerRes, err := a.repo.Login(ctx, customerEntity)
	if err != nil {
		return nil, err
	}
	j := util.NewJWT(a.config)
	token, err := j.GenerateToken(a.config, customerRes.Id, customerRes.Email, customerRes.Name)
	if err != nil {
		return nil, err
	}
	res := &LoginRes{
		AccessToken: token.AccessToken,
	}
	return res, nil
}

func (a *authenticationService) Register(req RegisterReq) error {
	ctx := context.Background()
	customerEntity := &entity.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return a.repo.Register(ctx, customerEntity)

}
