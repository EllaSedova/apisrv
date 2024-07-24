package newsportal

import (
	"apisrv/pkg/db"
)

func newNews(in *db.News) *News {
	if in == nil {
		return nil
	}
	return &News{
		News:     in,
		Category: newCategory(in.Category),
	}
}

func NewNewsList(in []db.News) (out []News) {
	for i := range in {
		out = append(out, *newNews(&in[i]))
	}

	return
}

func newCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}
	return &Category{
		Category: in,
	}
}

func newCategories(in []db.Category) (out []Category) {
	for i := range in {
		out = append(out, *newCategory(&in[i]))
	}
	return
}

func newTag(in *db.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{
		Tag: in,
	}
}

func newTags(in []db.Tag) (out []Tag) {
	for i := range in {
		out = append(out, *newTag(&in[i]))
	}
	return
}
