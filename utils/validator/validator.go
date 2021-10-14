package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	"github.com/go-playground/validator/v10"
	unTrans "github.com/go-playground/universal-translator"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"go-vue-api/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string,int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	//验证器注册翻译器
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}

	//将字段的英文转为中文
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE

}