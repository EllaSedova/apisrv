package rpc

import "time"

type NewsSummary struct {
	ID          int       `json:"newsId"`
	Title       string    `json:"title"`
	Category    Category  `json:"categoryId"`
	Foreword    string    `json:"foreword"`
	Tags        []Tag     `json:"tags,array"`
	Author      Author    `json:"author"`
	PublishedAt time.Time `json:"publishedAt"`
}
type News struct {
	ID          int       `json:"newsId"`
	Title       string    `json:"title"`
	Category    Category  `json:"categoryId"`
	Foreword    string    `json:"foreword"`
	Content     *string   `json:"content"`
	Tags        []Tag     `json:"tags,array"`
	Author      Author    `json:"author"`
	PublishedAt time.Time `json:"publishedAt"`
}

type Category struct {
	ID          int    `json:"categoryId"`
	Title       string `json:"title"`
	OrderNumber *int   `json:"orderNumber"`
	Alias       string `json:"alias"`
}

type Tag struct {
	ID    int    `json:"tagId"`
	Title string `json:"title"`
}

type Author struct {
	ID    int    `json:"authorId"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
