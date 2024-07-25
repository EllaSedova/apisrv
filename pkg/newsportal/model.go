package newsportal

import (
	"apisrv/pkg/db"
)

type News struct {
	*db.News
	Category *Category
	Tags     []Tag
	Author   *Author
}

type Category struct {
	*db.Category
}

type Tag struct {
	*db.Tag
}

type Author struct {
	*db.Author
}

type NewsList []News

func (nn NewsList) TagIDs() []int {
	// собираем все уникальные tagID

	tagIDMap := make(map[int]struct{})
	for _, summary := range nn {
		for _, tagId := range summary.TagIDs {
			tagIDMap[tagId] = struct{}{}
		}
	}

	// заполняем карту уникальных tagId
	var uniqueTagIDs []int
	for tagID := range tagIDMap {
		uniqueTagIDs = append(uniqueTagIDs, tagID)
	}
	return uniqueTagIDs
}
