package utilities

import (
	"github.com/gookit/validate"
	"github.com/kennedybaraka/fiber-api/pkg/models"
)

// login struct validation
func UserInputValidation(user models.User) error {
	v := validate.Struct(user)

	if !v.Validate() {
		return v.Errors
	}
	return nil
}
