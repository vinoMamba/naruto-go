package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/vinoMamba/naruto-go/cmd"
)

// @title            Naruto Go API
// @version         1.0
// @description

// @contact.name   vino
// @contact.url    https://baidu.com
// @contact.email  vino0908@outlook.com

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute()
}
