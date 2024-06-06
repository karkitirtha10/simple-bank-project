package inputmodel

// AddAccountInput input
type AddAccountInput struct {
	//Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}
