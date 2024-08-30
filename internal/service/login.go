package service

type ILoginService interface{}

type LoginService struct{}

func NewLoginService() ILoginService {
	return &LoginService{}
}
