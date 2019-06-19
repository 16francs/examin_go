package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/16francs/examin_go/domain/model"
)

var (
	tagMock = &model.Tag{
		Content: "tag",
	}

	createdTagMock = &model.Tag{
		Base: model.Base{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		Content: "tag",
	}
)

type TTagRepositoryMock struct {
}

func (m *TTagRepositoryMock) CreateTag(tag *model.Tag) (*model.Tag, error) {
	createdTag := createdTagMock
	return createdTag, nil
}

func TestTagService_CreateTag(t *testing.T) {
	target := NewTTagService(&TTagRepositoryMock{})
	want := createdTagMock
	got, err := target.CreateTag(TagMock)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
