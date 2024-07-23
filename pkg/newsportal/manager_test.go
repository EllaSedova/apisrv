package newsportal

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"

	"apisrv/pkg/db"
)

func ptrs(r string) *string { return &r }

var dbc *pg.DB
var nr db.NewsRepo
var nm *Manager
var realNews15 News
var realNews16 News
var realNews17 News
var wrongNews News

func TestMain(m *testing.M) {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "newsportal",
	}

	dbc = pg.Connect(opts)
	err := dbc.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	nr = db.NewNewsRepo(dbc)
	nm = NewManager(nr)

	realNews15 = News{
		News: &db.News{ID: 15,
			Title:       "AНовость5",
			CategoryID:  1,
			Foreword:    "Преамбула",
			Content:     ptrs("Контент"),
			TagIDs:      []int{1, 2, 3},
			Author:      "Автор",
			PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
			StatusID:    1,
			Category: &db.Category{
				ID:          1,
				Title:       "рр",
				OrderNumber: nil,
				Alias:       "к",
				StatusID:    1,
			},
		},
		Category: &Category{&db.Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
		},
		Tags: []Tag{
			{
				&db.Tag{
					ID:       1,
					Title:    "заголовок1",
					StatusID: 1,
				},
			},
			{
				&db.Tag{
					ID:       2,
					Title:    "заголовок2",
					StatusID: 1,
				},
			},
			{
				&db.Tag{
					ID:       3,
					Title:    "заголовок3",
					StatusID: 1,
				},
			},
		},
	}

	realNews16 = News{
		News: &db.News{ID: 16,
			Title:       "Новость10",
			CategoryID:  1,
			Foreword:    "Преамбула",
			Content:     ptrs("Контент"),
			TagIDs:      []int{1, 2},
			Author:      "Автор",
			PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
			StatusID:    1,
			Category: &db.Category{
				ID:          1,
				Title:       "рр",
				OrderNumber: nil,
				Alias:       "к",
				StatusID:    1,
			},
		},
		Category: &Category{&db.Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
		},
		Tags: []Tag{
			{
				&db.Tag{
					ID:       1,
					Title:    "заголовок1",
					StatusID: 1,
				},
			},
			{
				&db.Tag{
					ID:       2,
					Title:    "заголовок2",
					StatusID: 1,
				},
			},
		},
	}
	realNews17 = News{
		News: &db.News{ID: 17,
			Title:       "BНовость17",
			CategoryID:  1,
			Foreword:    "Преамбула",
			Content:     ptrs("Контент"),
			TagIDs:      []int{1, 2},
			Author:      "Автор",
			PublishedAt: time.Date(2024, time.July, 17, 18, 25, 28, 10745000, time.Local),
			StatusID:    1,
			Category: &db.Category{
				ID:          1,
				Title:       "рр",
				OrderNumber: nil,
				Alias:       "к",
				StatusID:    1,
			},
		},
		Category: &Category{&db.Category{
			ID:          1,
			Title:       "рр",
			OrderNumber: nil,
			Alias:       "к",
			StatusID:    1,
		},
		},
		Tags: []Tag{
			{
				&db.Tag{
					ID:       1,
					Title:    "заголовок1",
					StatusID: 1,
				},
			},
			{
				&db.Tag{
					ID:       2,
					Title:    "заголовок2",
					StatusID: 1,
				},
			},
		},
	}
	os.Exit(m.Run())
}
func TestNewsByID(t *testing.T) {
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    News
		wantErr assert.ErrorAssertionFunc
	}{
		{name: "valid args",
			args: args{
				ctx: context.Background(),
				ID:  15,
			},
			want:    realNews15,
			wantErr: assert.NoError,
		},
		{name: "invalid args",
			args: args{
				ctx: context.Background(),
				ID:  -15,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nm.NewsByID(tt.args.ctx, tt.args.ID)
			if !tt.wantErr(t, err, fmt.Sprintf("NewsByID(%v, %v)", tt.args.ctx, tt.args.ID)) {
				return
			}
			if tt.name == "invalid args" {
				assert.Nil(t, got)
			} else {
				assert.Equalf(t, &tt.want, got, "NewsByID(%v, %v)", tt.args.ctx, tt.args.ID)
			}
		})
	}
}

