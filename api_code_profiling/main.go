package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	pprof.Register(e)

	e.GET("/check", func(c echo.Context) error {
		word := c.QueryParam("w")

		isPalindrome := checkPalindrome(word)

		return c.JSON(http.StatusOK, echo.Map{
			"word":   word,
			"result": isPalindrome,
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func checkPalindrome(word string) bool {
	reversed := ""
	splitted := []byte(word)

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(splitted[i])
	}

	return reversed == word
}
