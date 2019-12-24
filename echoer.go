package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func headers(c echo.Context) error {

	response := make(map[string]string)
	for k, v := range c.Request().Header {
		response[k] = v[0]
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/*", headers)

	e.Logger.Fatal(e.Start(":" + port()))
}

func port() (port string) {
	port = "80"

	val, present := os.LookupEnv("PORT")
	if present {
		tPort, err := strconv.Atoi(val)
		if err != nil {
			port = strconv.Itoa(tPort)
		}
	}

	return
}
