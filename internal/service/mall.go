package service

type IMallService interface{}

type MallService struct{}

func NewMallService() IMallService {
	return &MallService{}
}
