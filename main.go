package main

import (
	"html"
	"myapp/app/models"
	"time"

	m "github.com/go-martini/martini"
	c "github.com/gocql/gocql"
)

var session *c.Session

func main() {
	cluster := c.NewCluster("127.0.0.1")
	cluster.Keyspace = "mykeyspace"
	cluster.Consistency = c.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	mc := m.Classic()
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
	var posts []*models.Post
	var id uint
	var title, tags, content string
	var date time.Time
	iter := session.Query(`SELECT id,title,tags,content,date FROM posts`).Iter()
	for iter.Scan(&id, &title, &tags, &content, &date) {
		newPost := models.Post{Id: id, Title: html.EscapeString(title), Tags: tags, Content: content, Date: date}
		posts = append(posts, &newPost)
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}
	return posts
}

// GetBlogSingle = get a single blog post by title
func GetBlogSingle(params m.Params) models.Post {
	var id uint
	var title, tags, content string
	var date time.Time
	err := session.Query(`SELECT id,title,tags,content,date FROM posts WHERE title = ? LIMIT 1`, params["title"]).Scan(&id, &title, &tags, &content, &date)
	if err != nil {
		panic(err)
	}
	newPost := models.Post{Id: id, Title: html.EscapeString(title), Tags: tags, Content: content, Date: date}
	return newPost
}
