package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	. "progress-manage-system/utils/ecode"
)

func Validate(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Printf("%+v\n", v)
		}
		return ErrValidFail //validate不通过
	}

	return nil
}
