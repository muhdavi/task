package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	funcMap := template.FuncMap{
		"add": func(x int, y int) int {
			return x + y
		},
	}

	r.SetFuncMap(funcMap)
	r.StaticFile("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*.html")

	Routes(r)

	//r.Run(":3003") // developmen
	r.Run()          // deployment
}
