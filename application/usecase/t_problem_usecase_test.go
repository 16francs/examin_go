package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
)

var (
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

	createdTagMock = &model.Tag{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		Content: "content",
	}

	createdProblemsTagMock = &model.ProblemsTag{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		ProblemID: createdProblemMock.Base.ID,
		TagID:     createdTagMock.Base.ID,
	}
)

type TProblemServiceMock struct {
}

func (m *TProblemServiceMock) CreateProblem(problem *model.Problem) (*model.Problem, error) {
	createdProblem := createdProblemMock
	return createdProblem, nil
}

type TProblemsTagServiceMock struct {
}

func (m *TProblemsTagServiceMock) CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error) {
	createdProblemsTag := createdProblemsTagMock
	return createdProblemsTag, nil
}

type TTagServiceMock struct {
}

func (m *TTagServiceMock) CreateTag(tag *model.Tag) (*model.Tag, error) {
	createdTag := createdTagMock
	return createdTag, nil
}

func TestTProblemUsecase_CreateProblem(t *testing.T) {
	target := NewTProblemUsecase(&TProblemServiceMock{}, &TProblemsTagServiceMock{}, &TTagServiceMock{})
	wantProblem := createdProblemMock
	wantTags := []*model.Tag{createdTagMock, createdTagMock, createdTagMock}
	gotProblem, gotTags, err := target.CreateProblem("title", "content", 1, []string{"aaa", "bbb", "ccc"})

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(gotProblem, wantProblem) {
		t.Fatalf("want %#v, but %#v", wantProblem, gotProblem)
	}

	if !reflect.DeepEqual(gotTags, wantTags) {
		t.Fatalf("want %#v, but %#v", wantTags, gotTags)
	}
}
