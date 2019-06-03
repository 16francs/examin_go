# aplication

ユースケース?ビジネスロジックの記述をする
handler から受け取った値を service に送る
service からの結果を handler に返す
また、エラー情報などはここで追加したりするらしい...

## テンプレート

```go
package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

// SampleUsecase - Usecase のサンプル
type SampleUsecase interface {
	GetSample() (*model.Sample, error)
	PostSample(name string) (*model.Sample, error)
}

type sampleUsecase struct {
	service service.SampleService
}

// NewSampleUsecase - sampleUsecase の生成
func NewSampleUsecase(s service.SampleService) SampleUsecase {
	return &sampleUsecase{service: s}
}

func (u *sampleUsecase) GetSample() (*model.Sample, error) {
	sample, err := u.service.GetSample()
	if err != nil {
		return nil, err
	}

	return sample, nil
}
```

## テストのテンプレート

```go
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
```
