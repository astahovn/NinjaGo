package post

import (
  "time"
  "github.com/astahovn/ninja/lib/db"
)

type Post struct {
  ID int
  Title string
  ShortText string
  FullText string
  DateAdded time.Time
}

func FetchLast() []Post {
  var posts []Post
  db.GetInstance().Order("id desc").Limit(3).Find(&posts)
  return posts
}