func TestFillTags(t *testing.T) {
	// Инициализация тестовых данных
	newsListTrue := []News{
		{News: &db.News{ID: 1, TagIDs: []int{1, 2}}},
		{News: &db.News{ID: 2, TagIDs: []int{2, 3}}},
	}
	newsListWrong := []News{
		{News: &db.News{ID: 1, TagIDs: []int{}}},
		{News: &db.News{ID: 2, TagIDs: []int{2, 13}}},
	}
	// Вызов метода
	err := nm.FillTags(context.Background(), newsListTrue)
	assert.NoError(t, err)

	// Проверка результатов
	trueTags := map[int][]Tag{
		1: {{&db.Tag{ID: 1, Title: "заголовок1", StatusID: 1}}, {&db.Tag{ID: 2, Title: "заголовок2", StatusID: 1}}},
		2: {{&db.Tag{ID: 2, Title: "заголовок2", StatusID: 1}}, {&db.Tag{ID: 3, Title: "заголовок3", StatusID: 1}}},
	}
	err = nm.FillTags(context.Background(), newsListWrong)
	assert.NoError(t, err)
	// Проверка результатов
	wrongTags := map[int][]Tag{
		1: nil,
		2: {{&db.Tag{ID: 2, Title: "заголовок2", StatusID: 1}}},
	}
	for _, news := range newsListTrue {
		assert.Equal(t, trueTags[news.ID], news.Tags)
	}
	for _, news := range newsListWrong {
		assert.Equal(t, wrongTags[news.ID], news.Tags)
	}
}

func TestNews(t *testing.T) {
	type args struct {
		ctx        context.Context
		categoryID *int
		tagID      *int
		page       *int
		pageSize   *int
	}
	tests := []struct {
		name    string
		args    args
		want    []News
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid args",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(1),
				tagID:      ptri(3),
				page:       nil,
				pageSize:   nil,
			},
			want:    []News{realNews15},
			wantErr: assert.NoError,
		},
		{
			name: "valid args return list",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(1),
				tagID:      ptri(1),
				page:       nil,
				pageSize:   nil,
			},
			want:    []News{realNews17, realNews16},
			wantErr: assert.NoError,
		},
		{
			name: "valid args return list with pagination",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(1),
				tagID:      ptri(1),
				page:       nil,
				pageSize:   ptri(10),
			},
			want:    []News{realNews17, realNews16, realNews15},
			wantErr: assert.NoError,
		},
		{
			name: "invalid args.categoryID",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(111),
				tagID:      ptri(3),
				page:       nil,
				pageSize:   nil,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nm.News(tt.args.ctx, tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)
			if !tt.wantErr(t, err, fmt.Sprintf("News(%v, %v, %v, %v, %v)", tt.args.ctx, tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)) {
				return
			}
			for i := range got {
				assert.Equal(t, tt.want[i].Tags, got[i].Tags)
				assert.Equal(t, tt.want[i].Category, got[i].Category)
			}
			assert.Equalf(t, tt.want, got, "News(%v, %v, %v, %v, %v)", tt.args.ctx, tt.args.categoryID, tt.args.tagID, tt.args.page, tt.args.pageSize)
		})
	}
}

func TestNewsCount(t *testing.T) {
	type args struct {
		ctx        context.Context
		categoryID *int
		tagID      *int
	}
	tests := []struct {
		name    string
		args    args
		want    *int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid args",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(1),
				tagID:      ptri(1),
			},
			want:    ptri(3),
			wantErr: assert.NoError,
		},
		{
			name: "invalid args",
			args: args{
				ctx:        context.Background(),
				categoryID: ptri(1),
				tagID:      ptri(133),
			},
			want:    ptri(0),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nm.NewsCount(tt.args.ctx, tt.args.categoryID, tt.args.tagID)
			if !tt.wantErr(t, err, fmt.Sprintf("NewsCount(%v, %v, %v)", tt.args.ctx, tt.args.categoryID, tt.args.tagID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewsCount(%v, %v, %v)", tt.args.ctx, tt.args.categoryID, tt.args.tagID)
		})
	}
}

func TestTagsByIDs(t *testing.T) {
	type args struct {
		ctx    context.Context
		tagIDs []int
	}
	tests := []struct {
		name    string
		args    args
		want    []Tag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid args",
			args: args{
				ctx:    context.Background(),
				tagIDs: []int{1, 2},
			},
			want:    []Tag{{&db.Tag{ID: 1, Title: "заголовок1", StatusID: 1}}, {&db.Tag{ID: 2, Title: "заголовок2", StatusID: 1}}},
			wantErr: assert.NoError,
		},
		{
			name: "invalid args",
			args: args{
				ctx:    context.Background(),
				tagIDs: []int{1, 13},
			},
			want:    []Tag{{&db.Tag{ID: 1, Title: "заголовок1", StatusID: 1}}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nm.TagsByIDs(tt.args.ctx, tt.args.tagIDs)
			if !tt.wantErr(t, err, fmt.Sprintf("TagsByIDs(%v, %v)", tt.args.ctx, tt.args.tagIDs)) {
				return
			}
			assert.Equalf(t, tt.want, got, "TagsByIDs(%v, %v)", tt.args.ctx, tt.args.tagIDs)
		})
	}
}
