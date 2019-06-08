package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
)

var (
	problemMock = &model.Problem{
		Title:   "title",
		Content: "content",
		UserID:  1,
	}

	createdProblemMock = &model.Problem{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		Title:   "title",
		Content: "content",
		UserID:  1,
	}
)

type TProblemRepositoryMock struct {
}

func (m *TProblemRepositoryMock) CreateProblem(problem *model.Problem) (*model.Problem, error) {
	createdProblem := createdProblemMock
	return createdProblem, nil
}

func TestProblemService_CreateProblem(t *testing.T) {
	target := NewTProblemService(&TProblemRepositoryMock{})
	want := createdProblemMock
	got, err := target.CreateProblem(problemMock)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
