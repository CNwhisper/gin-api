package requests
 
type RegisterUserRequest struct {
    Name string `form:"name" binding:"required,max=100"`
    Username string `form:"username" binding:"required,max=100"`
    Password string `form:"password" binding:"required,max=100"`
}
 
type LoginUserRequest struct {
    Username string `form:"username" binding:"required,max=100"`
    Password string `form:"password" binding:"required,max=100"`
}


