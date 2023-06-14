package main

import (
	"context"
	"fmt"

	"go-latihan1/connection"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

// struct => hampir mirip seperti object di javascript
type Blog struct {
	ID         int
	Nama       string
	Stardate   string
	Enddate    string
	Durasi     string
	Deskripsi  string
	Nodejs     bool
	React      bool
	Nextjs     bool
	Typescript bool
	Img        string
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()
	// e.GET("/hehe", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Static("/public", "public")
	e.Static("/view", "view")
	e.GET("rumah", halloWord)
	e.GET("hallo", hallo)
	e.GET("/", home)
	e.GET("tambahBlog", myBlog)
	e.GET("kontak", kontak)
	e.GET("/blogview", blogview)

	e.GET("/detail/:id", detail)

	e.POST("tambahBlog", tambahBlog)
	e.POST("/deleteBlog/:id", deleteBlog)
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
	// jalankan response template
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
	// jalankan response template

	return template.Execute(c.Response(), nil)
}

func detail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var Details = Blog{}

	ppp := connection.Conn.QueryRow(context.Background(), "SELECT * FROM datablog WHERE id=$1", id).Scan(
		&Details.ID, &Details.Nama, &Details.Stardate, &Details.Enddate, &Details.Deskripsi, &Details.Nodejs, &Details.React, &Details.Nextjs, &Details.Typescript, &Details.Img)

	if ppp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": ppp.Error()})
	}
	data := map[string]interface{}{
		"Blog": Details,
	}

	var tmpl, err = template.ParseFiles("view/detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}

func blogview(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, nama, stardate, enddate,durasi, deskripsi, nodejs, react, nextjs, typescript, img FROM coba5")

	var hasil []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Nama, &each.Stardate, &each.Enddate, &each.Durasi, &each.Deskripsi, &each.Nodejs, &each.React, &each.Nextjs, &each.Typescript, &each.Img)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		hasil = append(hasil, each)
	}
	var temp, err = template.ParseFiles("view/blogview.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": hasil}

	return temp.Execute(c.Response(), blogs)
}

func tambahBlog(c echo.Context) error {
	nama := c.FormValue("nama")
	stardate := c.FormValue("start")
	enddate := c.FormValue("endDate")
	deskripsi := c.FormValue("deskripsi")
	durasi := c.FormValue("deskripsi")
	// durasi := calculateDuration(stardate, endDate)
	var nodejs bool
	if c.FormValue("nodeJss") == "yes" {
		nodejs = true
	}

	var react bool
	if c.FormValue("react") == "yes" {
		react = true
	}

	var nextjs bool
	if c.FormValue("next") == "yes" {
		nextjs = true
	}

	var typescript bool
	if c.FormValue("typeScript") == "yes" {
		typescript = true
	}
	img := c.FormValue("imageUploud")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO coba5 (nama, stardate, enddate,durasi, deskripsi, nodejs, react, nextjs, typescript, img) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", nama, stardate, enddate, durasi, deskripsi, nodejs, react, nextjs, typescript, img)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blogview")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM coba5 WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blogview")
}
