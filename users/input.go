package user

// buat struct untuk menampung inputan user dari proses register
// pada proses register user hanya menginput field :
//  - name
//  - occupation
//  - email
//  - password

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required" validate:"min=5"`
}
