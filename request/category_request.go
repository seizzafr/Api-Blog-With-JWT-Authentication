package request

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
	
}
