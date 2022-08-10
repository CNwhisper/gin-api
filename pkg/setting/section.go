package setting
 
import "time"
 
type ServerSetting struct {
    RunMode      string
    HttpPort     string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}
 
type DatabaseSetting struct {
    DBType       string
    UserName     string
    Password     string
    Host         string
    DBName       string
    Charset      string
    ParseTime    bool
}
 
func (setting *Setting) ReadSection(k string, v interface{}) error {
    return setting.viper.UnmarshalKey(k, v)
}