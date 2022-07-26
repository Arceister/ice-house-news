package repository

import (
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
		name    string
		s       UsersRepository
		msgUUID string
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name:    "OK",
			s:       app,
			msgUUID: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "email", "password", "name", "bio", "web", "picture"},
				).
					AddRow("72908c48-b68c-4d67-ae74-d1305f84fc4d", "testemail@email.com", "123", "Jagad", "Bio", "Web", "Picture")

				mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").WillReturnRows(rows)
			},
			want: &entity.User{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetOneUserRepository(tt.msgUUID)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
				return
			}

			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Fatal(err)
				return
			}
		})
	}
}
