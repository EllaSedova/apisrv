package vt

import (
	"apisrv/pkg/db"
)

func NewAuthor(in *db.Author) *Author {
	if in == nil {
		return nil
	}

	author := &Author{
		ID:       in.ID,
		Name:     in.Name,
		Email:    in.Email,
		StatusID: in.StatusID,

		Status: NewStatus(in.StatusID),
	}

	return author
}

func NewAuthorSummary(in *db.Author) *AuthorSummary {
	if in == nil {
		return nil
	}

	return &AuthorSummary{
		ID:    in.ID,
		Name:  in.Name,
		Email: in.Email,

		Status: NewStatus(in.StatusID),
	}
}

func NewCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}

	category := &Category{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		Alias:       in.Alias,
		StatusID:    in.StatusID,

		Status: NewStatus(in.StatusID),
	}

	return category
}

func NewCategorySummary(in *db.Category) *CategorySummary {
	if in == nil {
		return nil
	}

	return &CategorySummary{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		Alias:       in.Alias,

		Status: NewStatus(in.StatusID),
	}
}

func NewNews(in *db.News) *News {
	if in == nil {
		return nil
	}

	news := &News{
		ID:          in.ID,
		Title:       in.Title,
		CategoryID:  in.CategoryID,
		Foreword:    in.Foreword,
		Content:     in.Content,
		AuthorID:    in.AuthorID,
		TagIDs:      in.TagIDs,
		PublishedAt: in.PublishedAt,
		StatusID:    in.StatusID,

		Category: NewCategorySummary(in.Category),
		Author:   NewAuthorSummary(in.Author),
		Status:   NewStatus(in.StatusID),
	}

	return news
}

func NewNewsSummary(in *db.News) *NewsSummary {
	if in == nil {
		return nil
	}

	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		CategoryID:  in.CategoryID,
		Foreword:    in.Foreword,
		Content:     in.Content,
		AuthorID:    in.AuthorID,
		PublishedAt: in.PublishedAt,

		Category: NewCategorySummary(in.Category),
		Author:   NewAuthorSummary(in.Author),
		Status:   NewStatus(in.StatusID),
	}
}

func NewTag(in *db.Tag) *Tag {
	if in == nil {
		return nil
	}

	tag := &Tag{
		ID:       in.ID,
		Title:    in.Title,
		StatusID: in.StatusID,

		Status: NewStatus(in.StatusID),
	}

	return tag
}

func NewTagSummary(in *db.Tag) *TagSummary {
	if in == nil {
		return nil
	}

	return &TagSummary{
		ID:    in.ID,
		Title: in.Title,

		Status: NewStatus(in.StatusID),
	}
}
