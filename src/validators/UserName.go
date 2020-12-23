package validators

import (
	"github.com/go-playground/validator/v10"
)

type UserName string

func init() {
	register("UserName", UserName("required,min=4").toFunc())
}

func (this UserName) toFunc() validator.Func {
	validatorError["UserName"] = "用户名必须大于4位数"
	return func(fl validator.FieldLevel) bool {
		uName, ok := fl.Field().Interface().(string)
		return ok && this.validate(uName)
	}
}

func (this UserName) validate(v string) bool {
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}

	return true
}
