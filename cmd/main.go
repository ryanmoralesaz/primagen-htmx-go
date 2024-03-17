package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	template *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		template: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

func main() {
	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	e := echo.New()
	e.Use(middleware.Logger())

	// Define the absolute path for the count.txt file
	absPath := filepath.Join("./", "count.txt")

	// Use the absolute path in loadCount
	count, err := loadCount(absPath)
	if err != nil {
		log.Printf("Error loading count: %v", err)
		count = 0
	}

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", Count{Count: count})
	})

	e.POST("/count", func(c echo.Context) error {
		count++
		// Use the absolute path in saveCount
		if err := saveCount(absPath, count); err != nil {
			log.Printf("Error saving count: %v", err)
			return err
		}
		return c.Render(200, "count", Count{Count: count})
	})

	e.Logger.Fatal(e.Start(":4200"))
}
