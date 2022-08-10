package errors
 
var (
    Success                   = NewError(200, "Success")
    ServerError               = NewError(500, "ServerError")
    InvalidParams             = NewError(400, "InvalidParams")
    NotFound                  = NewError(404, "NotFound")
    UnauthorizedAuthNotExist  = NewError(401, "UnauthorizedAuthNotExist")
    UnauthorizedTokenError    = NewError(401, "UnauthorizedTokenError")
    UnauthorizedTokenTimeout  = NewError(401, "UnauthorizedTokenTimeout")
    UnauthorizedTokenGenerate = NewError(401, "UnauthorizedTokenGenerate")
    TooManyRequests           = NewError(429, "TooManyRequests")
    CreateuserFailed          = NewError(429, "CreateuserFailed")
)