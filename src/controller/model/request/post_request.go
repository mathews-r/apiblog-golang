package request

type PostRequest struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Category  string `json:"category" binding:"required"`
	UserId    int    `json:"userid"`
	Published string `json:"published"`
	Updated   string `json:"updated"`
}

type PostUpdateRequest struct {
	Title    string `json:"title" binding:"omitempty,min=4,max=50"`
	Content  string `json:"content" binding:"omitempty,min=1"`
	Category string `json:"category" binding:"required"`
}
