package database

import (
	"errors"

	"github.com/rajesh4295/graphql-server/models"
)

var userData = []*models.User{
	{
		Id:       "1",
		Name:     "Bob",
		WalletId: "1",
	},
	{
		Id:       "2",
		Name:     "Alex",
		WalletId: "2",
	},
}

// var walletData = []*models.Wallet{}

// var transactionData = []*models.Transaction{}

func GetUsers() ([]*models.User, error) {
	return userData, nil
}

func GetUserById(id string) (*models.User, error) {
	for _, u := range userData {
		if u.Id == id {
			return u, nil
		}
	}
	return &models.User{}, errors.New("user not found")
}

func UpdateUserById(user *models.User) (*models.User, error) {
	existingUser, err := GetUserById(user.Id)
	if err != nil {
		return nil, err
	}
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.WalletId != "" {
		existingUser.WalletId = user.WalletId
	}
	return existingUser, nil
}

func AddUser(user *models.User) (*models.User, error) {
	userData = append(userData, user)
	return user, nil
}

func RemoveUserById(user *models.User) (*models.User, error) {
	var pos = -1
	for i, u := range userData {
		if user.Id == u.Id {
			pos = i
		}
	}
	if pos > -1 {
		userData = append(userData[:pos], userData[pos+1:]...)
		return user, nil
	}
	return &models.User{}, errors.New("user not found")
}
