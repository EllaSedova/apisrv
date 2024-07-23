package rpc

import (
	"apisrv/pkg/newsportal"
)

func newNews(in *newsportal.News) *News {
	if in == nil {
		return nil
	}

	return &News{
		ID:          in.ID,
		Title:       in.Title,
		Foreword:    in.Foreword,
		Content:     in.Content,
		Author:      in.Author,
		PublishedAt: in.PublishedAt,
		Tags:        newTags(in.Tags),
		Category:    *newCategory(in.Category),
	}

}

func newNewsSummary(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}

	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		Foreword:    in.Foreword,
		Author:      in.Author,
		PublishedAt: in.PublishedAt,
		Tags:        newTags(in.Tags),
		Category:    *newCategory(in.Category),
	}

}

func newCategory(in *newsportal.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		Alias:       in.Alias,
	}
}

func newCategories(in []newsportal.Category) (out []Category) {
	for i := range in {
		out = append(out, *newCategory(&in[i]))
	}
	return
}

func newTag(in *newsportal.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func newTags(in []newsportal.Tag) (out []Tag) {
	for i := range in {
		out = append(out, *newTag(&in[i]))
	}
	return
}
