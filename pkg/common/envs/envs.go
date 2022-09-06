package envs

import (
	"os"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {
	value := viper.Get(key)
	if value == "xxx" {
		return os.Getenv(key)
	}
	return value.(string)

}
