package main

import (
	"fmt"

	mysql "github.com/AltaProject/AltaSocialMedia/infrastructure/database"

	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	fmt.Println("==== STARTING PROGRAM ====")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
