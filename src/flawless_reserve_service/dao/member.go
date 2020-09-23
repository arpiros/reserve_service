package dao

import (
	"flawless_reserve_service/flawless_reserve_service/models"
)

func GetMember(mId int) *models.Member {
	var member models.Member
	err := commonDb.Where("id = ?", mId).First(&member).Error
	if err != nil {
		println("get member error")
	}
	return &member
}
