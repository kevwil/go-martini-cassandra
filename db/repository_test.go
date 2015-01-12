package db

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRepoStartStop(t *testing.T) {
    repo := NewRepository("127.0.0.1", "mykeyspace")
    assert.NotNil(t, repo)
    repo.Begin()
    assert.NotNil(t, repo.Sess)
    repo.Finish()
}

func TestGetAllPosts(t *testing.T) {
    repo := NewRepository("127.0.0.1", "mykeyspace")
    assert.NotNil(t, repo)
    repo.Begin()
    assert.NotNil(t, repo.Sess)

    posts := repo.GetAllPosts()
    assert.NotNil(t, posts)
    t.Log(posts)

    repo.Finish()
}
