package service

type IChannelService interface{}

type ChannelService struct{}

func NewChannelService() IChannelService {
	return &ChannelService{}
}
