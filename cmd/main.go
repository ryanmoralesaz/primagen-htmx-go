package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}
func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("aoeu", "jd@gmail.com"),
			newContact("Jane Doe", "cd@gmail.com"),
			newContact("Ivan Doe", "id@gmail.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	page := newPage()

	// data := newData()

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if page.Data.hasEmail(email) {
			page.Form.Errors["email"] = "Email already exists"
			return c.Render(422, "index", page)
		}
		// if page.Data.hasEmail(email) {
		// 	formData := newFormData()
		// 	formData.Values["name"] = name

		// 	formData.Errors["name"] = "Name already exists"
		// 	formData.Values["email"] = email
		// 	formData.Errors["email"] = "Email already exists"
		// 	return c.Render(422, "form", formData)
		// }
		page.Data.Contacts = append(page.Data.Contacts, newContact(name, email))
		return c.Render(http.StatusOK, "display", page)
	})
	e.Logger.Fatal(e.Start(":4200"))
}
