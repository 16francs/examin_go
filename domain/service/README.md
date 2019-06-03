# domain/service

アプリケーション層から渡された値に対して処理を行う　　
いわゆる MVCS(Model View Controller Service) の S

## テンプレート

```go
package service

import (
	"fmt"

	"github.com/16francs/examin_go/domain/model"
)

// SampleService - Service のサンプルインターフェース
type SampleService interface {
	GetSample() (*model.Sample, error)
	PostSample(name string) (*model.Sample, error)
}

type sampleService struct {
}

// NewSampleService - SampleService の生成
func NewSampleService() SampleService {
	return &sampleService{}
}

func (s *sampleService) GetSample() (*model.Sample, error) {
	sample := &model.Sample{Message: "Hello, World!!"}
	return sample, nil
}

func (s *sampleService) PostSample(name string) (*model.Sample, error) {
	message := fmt.Sprintf("Hello, %s!!", name)
	sample := &model.Sample{Message: message}
	return sample, nil
}
```

## テストのテンプレート

```go
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
```
