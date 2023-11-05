package model

import (
	"log"
)

type Post struct {
	Id      uint64 `json:"id"`
	Content string `json:"content"`
	UserId  uint64 `json:"user_id"`
}

func (p *Post) CreatePost() (*Post, error) {
	query := "insert into posts (content, user_id) values ($1, $2);"

	_, err := db.Query(query, p.Content, p.UserId)
	if err != nil {
		return nil, err

	}

	log.Printf("user_id: %v has created a new record in posts", p.UserId)

	return p, nil
}

func GetAllPosts() ([]Post, error) {
	var posts []Post

	query := "select * from posts;"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Content, &post.UserId)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	
	return posts, nil
}
