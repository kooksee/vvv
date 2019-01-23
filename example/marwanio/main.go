package main

import (
	"github.com/gopherjs/vecty"
	"github.com/kooksee/vvv/example/marwanio/components"
	"github.com/kooksee/vvv/example/marwanio/stores/blogposts"
)

func main() {
	must(blogposts.Fetch())
	vecty.SetTitle("Marwan - Software Engineer")
	vecty.RenderBody(&components.Body{})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
