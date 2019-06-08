package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/interface/middleware"
)

var (
	teacherMock = &model.User{
		LoginID:  "LoginID",
		Password: "LoginID",
		Name:     "name",
		School:   "school",
	}

	hashPassword, _ = middleware.GenerateHash("LoginID")

	createdTeacherMock = &model.User{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LoginID:  "LoginID",
		Password: hashPassword,
		Name:     "name",
		School:   "school",
		Role:     1,
	}
)

type TTeacherRepositoryMock struct {
}

func (m *TTeacherRepositoryMock) CreateTeacher(teacher *model.User) (*model.User, error) {
	createdTeacher := createdTeacherMock
	return createdTeacher, nil
}

func TestTeacherService_CreateTeacher(t *testing.T) {
	target := NewTTeacherService(&TTeacherRepositoryMock{})
	want := createdTeacherMock
	got, err := target.CreateTeacher(teacherMock)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
