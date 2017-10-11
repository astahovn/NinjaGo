package models

import (
  "time"
  "lib"
)

type Post struct {
  ID int
  Title string
  ShortText string
  FullText string
  DateAdded time.Time
}

func FetchPosts() []Post {
  db := lib.InitDb()
  defer db.Close()

  var posts []Post
  db.Find(&posts)
  return posts
}