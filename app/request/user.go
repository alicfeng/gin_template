package request

type FooRequest struct {
	FormRequest
	Name string `json:"name" comment:"名称" binding:"required"`
}
