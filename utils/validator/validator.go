package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()
	chinese  = zh.New()
	uni      = ut.New(chinese, chinese)
	trans, _ = uni.GetTranslator("zh")
)

func Validate(data interface{}) (string, error) {
	_ = zht.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(data); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), err
		}
	}
	return "", nil
}
