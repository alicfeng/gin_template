package request

import (
	validator2 "github.com/alicfeng/gin_template/app/validator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// FormRequest 规则手册 https://github.com/go-playground/validator/blob/master/README.md
type FormRequest struct {
}

type PageRequest struct {
	Page    int    `form:"page" comment:"当前页" binding:"required,gte=1"`
	Limit   int    `form:"limit" comment:"页数量" binding:"required,gte=1,lte=100"`
	Keyword string `form:"keyword" comment:"关键词"`
}

var (
	uni *ut.UniversalTranslator
	//validate *validator.Validate
	trans ut.Translator
)

func init() {
	// 注册翻译器
	langZh := zh.New()
	uni = ut.New(langZh, langZh)
	trans, _ = uni.GetTranslator("translator")
	// 获取gin的校验器
	validation := binding.Validator.Engine().(*validator.Validate)
	// 注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(validation, trans)
}

//Translate 翻译错误信息 */
func (request *FormRequest) Translate(value reflect.Value, err error) string {
	var result []string

	// 1.将错误转成 ValidationErrors 格式错误
	errors := err.(validator.ValidationErrors)

	// 2.获取 request 的反射 value 值 ⚠️一定是值不能是地址
	valueOf := value.Type()
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	// 3.将每一个表单不满足的错误信息翻译
	for _, err := range errors {
		// 3.1翻译错误信息
		zhError := err.Translate(trans)
		// 3.2加载自定义标签错误信息
		if text, exist := validator2.Message[err.Tag()]; true == exist {
			zhError = err.StructField() + text
		}
		// 3.3根据字段名称获取结构体标签键值对 约定.comment
		structField, success := valueOf.FieldByName(err.StructField())
		if success && "" != structField.Tag.Get("comment") {
			zhError = strings.Replace(zhError, err.StructField(), structField.Tag.Get("comment"), -1)
		}
		// 3.4错误信息拼接起来
		result = append(result, zhError)
	}

	return strings.Join(result, ",")
}
