package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Any("/*", echoRequest)

	e.Logger.Fatal(e.Start(":" + port()))
}

func echoRequest(c echo.Context) error {
	response := make(map[string]interface{})
	response["uri"] = uri(c.Request())
	response["headers"] = headers(c.Request())

	if shouldRenderBodyAsString() {
		response["body"] = string(body(c.Request()))
	} else {
		response["body"] = body(c.Request())
	}

	return c.JSON(http.StatusOK, response)
}

func uri(req *http.Request) string {
	return req.RequestURI
}

func headers(req *http.Request) map[string]string {
	headers := make(map[string]string)
	for k, v := range req.Header {
		headers[k] = v[0]
	}
	return headers
}

func body(req *http.Request) []byte {
	bodyArr := make([]byte, req.ContentLength)
	_, _ = req.Body.Read(bodyArr)

	return bodyArr
}

func shouldRenderBodyAsString() bool {
	val, present := os.LookupEnv("BODY_AS_STRING")
	if present {
		boolVal, err := strconv.ParseBool(val)
		if err == nil {
			return boolVal
		}
	}

	return true
}

func port() (port string) {
	port = "80"

	val, present := os.LookupEnv("PORT")
	if present {
		tPort, err := strconv.Atoi(val)
		if err == nil {
			port = strconv.Itoa(tPort)
		}
	}

	return
}
