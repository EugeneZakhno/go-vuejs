package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func main() {
	e := echo.New()
	e.Static("/", "ui/dist")
	e.File("/", "ui/dist/index.html")

	// We're going to face CORS issue since we are passing data between different apps.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8888", "http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	p := Person{
		Name:  "Nikola",
		Age:   37,
		Email: "nikola@tesla.genius",
	}
	e.GET("/person", func(c echo.Context) error {
		return c.JSON(202, p)
	})

	// Update Person's name
	e.POST("/person", func(c echo.Context) error {
		// Get the request
		r := c.Request()
		// Read the body
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("error in POST", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Body Request"})
		}
		n := PostPersonBody{
			Name: "default",
		}
		// equivalent of JSON.parse() in GO
		// By default Go passes arguments by value, meaning it creates a copy of the value, and a new pointer is created.
		// json.Unmarshall requires a reference (a pointer) to PostPersonBody and will update it internally.
		err = json.Unmarshal(b, &n)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}
		// Debug purpose
		fmt.Println(n.Name)
		// Update local instance (db...)
		p.Name = n.Name
		return c.JSON(http.StatusAccepted, n)
	})
	e.Logger.Fatal(e.Start(":8888"))
}

type Person struct {
	Name  string `json:"name"`
	Age   uint8  `json:"age"`
	Email string `json:"email"`
}
