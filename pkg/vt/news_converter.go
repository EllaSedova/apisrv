package vt

import (
	"apisrv/pkg/db"
	"apisrv/pkg/newsportal"
)

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
		TagIDs:      in.TagIDs,
		PublishedAt: in.PublishedAt,
		StatusID:    in.StatusID,
		AuthorID:    in.AuthorID,

		Category: NewCategorySummary(in.Category),
		Status:   NewStatus(in.StatusID),
		Author:   NewAuthorSummary(in.Author),
	}

	return news
}

func NewNewsSummary(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}

	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		CategoryID:  in.CategoryID,
		TagIDs:      in.TagIDs,
		PublishedAt: in.PublishedAt,
		Tags:        NewTagSummaryList(in.Tags),
		Category:    NewCategorySummary(in.Category.Category),
		Status:      NewStatus(in.StatusID),
	}
}

func NewTagSummaryList(in []newsportal.Tag) (out []TagSummary) {
	for i := range in {
		out = append(out, *NewTagSummary(in[i].Tag))
	}
	return
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
