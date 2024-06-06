package inputmodel

// binding tag is used for validation
// json tag is used to serialize json body to field in struct
type LoginInput struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required" validate:"required,max:1"`
// 	Password string `json:"password" binding:"required" validate:"required,max:1"`
// }
