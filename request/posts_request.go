package request

type PostsRequest struct {
	CategoryID uint   `form:"category_id" binding:"required"`
	Title      string `form:"title" binding:"required,min=3"`
	Content    string `form:"content" binding:"required,min=10"`
	UserID     uint   `form:"user_id" binding:"required"`
	Slug string `gorm:"-"`
}
