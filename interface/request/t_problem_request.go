package request

// TCreateProblem - 講師向けの問題集作成のリクエスト
type TCreateProblem struct {
	Title   string   `json:"title" binding:"required,max=31"`
	Content string   `json:"content" binding:"required,max=63"`
	Tags    []string `json:"tags" binding:"lt=5"`
}
