package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"
	"peanut/pkg/i18n"

	"gorm.io/gorm"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth
func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	i18n.SetupI18n()

	dbClient := dbConnect()
	infra.Migration(dbClient)
	server := infra.SetupServer(dbClient)

	server.Router.Run(":8081")
}

func dbConnect() *gorm.DB {
	db, err := infra.ConnectMySQL()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
