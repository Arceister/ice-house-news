package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestGetAllNewsCategoryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewCategoriesRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name    string
		app     *CategoriesRepository
		mock    func()
		want    []entity.Categories
		wantErr bool
	}{
		{
			name: "OK",
			app:  app,
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "name"},
				).
					AddRow("6fae13cb-e8a4-46c4-b412-a0d41662e024", "International").
					AddRow("28596a94-0ea8-4fd3-ad10-89df980decf3", "Sports")

				mock.ExpectPrepare("SELECT (.+) FROM categories").
					ExpectQuery().
					WillReturnRows(rows)
			},
			want: []entity.Categories{
				{
					Id:   uuid.MustParse("6fae13cb-e8a4-46c4-b412-a0d41662e024"),
					Name: "International",
				},
				{
					Id:   uuid.MustParse("28596a94-0ea8-4fd3-ad10-89df980decf3"),
					Name: "Sports",
				},
			},
		},
		{
			name: "Invalid SQL Query",
			app:  app,
			mock: func() {
				sqlmock.NewRows(
					[]string{"id", "name"},
				).
					AddRow("6fae13cb-e8a4-46c4-b412-a0d41662e024", "International").
					AddRow("28596a94-0ea8-4fd3-ad10-89df980decf3", "Sports")

				mock.ExpectPrepare("SELECT (.+) FROM categories").
					ExpectQuery().
					WillReturnError(errors.New("get all categories error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.app.GetAllNewsCategoryRepository()
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

func TestCreateCategoryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewCategoriesRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name    string
		app     *CategoriesRepository
		request entity.Categories
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			app:  app,
			request: entity.Categories{
				Id:   uuid.MustParse("28596a94-0ea8-4fd3-ad10-89df980decf3"),
				Name: "Howak",
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO categories").
					ExpectExec().
					WithArgs("28596a94-0ea8-4fd3-ad10-89df980decf3", "Howak").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Empty ID",
			app:  app,
			request: entity.Categories{
				Id:   uuid.Nil,
				Name: "Howak",
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO categories").
					ExpectExec().
					WithArgs("", "Howak").
					WillReturnError(errors.New("user not created"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.app.CreateCategoryRepository(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestCreateAndReturnCategoryRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewCategoriesRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name    string
		app     *CategoriesRepository
		request entity.Categories
		mock    func()
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "OK",
			app:  app,
			request: entity.Categories{
				Id:   uuid.MustParse("d0ff38ec-e438-4adb-9332-4a324d20a872"),
				Name: "Internal",
			},
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id"},
				).AddRow("d0ff38ec-e438-4adb-9332-4a324d20a872")

				mock.ExpectPrepare("INSERT INTO categories").
					ExpectQuery().
					WithArgs("d0ff38ec-e438-4adb-9332-4a324d20a872", "Internal").
					WillReturnRows(rows)
			},
			want: uuid.MustParse("d0ff38ec-e438-4adb-9332-4a324d20a872"),
		},
		{
			name: "Invalid Query",
			app:  app,
			request: entity.Categories{
				Id:   uuid.MustParse("d0ff38ec-e438-4adb-9332-4a324d20a872"),
				Name: "Internal",
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTOOOOOOOO categories").
					ExpectQuery().
					WithArgs("d0ff38ec-e438-4adb-9332-4a324d20a872", "Internal").
					WillReturnError(errors.New("invalid query"))
			},
			want:    uuid.Nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.app.CreateAndReturnCategoryRepository(tt.request)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCategoryByNameRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewCategoriesRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name     string
		app      *CategoriesRepository
		category string
		mock     func()
		want     entity.Categories
		wantErr  bool
	}{
		{
			name:     "OK",
			app:      app,
			category: "International",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "name"},
				).AddRow("d0ff38ec-e438-4adb-9332-4a324d20a872", "International")

				mock.ExpectPrepare("SELECT (.+) FROM categories WHERE name").
					ExpectQuery().
					WithArgs("International").
					WillReturnRows(rows)
			},
			want: entity.Categories{
				Id:   uuid.MustParse("d0ff38ec-e438-4adb-9332-4a324d20a872"),
				Name: "International",
			},
		},
		{
			name:     "Invalid Query",
			app:      app,
			category: "Intern",
			mock: func() {
				mock.ExpectPrepare("INSERT INTOOOOOOOO categories").
					ExpectQuery().
					WithArgs("Internal").
					WillReturnError(errors.New("invalid query"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.app.GetCategoryByNameRepository(tt.category)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
