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
