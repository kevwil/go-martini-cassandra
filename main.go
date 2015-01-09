package main

import (
	"github.com/go-martini/martini"
	"github.com/kevwil/go-martini-cassandra/db"
	"github.com/kevwil/go-martini-cassandra/models"
)

var repo db.Repository

func main() {
	repo := db.NewRepository("127.0.0.1", "mykeyspace")
	repo.Begin()
	defer repo.Finish()

	mc := martini.Classic()
	// mc.Get("/", GetHome)
	// mc.Get("/favicon.ico", func() (int,string) {
	//     return 404, "Not Found"
	// })
	mc.Get("/blog", GetBlogListing)
	mc.Get("/blog/:title", GetBlogSingle)

	mc.Run()
}

// GetBlogListing = retrieve list of blog posts
func GetBlogListing() []*models.Post {
	return repo.GetAllPosts()
}

// GetBlogSingle = get a single blog post by title
func GetBlogSingle(params martini.Params) models.Post {
	return repo.GetSinglePost(params["title"])
}
