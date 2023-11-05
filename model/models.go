package model

import (
	"log"
)

type Post struct {
	Id      uint64 `json:"id"`
	Content uint64 `json:"content"`
	UserId  uint64 `json:"user_id"`
}

func CreatePost(content string, userId uint64) error {
	query := "insert into posts (content, user_id) values ($1, $2)"

	_, err := db.Query(query, content, userId)
	if err != nil {
		return err

	}

	log.Printf("user_id: %v has created the new record in posts", userId)

	return nil
}
