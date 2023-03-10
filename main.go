package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
type Template struct {
	templates *template.Template
} 

type myProject struct{
	Title string
	Description string
	StartDate string
	EndDate string
	Duration string
	technologies []string
	
}

var kumpulanProject = []myProject{
	{
		Title: "Ya Allah Berilah Hamba Kemudahan",
		Description: "Aaamiiinnn",
		StartDate: "2021-02-10",
		EndDate: "2021-02-20",
		Duration: "3 Minggu",
		
	},
	{
		Title: "Ya Allah Berilah Hamba Kemudahan",
		Description: "Aaamiiinnn",
		StartDate: "2021-02-10",
		EndDate: "2021-02-20",
		Duration: "4 Bulan",
		
	},
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
	e.GET("/delete/:id", deletePost)
	
	
	// fmt.Println("Helo")

	

	e.Logger.Fatal(e.Start("localhost:5008"))
}

// func haloDunia(c echo.Context) error {
// 	return c.String(http.StatusOK, "Halo Dunia!")
// }
func home(c echo.Context) error {

	project := map[string]interface{}{
		"projects": kumpulanProject,
	}

	return c.Render(http.StatusOK, "index.html", project)
}


func contactMe(c echo.Context) error {
	return c.Render(http.StatusOK, "contactForm.html", nil)
}
func addProject(c echo.Context) error {
	return c.Render(http.StatusOK, "addProject.html", nil)
}


func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) 

	var ProjectDetail = myProject{}

	for i, data := range kumpulanProject {
		if id == i {
			ProjectDetail = myProject {
				Title: data.Title,
				Description: data.Description,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				
			}
		}

	}

	dataProject := map[string]interface{}{ 
		"ProjectKuy": ProjectDetail,
	}

	return c.Render(http.StatusOK, "projectDetail.html", dataProject)
	
}

func postProject(c echo.Context) error  {
	title := c.FormValue("title")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("Description")

	date1, _ := time.Parse("2006-01-02", startDate)
	date2, _ := time.Parse("2006-01-02", endDate)

	durationDays := date2.Sub(date1) / (24)
	var durasi = ""

	if (durationDays > 365){
       durasi = durationDays.String() + "tahun"
    } else if (durationDays > 30){
		durasi = durationDays.String() + "bulan"
	} else if (durationDays > 7){
		durasi = durationDays.String() + "minggu"
	} else {
		durasi = durationDays.String() + "hari"
	}

	var projectBaru = myProject{
		Title: title,
		Description: description,
		StartDate: startDate,
		EndDate: endDate,
		Duration: durasi,
	}

	kumpulanProject = append(kumpulanProject, projectBaru)

	println(title)
	println(startDate)
	println(endDate)
	
	// durasi := startDate.Sub(endDate)
	// fmt.Println(durasi)
	
	println(description)

	//mencari tahu jenis tipe data
	// fmt.Println(reflect.TypeOf(startDate))
	// fmt.Println(reflect.TypeOf(node))
	// fmt.Println(reflect.TypeOf(react))
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deletePost(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))

	kumpulanProject = append(kumpulanProject[:id], kumpulanProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}