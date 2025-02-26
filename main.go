package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	// Note that this must be the last statement of the main function
	e.Logger.Fatal(e.Start(":8888"))
}
