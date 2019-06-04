package service

import (
	"time"
	"reflect"
	"testing"

	"github.com/16francs/examin_go/domain/model"
)

type TTeacherRepositoryMock struct {
}

func (m *TTeacherRepositoryMock) CreateTeacher(teacher *model.User) (*model.User, error) {
	createdTeacher := &model.User{
		Base:	model.Base{
			ID: 1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LoginID: "LoginID",
		Password: "password",
		Name: "name",
		School: "school",
		Role: 0,
	}
	return createdTeacher, nil
}

func TestTeacherService_CreateTeacher(t *testing.T) {
	target := NewTTeacherService(&TTeacherRepositoryMock{})
	want := &model.User{
		Base: model.Base{
			ID: 1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LoginID: "LoginID",
		Password: "password",
		Name: "name",
		School: "school",
		Role: 0,
	}
	got, err := target.CreateTeacher("LoginID", "name", "school")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
