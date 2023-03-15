package main

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Template)
	e.Logger.Fatal(e.Start(":8000"))
}

func Template(c echo.Context) error{
	return c.JSON(http.StatusOK, gin.H{
        "msg": "--success--",
    })
}