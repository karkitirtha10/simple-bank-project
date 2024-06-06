package inputmodel

type AddUserInput struct {
	Name     string `json:"name" binding:"required" validate:"required,max:255"`
	Email    string `json:"email" binding:"required" validate:"required,max:255"`
	Password string `json:"password" binding:"required" validate:"required,max:255"`
	// u_updated_at TIMESTAMPTZ
}
