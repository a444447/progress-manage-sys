package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"progress-manage-system/utils/errmsg"
)

func Validate(data interface{}) int {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Printf("%+v\n", v)
		}
		return errmsg.ErrorValidate //validate不通过
	}

	return errmsg.SUCCESS
}
