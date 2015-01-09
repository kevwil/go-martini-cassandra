package db

import (
	"github.com/gocql/gocql"
	"github.com/kevwil/go-martini-cassandra/models"
	"html"
	"time"
)

type Repository interface {
	Begin()
	Finish()
	GetAllPosts() []*models.Post
	GetSinglePost(title string) models.Post
}

type repository struct {
	cc   *gocql.ClusterConfig
	Sess *gocql.Session
}

func NewRepository(clusterAddr string, keyspace string) Repository {
	cluster := gocql.NewCluster(clusterAddr)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.One
	return &repository{cc: cluster}
}

func (r *repository) Begin() {
	session, err := r.cc.CreateSession()
	if err != nil {
		panic(err)
	}
	r.Sess = session
}

func (r *repository) Finish() {
	r.Sess.Close()
}

func (r *repository) GetAllPosts() []*models.Post {
	var posts []*models.Post
	var id uint
	var title, tags, content string
	var date time.Time
	iter := r.Sess.Query(`SELECT id,title,tags,content,date FROM posts`).Iter()
	for iter.Scan(&id, &title, &tags, &content, &date) {
		newPost := models.Post{Id: id, Title: html.EscapeString(title), Tags: tags, Content: content, Date: date}
		posts = append(posts, &newPost)
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}
	return posts
}

func (r *repository) GetSinglePost(title string) models.Post {
	var id uint
	var mytitle, tags, content string
	var date time.Time
	err := r.Sess.Query(`SELECT id,title,tags,content,date FROM posts WHERE title = ? LIMIT 1`, title).Scan(&id, &mytitle, &tags, &content, &date)
	if err != nil {
		panic(err)
	}
	newPost := models.Post{Id: id, Title: html.EscapeString(mytitle), Tags: tags, Content: content, Date: date}
	return newPost
}
