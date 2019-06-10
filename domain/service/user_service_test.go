package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/interface/middleware"
)

var (
	userHashPassword, _ = middleware.GenerateHash("LoginID")

	createdUserMock = &model.User{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		LoginID:  "LoginID",
		Password: userHashPassword,
		Name:     "name",
		School:   "school",
		Role:     1,
	}
)

type UserRepositoryMock struct {
}

func (m *UserRepositoryMock) Create(user *model.User) error {
	return nil
}

func (m *UserRepositoryMock) FindByLoginID(LoginID string) (model.User, error) {
	createdUser := createdUserMock
	return *createdUser, nil
}

func TestUserService_Show(t *testing.T) {
	target := NewUserService(&UserRepositoryMock{})
	want := createdUserMock
	got, err := target.Show("loginID")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
