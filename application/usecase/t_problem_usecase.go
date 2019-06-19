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
	tProblemService     service.TProblemService
	tProblemsTagService service.TProblemsTagService
	tTagService         service.TTagService
}

// NewTProblemUsecase - tProblemUsecase の生成
func NewTProblemUsecase(ps service.TProblemService, pts service.TProblemsTagService, ts service.TTagService) TProblemUsecase {
	return &tProblemUsecase{ps, pts, ts}
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
	// TODO: 全てのタグを取得
	createdTags := make([]*model.Tag, 0, 4)
	for _, v := range tags {
		tag := &model.Tag{Content: v}

		// TODO: タグがすでに存在する場合、タグ情報の取得
		// TODO: そうでない場合は、タグを登録する
		createdTag, err := u.tTagService.CreateTag(tag)
		if err != nil {
			return nil, nil, err
		}

		createdTags = append(createdTags, createdTag)
	}

	// 問題集とタグを関連付け
	for _, v := range createdTags {
		problemsTag := &model.ProblemsTag{
			ProblemID: createdProblem.Base.ID,
			TagID:     v.Base.ID,
		}

		_, err := u.tProblemsTagService.CreateProblemsTag(problemsTag)
		if err != nil {
			return nil, nil, err
		}
	}

	return createdProblem, createdTags, nil
}
