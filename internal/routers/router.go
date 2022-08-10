package routers
 
// import "github.com/gin-gonic/gin"
 
// func NewRouter() *gin.Engine {
//     router := gin.New()
 
//     router.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{})
//     })
 
//     router.GET("/test", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "success": true,
//             "data": "Hello world.",
//         })
//     })
 
//     return router
// }
 
import (
    "github.com/CNwhisper/gin-api/internal/routers/api"
	"github.com/CNwhisper/gin-api/internal/middleware"
    "github.com/gin-gonic/gin"
)


	
 
func NewRouter() *gin.Engine {
	// router := gin.New()
	
    router := gin.New()
 	router.Use(middleware.Translations())

    hello := api.NewHello()
    router.GET("/", hello.Index)
    router.GET("/test", hello.Test)
	router.POST("/form", hello.PostForm)
	router.POST("/json", hello.PostJson)
	router.GET("/:id", hello.Get)

	user := api.NewUser()
	router.POST("/register", user.Register)
	router.POST("/login", user.Login)
 
    return router
}