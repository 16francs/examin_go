package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

// TProblemUsecase - 講師用の 問題集ユースケース
type TProblemUsecase interface {
	CreateProblem(title, content string, userID uint, tags []string) (*model.Problem, []*model.Tag, error)
}

type tProblemUsecase struct {
	tProblemService service.TProblemService
	tTagService     service.TTagService
}

// NewTProblemUsecase - tProblemUsecase の生成
func NewTProblemUsecase(ps service.TProblemService, ts service.TTagService) TProblemUsecase {
	return &tProblemUsecase{ps, ts}
}

func (u *tProblemUsecase) CreateProblem(title, content string, userID uint, tags []string) (*model.Problem, []*model.Tag, error) {
	// 問題集の登録
	problem := &model.Problem{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	createdProblem, err := u.tProblemService.CreateProblem(problem)
	if err != nil {
		return nil, nil, err
	}

	// 問題集のタグを登録
	createdTags := make([]*model.Tag, 0, 4)
	for _, v := range tags {
		tag := &model.Tag{Content: v}

		createdTag, err := u.tTagService.CreateTag(tag)
		if err != nil {
			return nil, nil, err
		}

		createdTags = append(createdTags, createdTag)
	}

	// TODO: 問題集とタグを関連付け

	return createdProblem, createdTags, nil
}
