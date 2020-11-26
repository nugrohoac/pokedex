package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	middleware_echo "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/nugrohoac/pokedex/docs"
	delivery "github.com/nugrohoac/pokedex/internal/http"
	"github.com/nugrohoac/pokedex/internal/middleware"
	"github.com/nugrohoac/pokedex/internal/mysql"
	"github.com/nugrohoac/pokedex/pokemon"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsnMysql := os.Getenv("DSN")
	mysqlDB, err := sql.Open("mysql", dsnMysql)
	if err != nil {
		logrus.Fatal("FAILED CONNECT TO MYSQL", err.Error())
	}

	if mysqlDB != nil {
		mysqlDB.SetConnMaxLifetime(time.Duration(5) * time.Second)
		mysqlDB.SetMaxIdleConns(3)
		mysqlDB.SetConnMaxLifetime(5)
	}

	pokemonDelivery := mysql.NewPokemonRepository(mysqlDB)
	pokemonService := pokemon.NewPokemonService(pokemonDelivery)

	e := echo.New()
	e.Use(middleware_echo.Logger())
	e.Use(middleware_echo.Recover())
	e.Use(middleware.ErrHandler)
	e.Use(middleware_echo.CORSWithConfig(middleware_echo.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))

	timeOutString := os.Getenv("TIMEOUT")
	timeOutInt, err := strconv.Atoi(timeOutString)
	if err != nil {
		logrus.Fatal("FAILED LOAD TIMEOUT", err.Error())

	}

	// swagger start
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// swagger end

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	delivery.NewPokemonDelivery(e, pokemonService, time.Duration(timeOutInt)*time.Second, *validator.New())

	e.Logger.Fatal(e.Start(":4000"))
}
