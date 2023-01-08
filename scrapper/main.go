package main

import (
	"github.com/labstack/echo"
	"os"
	"strings"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

const fileName string = "jobs.csv"

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(CleanString(c.FormValue("term")))
	//fmt.Println(term)
	Scrapper(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1234"))
	//Scrapper("total-salary")
}
