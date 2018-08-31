package main

import (
	"fmt"
	"os"

	brandHTTP "github.com/hardyantz/go-clean-arch/src/brand/delivery/http"
	brandRepo "github.com/hardyantz/go-clean-arch/src/brand/repository"
	brandUseCase "github.com/hardyantz/go-clean-arch/src/brand/usecase"
	"github.com/hardyantz/go-clean-arch/conn"
	_ "github.com/arangodb/go-driver"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	db, err := conn.ArangoHandler()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e := echo.New()

	brand := brandRepo.NewArangoBrandRepository(db)
	brandUCase := brandUseCase.NewBrandUsecase(brand)
	brandHTTP.NewBrandHTTPHandler(e, brandUCase)

	e.Start(":" + os.Getenv("SITE_PORT"))
}
