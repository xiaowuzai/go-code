package translate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Locale string

const (
	LOCATE_ZH Locale = "zh"
	LOCATE_EN Locale = "en"
)

func InitTrans(locale Locale) (ut.Translator, error) {
	//修改 gin 框架中的 validator 引擎属性, 实现定制
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil, errors.New("not a validator engine")
	}

	//注册一个获取 json的 tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := ut.New(enT, zhT, enT)

	trans, ok := uni.GetTranslator(string(locale))
	if !ok {
		return nil, fmt.Errorf("uni.GetTranslator(%s)", locale)
	}

	switch locale {
	case "en":
		en_translations.RegisterDefaultTranslations(v, trans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(v, trans)
	default:
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	return trans, nil
}
