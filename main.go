package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type Template struct {
	templates *template.Template
} 

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main(){
	e := echo.New()

	e.Static("/public", "public")

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	// e.GET("/hello", haloDunia) //localhost:5008/hello
	e.GET("/", home) //localhost:5008/
	e.GET("/contactMe", contactMe)
	e.GET("/addProject", addProject)
	e.GET("/projectDetail/:id", projectDetail)
	e.POST("/postProject", postProject)
	
	
	// fmt.Println("Helo")


	e.Logger.Fatal(e.Start("localhost:5008"))
}

// func haloDunia(c echo.Context) error {
// 	return c.String(http.StatusOK, "Halo Dunia!")
// }
func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
func contactMe(c echo.Context) error {
	return c.Render(http.StatusOK, "contactForm.html", nil)
}
func addProject(c echo.Context) error {
	return c.Render(http.StatusOK, "addProject.html", nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) 

	data := map[string]interface{}{ 
		"Id":      id,
		"Title":   "Dumbways Mobile Apps",
		"Content": "Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam, molestiae numquam! Deleniti maiores expedita eaque deserunt quaerat! Dicta, eligendi debitis?",
	}

	return c.Render(http.StatusOK, "projectDetail.html", data)
	
}

func postProject(c echo.Context) error  {
	title := c.FormValue("title")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	node := c.FormValue("node")
	react := c.FormValue("react")
	next := c.FormValue("next")
	typescript := c.FormValue("typescript")
	Description := c.FormValue("Description")

	println(title)
	println(startDate)
	println(endDate)
	if node == "on"  {
		println("node=", node)
	} else {
		println("node = off")
	}
	if next == "on"  {
		println("next=", next)
	} else {
		println("next = off")
	}
	if react == "on"  {
		println("react=", react)
	} else {
		println("react = off")
	}
	if typescript == "on"  {
		println("typescript=", typescript)
	} else {
		println("typescript = off")
	}
	
	
	println(Description)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}