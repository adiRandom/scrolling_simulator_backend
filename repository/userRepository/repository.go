package userRepository

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib/functional"
	"backend_scrolling_simulator/repository"
)

func FindUserById(id uint) (*dtos.User, error) {
	user := dtos.User{}
	db := repository.GetDB()
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUsersByIds(ids []uint) ([]dtos.User, error) {
	var users []dtos.User
	db := repository.GetDB()
	err := db.Where("id IN ?", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}

	sorted := functional.Associate(ids, users, func(id uint, user dtos.User) bool {
		return user.Id == id
	})

	return functional.Values(sorted), nil
}
