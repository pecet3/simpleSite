package model

import (
	"log"
)

type Post struct {
	Id      uint64 `json:"id"`
	Content uint64 `json:"content"`
	UserId  uint64 `json:"user_id"`
}

func (p *Post) CreatePost() (*Post, error) {
	query := "insert into posts (content, user_id) values ($1, $2)"

	_, err := db.Query(query, p.Content, p.UserId)
	if err != nil {
		return nil, err

	}

	log.Printf("user_id: %v has created the new record in posts", p.UserId)

	return p, nil
}
