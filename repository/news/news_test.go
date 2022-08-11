package repository

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetNewsListRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewNewsRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name           string
		mockRepository *NewsRepository
		scope          int
		category       string
		mock           func()
		want           []entity.NewsListOutput
		wantErr        bool
	}{
		{
			name:           "OK",
			mockRepository: mockRepository,
			scope:          3,
			category:       "Howak",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"n.id", "n.title", "n.slug_url", "n.cover_image", "additional_images", "n.nsfw",
						"c.id", "c.name",
						"u.id", "u.name", "u.picture",
						"nc.upvote", "nc.downvote", "comment", "nc.view",
						"n.created_at"},
				).
					AddRow("922c7afd-643e-4e44-ab51-c80dc137674a", "News Title", "news-title", "Cover", `[{"image": "ABC"}]`, false,
						"d414197c-0fa0-46c1-ac29-69c4cdc0ed11", "Howak",
						"e65d7793-bcc6-467c-88b1-9636ee745f45", "Name", "Picture",
						10, 23, 2, 1000,
						time.Time{})

				mock.ExpectPrepare("SELECT (.+) FROM news").ExpectQuery().WillReturnRows(rows)
			},
			want: []entity.NewsListOutput{
				{
					Id:               uuid.MustParse("922c7afd-643e-4e44-ab51-c80dc137674a"),
					Title:            "News Title",
					SlugUrl:          "news-title",
					CoverImage:       func(val string) *string { return &val }("Cover"),
					AdditionalImages: []string{"ABC"},
					CreatedAt:        time.Time{},
					Nsfw:             false,
					Category: entity.NewsCategory{
						Id:   uuid.MustParse("d414197c-0fa0-46c1-ac29-69c4cdc0ed11"),
						Name: "Howak",
					},
					Author: entity.NewsAuthor{
						Id:      uuid.MustParse("e65d7793-bcc6-467c-88b1-9636ee745f45"),
						Name:    "Name",
						Picture: func(val string) *string { return &val }("Picture"),
					},
					Counter: entity.NewsCounter{
						Upvote:   10,
						Downvote: 23,
						Comment:  2,
						View:     1000,
					},
				},
			},
		},
		{
			name:           "Error",
			mockRepository: mockRepository,
			scope:          3,
			category:       "Howak",
			mock: func() {
				mock.ExpectPrepare("SELECT (.+) FROM news").ExpectQuery().WillReturnError(errors.New("get news list error"))
			},
			wantErr: true,
		},
		{
			name:           "Invalid Query",
			mockRepository: mockRepository,
			scope:          3,
			category:       "Howak",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"n.id", "n.title", "n.slug_url", "n.cover_image", "additional_images", "n.nsfw",
						"c.id", "c.name",
						"u.id", "u.name", "u.picture",
						"nc.upvote", "nc.downvote", "comment", "nc.view",
						"n.created_at"},
				).
					AddRow("922c7afd-643e-4e44-ab51-c80dc137674a", "News Title", "news-title", "Cover", `[{"image": "ABC"}]`, false,
						"d414197c-0fa0-46c1-ac29-69c4cdc0ed11", "Howak",
						"e65d7793-bcc6-467c-88b1-9636ee745f45", "Name", "Picture",
						10, 23, 2, 1000,
						time.Time{})

				mock.ExpectPrepare("SELECTTTTTTTTTT (.+) FROM news").ExpectQuery().WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.mockRepository.GetNewsListRepository(tt.scope, tt.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNewsDetailRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewNewsRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name           string
		mockRepository *NewsRepository
		newsId         string
		mock           func()
		want           entity.NewsDetail
		wantErr        bool
	}{
		{
			name:           "OK",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"n.id", "n.title", "n.slug_url", "n.cover_image", "additional_images", "n.nsfw",
						"c.id", "c.name",
						"u.id", "u.name", "u.picture",
						"nc.upvote", "nc.downvote", "comment", "nc.view",
						"n.created_at", "n.isi"},
				).
					AddRow("922c7afd-643e-4e44-ab51-c80dc137674a", "News Title", "news-title", "Cover", `[{"image": "ABC"}]`, false,
						"d414197c-0fa0-46c1-ac29-69c4cdc0ed11", "Howak",
						"e65d7793-bcc6-467c-88b1-9636ee745f45", "Name", "Picture",
						10, 23, 2, 1000,
						time.Time{}, "Lorem Ipsum")

				mock.ExpectBegin()
				mock.ExpectPrepare("SELECT (.+) FROM news").ExpectQuery().WillReturnRows(rows)
				mock.ExpectPrepare("UPDATE news_counter (.+)").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: entity.NewsDetail{
				Id:               uuid.MustParse("922c7afd-643e-4e44-ab51-c80dc137674a"),
				Title:            "News Title",
				SlugUrl:          "news-title",
				CoverImage:       func(val string) *string { return &val }("Cover"),
				AdditionalImages: []string{"ABC"},
				CreatedAt:        time.Time{},
				Nsfw:             false,
				Category: entity.NewsCategory{
					Id:   uuid.MustParse("d414197c-0fa0-46c1-ac29-69c4cdc0ed11"),
					Name: "Howak",
				},
				Author: entity.NewsAuthor{
					Id:      uuid.MustParse("e65d7793-bcc6-467c-88b1-9636ee745f45"),
					Name:    "Name",
					Picture: func(val string) *string { return &val }("Picture"),
				},
				Counter: entity.NewsCounter{
					Upvote:   10,
					Downvote: 23,
					Comment:  2,
					View:     1000,
				},
				Content: "Lorem Ipsum",
			},
		},
		{
			name:           "Failed at First Transaction",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectPrepare("SELECT (.+) FROM news").ExpectQuery().WillReturnError(errors.New("error message"))
			},
			wantErr: true,
		},
		{
			name:           "Failed at Second Transaction",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"n.id", "n.title", "n.slug_url", "n.cover_image", "additional_images", "n.nsfw",
						"c.id", "c.name",
						"u.id", "u.name", "u.picture",
						"nc.upvote", "nc.downvote", "comment", "nc.view",
						"n.created_at", "n.isi"},
				).
					AddRow("922c7afd-643e-4e44-ab51-c80dc137674a", "News Title", "news-title", "Cover", `[{"image": "ABC"}]`, false,
						"d414197c-0fa0-46c1-ac29-69c4cdc0ed11", "Howak",
						"e65d7793-bcc6-467c-88b1-9636ee745f45", "Name", "Picture",
						10, 23, 2, 1000,
						time.Time{}, "Lorem Ipsum")

				mock.ExpectBegin()
				mock.ExpectPrepare("SELECT (.+) FROM news").ExpectQuery().WillReturnRows(rows)
				mock.ExpectPrepare("UPDATE news_counter (.+)").ExpectExec().WillReturnError(errors.New("update failed"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.mockRepository.GetNewsDetailRepository(tt.newsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNewsUserRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewNewsRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name           string
		mockRepository *NewsRepository
		newsId         string
		mock           func()
		want           string
		wantErr        bool
	}{
		{
			name:           "OK",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"users_id"},
				).
					AddRow("b46a18e7-6fae-4a0d-8179-317b856dd1ac")

				mock.ExpectPrepare("SELECT (.+) FROM news").
					ExpectQuery().
					WithArgs("922c7afd-643e-4e44-ab51-c80dc137674a").
					WillReturnRows(rows)
			},
			want: "b46a18e7-6fae-4a0d-8179-317b856dd1ac",
		},
		{
			name:           "Error",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				mock.ExpectPrepare("SELECT (.+) FROM news").
					ExpectQuery().
					WithArgs("922c7afd-643e-4e44-ab51-c80dc137674a").
					WillReturnError(errors.New("error message"))
			},
			wantErr: true,
		},
		{
			name:           "Invalid Query",
			mockRepository: mockRepository,
			newsId:         "922c7afd-643e-4e44-ab51-c80dc137674a",
			mock: func() {
				mock.ExpectPrepare("SELECTTTTTTTTT (.+) FROM news").
					ExpectQuery().
					WithArgs("922c7afd-643e-4e44-ab51-c80dc137674a").
					WillReturnError(errors.New("error message"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.mockRepository.GetNewsUserRepository(tt.newsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddNewNewsRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewNewsRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name           string
		mockRepository *NewsRepository
		newsData       entity.NewsInsert
		mock           func()
		want           string
		wantErr        bool
	}{
		{
			name:           "OK",
			mockRepository: mockRepository,
			newsData: entity.NewsInsert{
				Id:         uuid.MustParse("2fe834c0-b5f7-4cb8-8cc3-b0a2d8cdba8a"),
				UserId:     uuid.MustParse("624def9b-1314-4a7d-a98a-0952eea8cdc2"),
				CategoryId: uuid.MustParse("701dec5c-2984-490c-950a-622ba9c95bce"),
				NewsInputRequest: entity.NewsInputRequest{
					Title:            "Judul Berita",
					SlugUrl:          "judul-berita",
					CoverImage:       func(val string) *string { return &val }("Cover Image"),
					AdditionalImages: []string{"ABC"},
					CreatedAt:        time.Time{},
					Nsfw:             false,
					Content:          "Lorem",
					Category:         "Howak",
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectPrepare("INSERT INTO news").
					ExpectExec().
					WithArgs("2fe834c0-b5f7-4cb8-8cc3-b0a2d8cdba8a",
						"624def9b-1314-4a7d-a98a-0952eea8cdc2",
						"701dec5c-2984-490c-950a-622ba9c95bce",
						"Judul Berita", "Lorem", "judul-berita", "Cover Image", false, AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectPrepare("INSERT INTO news_additional_images").ExpectExec().WithArgs("2fe834c0-b5f7-4cb8-8cc3-b0a2d8cdba8a", "ABC").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectPrepare("INSERT INTO news_counter").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: "b46a18e7-6fae-4a0d-8179-317b856dd1ac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.mockRepository.AddNewNewsRepository(tt.newsData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
