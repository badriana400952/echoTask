package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/hehe", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Static("/public", "public")
	e.Static("/view", "view")
	e.GET("rumah", halloWord)
	e.GET("hallo", hallo)
	e.GET("/", home)
	e.GET("myblog", myBlog)
	e.GET("kontak", kontak)
	e.GET("detail/:id", detail)
	e.Logger.Fatal(e.Start("localhost:3000"))
}
func halloWord(c echo.Context) error {
	return c.String(http.StatusOK, "view/index.html")
}
func hallo(c echo.Context) error {
	return c.String(http.StatusOK, "hallo")
}

func home(c echo.Context) error {
	var template, err = template.ParseFiles("view/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func myBlog(c echo.Context) error {
	var template, err = template.ParseFiles("view/MyBlog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}
func kontak(c echo.Context) error {
	var template, err = template.ParseFiles("view/kontak.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func detail(c echo.Context) error {
	Nim, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Nim":     Nim,
		"Title":   "lorem ipsum dolor sit",
		"Content": "is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum",
	}

	var template, err = template.ParseFiles("view/detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), data)

}
