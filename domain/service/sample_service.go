package service

import (
	"github.com/16francs/examin_go/domain/model"
)

// SampleService - Service のサンプルインターフェース
type SampleService interface {
	GetSample() (*model.Sample, error)
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
