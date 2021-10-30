package userauth

import "AltaStore/business/user"

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	UserLogin(email string, password string) (string, error)

	//Create Token If Failed will return error
	CreateToken(user *user.User) (*TokenDetails, error)
}
