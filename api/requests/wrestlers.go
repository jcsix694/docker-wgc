package requests

type CreateWrestlerRequest struct {
	Name string `json:"name" binding:"required"`
}
