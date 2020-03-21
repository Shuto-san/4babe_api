package repository_test

import (
	"regexp"
	"testing"

	"github.com/Shuto-san/4babe_api/testutil"

	"github.com/Shuto-san/4babe_api/interface/repository"

	"github.com/stretchr/testify/assert"

	"github.com/Shuto-san/4babe_api/domain/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func setup(t *testing.T) (r repository.QuestionRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := testutil.NewDBMock(t)

	r = repository.NewQuestionRepository(db)

	return r, mock, func() {
		db.Close()
	}
}

func TestQuestionRepository_FindAll(t *testing.T) {
	t.Helper()
	r, mock, teardown := setup(t)
	defer teardown()

	cases := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, u []*model.Question, err error)
	}{
		"正常なSQLのとき": {
			arrange: func(t *testing.T) {
				u := model.Question{ID: 1, Name: "manato", Age: "20"}
				c := model.CreditCard{ID: 1, QuestionID: 1, Number: "111111"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(u.ID, u.Name, u.Age))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "question_id", "number"}).AddRow(c.ID, c.QuestionID, c.Number))
			},
			assert: func(t *testing.T, u []*model.Question, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1, len(u))
				assert.NotEmpty(t, u[0].CreditCards)
			},
		},
		"エラーのSQLのとき": {
			arrange: func(t *testing.T) {
				u := model.Question{ID: 1, Name: "manato", Age: "20"}
				c := model.CreditCard{ID: 1, QuestionID: 1, Number: "111111"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(u.ID, u.Name, u.Age))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).WillReturnRows(sqlmock.NewRows([]string{"id", "question_id", "number"}).AddRow(c.ID, c.QuestionID, c.Number))
			},
			assert: func(t *testing.T, u []*model.Question, err error) {
				//assert.NotEmpty(t, err)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			tt.arrange(t)

			u, err := r.FindAll([]*model.Question{})

			tt.assert(t, u, err)
		})
	}
}
