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

func (u *sampleUsecase) PostSample(name string) (*model.Sample, error) {
	sample, err := u.service.PostSample(name)
	if err != nil {
		return nil, err
	}

	return sample, nil
}
