package types

import (
	"github.com/khaledibrahim1015/hotel-reservation/utils"
	"github.com/khaledibrahim1015/hotel-reservation/utils/validator"
)

// request body
type CreateUserParam struct {
	FirstName string `validate:"required,min=3,max=50" json:"firstName"`
	LastName  string `validate:"required,min=3,max=50" json:"lastName"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=7,max=50" json:"password"`
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

func (ve *CreateUserParam) Validate() error {

	v := validator.New()
	if err := v.Validate(ve); err != nil {
		return err
	}
	return nil
}
