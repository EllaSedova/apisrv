//nolint:dupl
package vt

import (
	"time"

	"apisrv/pkg/db"
)

type Category struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,max=255"`
	OrderNumber *int   `json:"orderNumber"`
	Alias       string `json:"alias" validate:"required,alias,max=255"`
	StatusID    int    `json:"statusId" validate:"required,status"`

	Status *Status `json:"status"`
}

func (c *Category) ToDB() *db.Category {
	if c == nil {
		return nil
	}

	category := &db.Category{
		ID:          c.ID,
		Title:       c.Title,
		OrderNumber: c.OrderNumber,
		Alias:       c.Alias,
		StatusID:    c.StatusID,
	}

	return category
}

type CategorySearch struct {
	ID          *int    `json:"id"`
	Title       *string `json:"title"`
	OrderNumber *int    `json:"orderNumber"`
	Alias       *string `json:"alias"`
	StatusID    *int    `json:"statusId"`
	IDs         []int   `json:"ids"`
	NotID       *int    `json:"notId"`
}

func (cs *CategorySearch) ToDB() *db.CategorySearch {
	if cs == nil {
		return nil
	}

	return &db.CategorySearch{
		ID:          cs.ID,
		TitleILike:  cs.Title,
		OrderNumber: cs.OrderNumber,
		Alias:       cs.Alias,
		StatusID:    cs.StatusID,
		IDs:         cs.IDs,
		NotID:       cs.NotID,
	}
}

type CategorySummary struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	OrderNumber *int   `json:"orderNumber"`
	Alias       string `json:"alias"`

	Status *Status `json:"status"`
}

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,max=255"`
	CategoryID  int       `json:"categoryId" validate:"required"`
	Foreword    string    `json:"foreword" validate:"required,max=1024"`
	Content     *string   `json:"content"`
	TagIDs      []int     `json:"tagIds" validate:"required"`
	Author      string    `json:"author" validate:"required,max=64"`
	PublishedAt time.Time `json:"publishedAt" validate:"required"`
	StatusID    int       `json:"statusId" validate:"required,status"`

	Category *CategorySummary `json:"category"`
	Status   *Status          `json:"status"`
}

func (n *News) ToDB() *db.News {
	if n == nil {
		return nil
	}

	news := &db.News{
		ID:          n.ID,
		Title:       n.Title,
		CategoryID:  n.CategoryID,
		Foreword:    n.Foreword,
		Content:     n.Content,
		TagIDs:      n.TagIDs,
		Author:      n.Author,
		PublishedAt: n.PublishedAt,
		StatusID:    n.StatusID,
	}

	return news
}

type NewsSearch struct {
	ID          *int       `json:"id"`
	Title       *string    `json:"title"`
	CategoryID  *int       `json:"categoryId"`
	Foreword    *string    `json:"foreword"`
	TagIDs      *int       `json:"tagIds"`
	Author      *string    `json:"author"`
	PublishedAt *time.Time `json:"publishedAt"`
	StatusID    *int       `json:"statusId"`
	IDs         []int      `json:"ids"`
}

func (ns *NewsSearch) ToDB() *db.NewsSearch {
	if ns == nil {
		return nil
	}

	return &db.NewsSearch{
		ID:            ns.ID,
		TitleILike:    ns.Title,
		CategoryID:    ns.CategoryID,
		ForewordILike: ns.Foreword,
		TagIDILike:    ns.TagIDs,
		AuthorILike:   ns.Author,
		PublishedAt:   ns.PublishedAt,
		StatusID:      ns.StatusID,
		IDs:           ns.IDs,
	}
}

type NewsSummary struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	CategoryID  int              `json:"categoryId"`
	TagIDs      []int            `json:"tagIds"`
	Author      string           `json:"author"`
	PublishedAt time.Time        `json:"publishedAt"`
	Tags        []TagSummary     `json:"tags"`
	Category    *CategorySummary `json:"category"`
	Status      *Status          `json:"status"`
}

type Tag struct {
	ID       int    `json:"id"`
	Title    string `json:"title" validate:"required,max=128"`
	StatusID int    `json:"statusId" validate:"required,status"`

	Status *Status `json:"status"`
}

func (t *Tag) ToDB() *db.Tag {
	if t == nil {
		return nil
	}

	tag := &db.Tag{
		ID:       t.ID,
		Title:    t.Title,
		StatusID: t.StatusID,
	}

	return tag
}

type TagSearch struct {
	ID       *int    `json:"id"`
	Title    *string `json:"title"`
	StatusID *int    `json:"statusId"`
	IDs      []int   `json:"ids"`
}

func (ts *TagSearch) ToDB() *db.TagSearch {
	if ts == nil {
		return nil
	}

	return &db.TagSearch{
		ID:         ts.ID,
		TitleILike: ts.Title,
		StatusID:   ts.StatusID,
		IDs:        ts.IDs,
	}
}

type TagSummary struct {
	ID    int    `json:"id"`
	Title string `json:"title"`

	Status *Status `json:"status"`
}
