package helper

import "github.com/go-playground/validator/v10"

// disini kita akan membuat response dari json ketika user berhasil ataupun gagal meng-hit endpoint yang telah disediakan

/*
	meta : {
		"message"	: "Berhasil Register"
		"code"		: 200
		"status"	: "Berhasil"
	},
	data : {
		"id"		: 1,
		"name"		: "Abdas Fernanda",
		"email" 	: "abdas@gmail.com",
		"token" 	: "abdas123rtybnm",
		"occupation": "Backend"
	}
*/

// kita harus bikin 2 struct [respons dan meta] untuk data itu dia dianamis nantinya jadi kita biki tipe dia jadi inteface saja

type respons struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponsAPI(message, status string, code int, data interface{}) respons {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responsAPI := respons{
		Meta: meta,
		Data: data,
	}

	return responsAPI
}

// format error
func ErrorFormater(err error) []string {
	var myerror []string

	for _, e := range err.(validator.ValidationErrors) {
		myerror = append(myerror, e.Error())
	}

	return myerror
}
