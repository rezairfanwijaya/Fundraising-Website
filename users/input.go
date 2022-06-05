package user

// buat struct untuk menampung inputan user dari proses register
// pada proses register user hanya menginput field :
//  - name
//  - occupation
//  - email
//  - password

type RegisterUserInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}
