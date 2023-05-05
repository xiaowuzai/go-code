package main

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/xiaowuzai/go-code/learn_gin/translate"
	mvalidator "github.com/xiaowuzai/go-code/learn_gin/validator"
)

func main() {
	trans, err := translate.InitTrans(translate.LOCATE_ZH)
	if err != nil {
		panic(err)
	}
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", mvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
