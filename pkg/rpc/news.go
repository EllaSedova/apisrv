package rpc

import (
	"context"

	"github.com/vmkteam/zenrpc/v2"

	"apisrv/pkg/newsportal"
)

type NewsService struct {
	zenrpc.Service
	m *newsportal.Manager
}

func NewNewsService(m *newsportal.Manager) *NewsService {
	return &NewsService{
		m: m,
	}
}

// NewsByID получение новости по id
func (rs NewsService) NewsByID(ctx context.Context, id int) (*News, error) {
	news, err := rs.m.NewsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return newNews(news), nil
}

// Categories получение всех категорий
func (rs NewsService) Categories(ctx context.Context) ([]Category, error) {
	categories, err := rs.m.Categories(ctx)
	if err != nil {
		return nil, err
	}

	return newCategories(categories), err
}

// Tags получение всех тегов
func (rs NewsService) Tags(ctx context.Context) ([]Tag, error) {
	tags, err := rs.m.Tags(ctx)
	if err != nil {
		return nil, err
	}

	return newTags(tags), err
}

// NewsWithFilters получение новости с фильтрами
func (rs NewsService) NewsWithFilters(ctx context.Context, categoryID, tagID, page, pageSize *int) ([]NewsSummary, error) {
	newsResponse, err := rs.m.News(ctx, categoryID, tagID, page, pageSize)
	if err != nil {
		return nil, err
	}

	var newNewsList []NewsSummary
	for _, summary := range newsResponse {
		newNews := newNewsSummary(&summary)
		newNewsList = append(newNewsList, *newNews)
	}

	return newNewsList, nil
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (rs NewsService) NewsCountWithFilters(ctx context.Context, categoryID, tagID *int) (*int, error) {
	count, err := rs.m.NewsCount(ctx, categoryID, tagID)

	return count, err
}
