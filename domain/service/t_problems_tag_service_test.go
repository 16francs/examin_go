package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
)

var (
	problemsTagMock = &model.ProblemsTag{
		ProblemID: 1,
		TagID:     1,
	}

	createdProblemsTagMock = &model.ProblemsTag{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		ProblemID: 1,
		TagID:     1,
	}
)

type TProblemsTagRepositoryMock struct {
}

func (m *TProblemsTagRepositoryMock) CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error) {
	createdProblemsTag := createdProblemsTagMock
	return createdProblemsTag, nil
}

func TestProblemsTagService_CreateProblemsTag(t *testing.T) {
	target := NewTProblemsTagService(&TProblemsTagRepositoryMock{})
	want := createdProblemsTagMock
	got, err := target.CreateProblemsTag(problemsTagMock)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
