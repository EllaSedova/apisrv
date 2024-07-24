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
		Author:      in.Author,
		PublishedAt: in.PublishedAt,
		StatusID:    in.StatusID,

		Category: NewCategorySummary(in.Category),
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
		TagIDs:      in.TagIDs,
		Author:      in.Author,
		PublishedAt: in.PublishedAt,

		Category: NewCategorySummary(in.Category),
		Status:   NewStatus(in.StatusID),
	}
}
func NewNewsSummaryFromNewsportal(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}

	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		CategoryID:  in.CategoryID,
		TagIDs:      in.TagIDs,
		Author:      in.Author,
		PublishedAt: in.PublishedAt,
		Tags:        NewTagSummaryList(NewTagSummaryListFromNewsportal(in.Tags)),
		Category:    NewCategorySummary(in.Category.Category),
		Status:      NewStatus(in.StatusID),
	}
}

func NewTagSummaryListFromNewsportal(in []newsportal.Tag) (out []db.Tag) {
	for i := range in {
		out = append(out, *in[i].Tag)
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

func NewTagSummaryList(in []db.Tag) (out []TagSummary) {
	for i := range in {
		out = append(out, *NewTagSummary(&in[i]))
	}
	return
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
