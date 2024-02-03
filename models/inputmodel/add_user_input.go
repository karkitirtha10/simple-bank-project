package inputmodel

type AddUserInput struct {
	Name     string `json:"name" binding:"required" validate:"required,max:2"`
	Email    string `json:"email" binding:"required" validate:"required,max:2"`
	Password string `json:"password" binding:"required" validate:"required,max:2"`
	// u_updated_at TIMESTAMPTZ
}
