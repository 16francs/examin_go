package request

// TCreateTeacher - 講師向けの講師作成のリクエスト
type TCreateTeacher struct {
	LoginID string `json:"login_id" binding:"required,max=31"`
	Name    string `json:"name" binding:"required,max=63"`
	School  string `json:"school" binding:"required,max=63"`
}
