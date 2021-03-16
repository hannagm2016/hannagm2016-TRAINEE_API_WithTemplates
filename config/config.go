package config
import (
"github.com/tkanos/gonfig"
"fmt"
)
type Configuration struct {
  ClientID string
  ClientSecret string
  RedirectURL string
}
func GetConfig(env string) Configuration {
  configuration := Configuration{}
  fileName := fmt.Sprintf("config/%s_config.json", env)
  gonfig.GetConf(fileName, &configuration)
  return configuration
}
