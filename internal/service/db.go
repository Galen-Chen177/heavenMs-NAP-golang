package service

import (
	"gorm.io/gorm"

	"heavenMs-NAP-golang/internal/db"
	"heavenMs-NAP-golang/internal/model"
)

type IDbService interface {
	// 将所有用户登录状态置为未登录
	ResetAllLoginStatus() (int64, error)
	// 重置最后获得雇佣商人的时间
	ResetAllLastGainHiredMerchant() (int64, error)
}

type Service struct {
	db *gorm.DB
}

func NewDbService() IDbService {
	return &Service{
		db: db.GormDB,
	}
}

func (s *Service) ResetAllLoginStatus() (int64, error) {
	var count int64
	err := s.db.Model(&model.Account{}).Update("loggedin", 0).Count(&count).Error
	return count, err
}

func (s *Service) ResetAllLastGainHiredMerchant() (int64, error) {
	var count int64
	err := s.db.Model(&model.Account{}).Update("lastGainHM", 0).Count(&count).Error
	return count, err
}
