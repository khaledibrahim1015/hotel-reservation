package types

import "github.com/khaledibrahim1015/hotel-reservation/utils"

// request body
type CreateUserParam struct {
	FirstName string ` json:"firstName"`
	LastName  string ` json:"lastName"`
	Email     string ` json:"email"`
	Password  string ` json:"password"`
}

type User struct {
	// ID        string `bson:"_id" json:"id,omitempty"`
	ID               string `bson:"_id,omitempty" json:"id"`
	FirstName        string `bson:"firstName" json:"firstName"`
	LastName         string `bson:"lastName"lastName"`
	Email            string `bson:"email" json:"email"`
	EncryptedPaaword string `bson:"encryptedPaaword" json:"-"`
}

func NewUserFromParams(param CreateUserParam) (*User, error) {

	encryptedPassword, err := utils.HashPassword(param.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:        param.FirstName,
		LastName:         param.LastName,
		Email:            param.Email,
		EncryptedPaaword: encryptedPassword,
	}, nil

}
