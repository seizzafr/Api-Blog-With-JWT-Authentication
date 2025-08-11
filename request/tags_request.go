package request

type TagsRequest struct {
	Name string `json:"name" binding:"required"`
	
}
