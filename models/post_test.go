package models

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
    p := &Post{Id: 123, Title: "Worlds Greatest Post", Tags: "#testing #go", Content: "nothing to see here", Date: time.Now()}
    assert.NotNil(t, p)
}
