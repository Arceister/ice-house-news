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

func TestGetOneUserRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewUsersRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name     string
		s        UsersRepository
		userUUID string
		mock     func()
		want     entity.User
		wantErr  bool
	}{
		{
			name:     "OK",
			s:        app,
			userUUID: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "email", "password", "name", "bio", "web", "picture"},
				).
					AddRow("72908c48-b68c-4d67-ae74-d1305f84fc4d", "testemail@email.com", "123", "Jagad", "Bio", "Web", "Picture")

				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").WillReturnRows(rows)
			},
			want: entity.User{
				Id:       uuid.MustParse("72908c48-b68c-4d67-ae74-d1305f84fc4d"),
				Email:    "testemail@email.com",
				Password: "123",
				Name:     "Name",
				Bio:      func(val string) *string { return &val }("Bio"),
				Web:      func(val string) *string { return &val }("Web"),
				Picture:  func(val string) *string { return &val }("Picture"),
			},
			wantErr: false,
		},
		{
			name:     "Not Found",
			s:        app,
			userUUID: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "email", "password", "name", "bio", "web", "picture"},
				)
				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetOneUserRepository(tt.userUUID)
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

func TestCreateUserRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	app := NewUsersRepository(
		lib.DB{
			DB: db,
		},
	)
	tests := []struct {
		name     string
		s        UsersRepository
		userUUID uuid.UUID
		request  entity.User
		mock     func()
		wantErr  bool
	}{
		{
			name:     "OK",
			s:        app,
			userUUID: uuid.MustParse("72908c48-b68c-4d67-ae74-d1305f84fc4d"),
			request: entity.User{
				Email:    "testemail@email.com",
				Password: "123",
				Name:     "Jagad",
				Bio:      func(val string) *string { return &val }("Bio"),
				Web:      func(val string) *string { return &val }("Web"),
				Picture:  func(val string) *string { return &val }("Picture"),
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO users").
					WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d", "testemail@email.com", "123", "Jagad", "Bio", "Web", "Picture").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:     "Empty Email",
			s:        app,
			userUUID: uuid.MustParse("72908c48-b68c-4d67-ae74-d1305f84fc4d"),
			request: entity.User{
				Email:    "",
				Password: "123",
				Name:     "Jagad",
				Bio:      func(val string) *string { return &val }("Bio"),
				Web:      func(val string) *string { return &val }("Web"),
				Picture:  func(val string) *string { return &val }("Picture"),
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO users").
					WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d", "", "123", "Jagad", "Bio", "Web", "Picture").
					WillReturnError(errors.New("empty email"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.s.CreateUserRepository(tt.userUUID, tt.request)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}
		})
	}
}
