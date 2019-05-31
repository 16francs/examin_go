package service

import (
	"reflect"
	"testing"

	"github.com/16francs/examin_go/domain/model"
)

func TestSampleService_GetSample(t *testing.T) {
	target := NewSampleService()
	want := &model.Sample{Message: "Hello, World!!"}
	got, err := target.GetSample()

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestSampleService_PostSample(t *testing.T) {
	target := NewSampleService()
	want := &model.Sample{Message: "Hello, Taro!!"}
	got, err := target.PostSample("Taro")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
