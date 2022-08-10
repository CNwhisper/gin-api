package errors
 
import "fmt"
 
type Error struct {
    Code    int  `json:"cond"`
    Message string   `json:"message"`
    Details []string `json:"details"`
}
 
func NewError(code int, message string) *Error {
    return &Error{Code: code, Message: message}
}
 
func (e *Error) GetError() string {
    return fmt.Sprintf("Error: %d, message: %s", e.GetCode(), e.GetMessage())
}
 
func (e *Error) GetCode() int {
    return e.Code
}
 
func (e *Error) GetMessage() string {
    return e.Message
}
 
func (e *Error) GetDetails() []string {
    return e.Details
}
 
func (e *Error) WithDetails(details ...string) *Error {
    newError := *e
    newError.Details = append(newError.Details, details...)
    return &newError
}