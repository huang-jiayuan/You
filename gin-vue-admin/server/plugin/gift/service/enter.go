package service

var Service = new(service)

type service struct {
	UserGiftBackpack userGiftBackpack
	GiftSendRecord   giftSendRecord
	GiftInfo         giftInfo
}
