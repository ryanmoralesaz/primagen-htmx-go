package main

import (
	"io"
	"log"
	"os"
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

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John Doe", "jd@gmail.com"),
			newContact("Jane Doe", "cd@gmail.com"),
		},
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

	data := newData()

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "index", data)
	})

	e.Logger.Fatal(e.Start(":4200"))
}
