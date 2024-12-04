package main

import (
	"github.com/go-martini/martini"
	// "net/http"
  	// "github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main(){
	m := martini.Classic()
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"status": "world", "result": "OK"})
	  })
	m.RunOnAddr(":7001")
}