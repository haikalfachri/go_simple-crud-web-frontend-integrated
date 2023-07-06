package routes

import (
	"biodata/controllers"
	"biodata/utils"
	"net/http"

	"github.com/labstack/echo"
)

func SetUpRoutes(e *echo.Echo) {
	biodataCtrl := controllers.InitBiodataContoller()
	e.Renderer = utils.NewRenderer("./public/views/*.html", true)

	// e.GET("/index", func(c echo.Context) error {
	// 	data := ""
	// 	return c.Render(http.StatusOK, "index.html", data)
	// })

	e.GET("/home", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", "")
	})

	e.GET("/form", func(c echo.Context) error {
		return c.Render(http.StatusOK, "form.html", "")
	})

	e.GET("/biodata", biodataCtrl.GetAll)
	e.GET("/biodata/:id", biodataCtrl.GetById)
	e.POST("/biodata", biodataCtrl.Create)
	e.DELETE("/biodata/:id", biodataCtrl.Delete)

	e.Static("/css", "./public/views/css/style.css")
	e.Static("/icon", "./public/assets/icon")
	e.Static("/picture", "./public/assets/picture")
}