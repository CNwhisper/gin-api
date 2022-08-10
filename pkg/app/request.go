package app
 
import (
    "github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
)
 
// func BindAndValidation(c *gin.Context, v interface{}) (bool, []string) {
//     var errors []string
 
//     if err := c.ShouldBind(v); err != nil {
//         for _, e := range err.(validator.ValidationErrors) {
//             errors = append(errors, e.Error())
//         }
//         return false, errors
//     }
 
//     return true, nil
// }

func BindAndValidation(c *gin.Context, v interface{}) (bool, []string) {
    var errors []string
 
    if err := c.ShouldBind(v); err != nil {
        translatorVal := c.Value("translator")  // 抓到中介層抓到的語系
        translator, _ := translatorVal.(ut.Translator)
        validationErrors, ok := err.(validator.ValidationErrors)
        if !ok {
            return false, errors
        }
 
        for _, value := range validationErrors.Translate(translator) {
            errors = append(errors, value)
        }
 
        return false, errors
    }
 
    return true, nil
}