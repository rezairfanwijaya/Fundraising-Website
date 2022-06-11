package user

// buat struct untuk menampung inputan user dari proses register
// pada proses register user hanya menginput field :
//  - name
//  - occupation
//  - email
//  - password

// pada proses login user hanya menginput field :
// - email
// - password

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=5"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type EmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
