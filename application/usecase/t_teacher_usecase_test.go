package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
)

type TTeacherServiceMock struct {
}

func (m *TTeacherServiceMock) CreateTeacher(loginID, name, school string) (*model.User, error) {
	teacher := &model.User{
		Base:	model.Base{
			ID: 1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LoginID: loginID,
		Password: "password",
		Name: name,
		School: school,
		Role: 0,
	}
	return teacher, nil
}

func TestTTeacherUsecase_CreateTeacher(t *testing.T) {
	target := NewTTeacherUsecase(&TTeacherServiceMock{})
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
