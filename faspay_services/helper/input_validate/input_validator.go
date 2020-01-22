package inputvalidate

import (
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func InputValidateStruct(req interface{}) error {
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {
		//fmt.Println("InvalidValidationError err: ", err)
		//validationErrors := err.(validator.ValidationErrors)
		//fmt.Println("InvalidValidationError validationErrors: ", validationErrors)
	}
	return err
}
