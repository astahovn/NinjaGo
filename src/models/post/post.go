package post

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

func FetchLast() []Post {
  db := lib.InitDb()
  defer db.Close()

  var posts []Post
  db.Order("id desc").Limit(3).Find(&posts)
  return posts
}