package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/kevwil/go-martini-cassandra/db"
	"github.com/kevwil/go-martini-cassandra/models"
)

var repo db.Repository

func main() {
	repo = *db.NewRepository("127.0.0.1", "mykeyspace")
	repo.Begin()
	defer repo.Finish()

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Extensions: []string{".html"},
	}))
	m.Get("/", GetHome)
	m.Get("/favicon.ico", func() (int,string) {
		return 404, "Not Found"
	})
	m.Get("/blog", GetBlogListing)
	m.Get("/blog/:title", GetBlogSingle)

	m.Run()
}

// GetHome = render home page
func GetHome(r render.Render) {
	r.HTML(200, "index", map[string]interface{}{"PageTitle": "Home"})
}

type Listing struct {
	PageTitle string
	Posts []models.Post
}

// GetBlogListing = retrieve list of blog posts
func GetBlogListing(r render.Render) {
	r.HTML(200, "listing", Listing{"Blog",repo.GetAllPosts()})
}

type PostOutput struct {
	PageTitle string
	Post models.Post
}

// GetBlogSingle = get a single blog post by title
func GetBlogSingle(params martini.Params, r render.Render) {
	r.HTML(200, "single", PostOutput{params["title"],repo.GetSinglePost(params["title"])})
}
