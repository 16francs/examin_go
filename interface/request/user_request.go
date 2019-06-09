package request

type CreateUser struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password"  binding:"required"`
	Name     string `json:"name"  binding:"required"`
	School   string `json:"school"`
}
