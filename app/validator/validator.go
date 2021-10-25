package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Register() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("exist", Exist)
		err = v.RegisterValidation("unique", Unique)
	}

	return err
}
