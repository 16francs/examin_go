package request

// PostSample - PostSample のリクエスト
type PostSample struct {
	Name string `json:"name" binding:"required"`
}
