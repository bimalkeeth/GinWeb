package controllers

import (
	"GinWeb/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RegisterRoutes() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employee/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timeOff, ok := models.TimesOff[id]
		if !ok {
			c.String(http.StatusNotFound, "Data Not found")
			return
		}
		c.HTML(http.StatusOK, "vacation_overview.html", map[string]interface{}{
			"TimesOff": timeOff,
		})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	admin.POST("/employee/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			pto, err := strconv.ParseFloat(c.PostForm("pto"), 32)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			startDate, err := time.Parse("2006-01-02", c.PostForm("startDate"))
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			var emp models.Employee
			emp.ID = 42
			emp.TotalPTO = float32(pto)
			emp.StartDate = startDate
			emp.FirstName = c.PostForm("firstName")
			emp.LastName = c.PostForm("lastName")
			emp.Status = "Active"
			models.Employees["42"] = emp
			c.Redirect(http.StatusMovedPermanently, "/admin/employees/42")
		}
	})

	r.Static("/public", "./public")

	return r
}
