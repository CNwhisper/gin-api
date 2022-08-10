package middleware
 
import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/locales/en"
    "github.com/go-playground/locales/zh_Hant_TW"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    enTranslation "github.com/go-playground/validator/v10/translations/en"
    zhTranslation "github.com/go-playground/validator/v10/translations/zh_tw"
)
 
func Translations() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        universalTranslator := ut.New(en.New(), zh_Hant_TW.New())  // 英文與繁體中文
        locale := ctx.GetHeader("locale") // 從 header 抓語系
        translator, _ := universalTranslator.GetTranslator(locale)
        v, ok := binding.Validator.Engine().(*validator.Validate)
        if ok {
            switch locale {
            case "en":
                _ = enTranslation.RegisterDefaultTranslations(v, translator)
                break
            default:
                _ = zhTranslation.RegisterDefaultTranslations(v, translator) // 預設中文
                break
            }
            ctx.Set("translator", translator) // 存一個值供之後拿
        }
 
        ctx.Next()  // 允許通過這個中介層
    }
}