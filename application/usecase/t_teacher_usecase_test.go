package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/interface/middleware"
)

var (
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
		Role:     0,	
	}
)

type TTeacherServiceMock struct {
}

func (m *TTeacherServiceMock) CreateTeacher(teacher *model.User) (*model.User, error) {
	createdTeacher := createdTeacherMock
	return createdTeacher, nil
}

func TestTTeacherUsecase_CreateTeacher(t *testing.T) {
	target := NewTTeacherUsecase(&TTeacherServiceMock{})
	want := createdTeacherMock
	got, err := target.CreateTeacher("LoginID", "name", "school")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
