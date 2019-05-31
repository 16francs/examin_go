package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/16francs/examin_go/domain/model"
)

type SampleServiceMock struct {
}

func (m *SampleServiceMock) GetSample() (*model.Sample, error) {
	sample := &model.Sample{Message: "Hello, World"}
	return sample, nil
}

func (m *SampleServiceMock) PostSample(name string) (*model.Sample, error) {
	message := fmt.Sprintf("Hello, %s!!", name)
	sample := &model.Sample{Message: message}
	return sample, nil
}

func TestSampleUsecase_GetSample(t *testing.T) {
	target := NewSampleUsecase(&SampleServiceMock{})
	want := &model.Sample{Message: "Hello, World"}
	got, err := target.GetSample()

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestSampleUsecase_PostSample(t *testing.T) {
	target := NewSampleUsecase(&SampleServiceMock{})
	want := &model.Sample{Message: "Hello, Taro!!"}
	got, err := target.PostSample("Taro")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}	
}
