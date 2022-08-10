package repository

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

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
