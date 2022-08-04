package repository

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetCommentOnNewsRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewCommentRepository(
		lib.DB{
			DB: db,
		},
	)

	timePlaceholder := time.Now()

	tests := []struct {
		name    string
		app     repository.ICommentRepository
		newsId  string
		mock    func()
		want    []entity.Comment
		wantErr bool
	}{
		{
			name:   "OK",
			app:    mockRepository,
			newsId: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"nc.id", "nc.description", "u.id", "u.name", "u.picture", "nc.created_at"},
				).
					AddRow("ec0a3406-faca-48f0-8e6f-09df1275708d", "Comment 1", "4ece409a-866d-4cec-8ea8-996cd04a4a37", "User 1", "Picture 1", timePlaceholder).
					AddRow("89fdc5ce-c1a2-4351-b5ba-dd83e7cf6836", "Comment 2", "c60eb25a-d32f-48eb-83cd-a942191e7793", "User 2", "Picture 2", timePlaceholder)

				mock.ExpectPrepare("SELECT (.+) FROM (.+) JOIN (.+)").ExpectQuery().WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").WillReturnRows(rows)
			},
			want: []entity.Comment{
				{
					Id:          uuid.MustParse("ec0a3406-faca-48f0-8e6f-09df1275708d"),
					Description: "Comment 1",
					User: entity.Commentator{
						Id:      uuid.MustParse("4ece409a-866d-4cec-8ea8-996cd04a4a37"),
						Name:    "User 1",
						Picture: "Picture 1",
					},
					CreatedAt: timePlaceholder,
				},
				{
					Id:          uuid.MustParse("89fdc5ce-c1a2-4351-b5ba-dd83e7cf6836"),
					Description: "Comment 2",
					User: entity.Commentator{
						Id:      uuid.MustParse("c60eb25a-d32f-48eb-83cd-a942191e7793"),
						Name:    "User 2",
						Picture: "Picture 2",
					},
					CreatedAt: timePlaceholder,
				},
			},
		},
		{
			name:   "Error",
			app:    mockRepository,
			newsId: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				mock.ExpectPrepare("SELECT (.+) FROM (.+) JOIN (.+)").
					ExpectQuery().
					WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name:   "Wrong Query",
			app:    mockRepository,
			newsId: "72908c48-b68c-4d67-ae74-d1305f84fc4d",
			mock: func() {
				mock.ExpectPrepare("SELECTTTTTTT (.+) FROM (.+) JOIN (.+)").
					ExpectQuery().
					WithArgs("72908c48-b68c-4d67-ae74-d1305f84fc4d").
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.app.GetCommentsOnNewsRepository("72908c48-b68c-4d67-ae74-d1305f84fc4d")
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

func TestInsertCommentRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockRepository := NewCommentRepository(
		lib.DB{
			DB: db,
		},
	)

	tests := []struct {
		name    string
		app     repository.ICommentRepository
		request entity.CommentInsert
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			app:  mockRepository,
			request: entity.CommentInsert{
				Id:     uuid.MustParse("eb13e525-2e6b-42ea-832a-021bf39932cb"),
				UserId: uuid.MustParse("c6b1b9f9-92e5-4638-806f-27fbb94546a2"),
				NewsId: uuid.MustParse("3f441da6-4c57-4acb-869f-78469671c0fe"),
				CommentInsertRequest: entity.CommentInsertRequest{
					Description: "Comment",
				},
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO news_comment").
					ExpectExec().
					WithArgs("eb13e525-2e6b-42ea-832a-021bf39932cb",
						"3f441da6-4c57-4acb-869f-78469671c0fe",
						"c6b1b9f9-92e5-4638-806f-27fbb94546a2",
						"Comment",
						AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Empty Row",
			app:  mockRepository,
			request: entity.CommentInsert{
				Id:     uuid.MustParse("eb13e525-2e6b-42ea-832a-021bf39932cb"),
				UserId: uuid.MustParse("c6b1b9f9-92e5-4638-806f-27fbb94546a2"),
				NewsId: uuid.MustParse("3f441da6-4c57-4acb-869f-78469671c0fe"),
				CommentInsertRequest: entity.CommentInsertRequest{
					Description: "Comment",
				},
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO news_comment").
					ExpectExec().
					WithArgs().
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.app.InsertCommentRepository(tt.request)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}
		})
	}
}
