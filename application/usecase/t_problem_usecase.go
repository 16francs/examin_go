package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

// TProblemUsecase - 講師用の 問題集ユースケース
type TProblemUsecase interface {
	CreateProblem(title, content, loginID string) (*model.Problem, error)
}

type tProblemUsecase struct {
	TProblemService service.TProblemService
	UserService     service.UserService
}

// NewTProblemUsecase - tProblemUsecase の生成
func NewTProblemUsecase(ps service.TProblemService, us service.UserService) TProblemUsecase {
	return &tProblemUsecase{ps, us}
}

func (u *tProblemUsecase) CreateProblem(title, content, loginID string) (*model.Problem, error) {
	// ログイン情報の取得
	teacher, err := u.UserService.Show(loginID)
	if err != nil {
		return nil, err
	}

	problem := &model.Problem{
		Title:   title,
		Content: content,
		UserID:  teacher.ID,
	}

	createdProblem, err := u.TProblemService.CreateProblem(problem)
	if err != nil {
		return nil, err
	}

	return createdProblem, nil
}
