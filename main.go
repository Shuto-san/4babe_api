package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Shuto-san/4babe_api/config"
	"github.com/Shuto-san/4babe_api/infrastructure/datastore"
	"github.com/Shuto-san/4babe_api/infrastructure/router"
	"github.com/Shuto-san/4babe_api/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	gin := gin.New()
	gin = router.NewRouter(gin, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := gin.Run(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
