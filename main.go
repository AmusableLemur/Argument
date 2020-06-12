package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main() {
    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
        Extensions: []string{".html"},
        Layout: "layout",
    }))

    m.Get("/", func(r render.Render) {
        r.HTML(200, "index")
    })

    m.Run()
}
