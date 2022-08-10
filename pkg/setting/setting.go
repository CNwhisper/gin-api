package setting
 
import "github.com/spf13/viper"
 
type Setting struct {
    viper *viper.Viper
}
 
func NewSetting() (*Setting, error) {
    viper := viper.New()
    viper.SetConfigName("config")
    viper.AddConfigPath("configs/")
    viper.SetConfigType("yaml")
    err := viper.ReadInConfig()
 
    if err != nil {
        return nil, err
    }
 
    return &Setting{viper}, nil
}