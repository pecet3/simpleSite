package model

import (
	"log"
)

type Post struct {
	Id      uint64 `json:"id"`
	Content string `json:"content"`
	UserId  uint64 `json:"user_id"`
}
type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
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

func (u *User) RegisterUser() (*User, error) {
	query := "insert into users (name, password) values ($1, $2);"

	_, err := db.Query(query, u.Name, u.Password)
	if err != nil {
		return nil, err

	}

	log.Printf("user of id: %v has created created", u.Id)

	return u, nil
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

func GetUserByName(name string) ([]User, error) {
	var users []User

	query := "select * from users where name = $1;"

	rows, err := db.Query(query, name)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
