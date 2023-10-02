package server

import (
	"fmt"
	"github.com/grip211/lessonsL2/develop/dev11/config"
	"log"
	"net/http"

	"gorm.io/gorm"
)

// Up поднимает сервер.
func Up(dbConnection *gorm.DB, cfg config.Server) {
	handling(dbConnection)

	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
	if err != nil {
		log.Fatalf("unable start the server: %s\n", err.Error())
	}
}
