package httpv1

import "github.com/AlexeyLoychenko/person_api/internal/model"

func OK() model.WebResponse[string] {
	return model.WebResponse[string]{Data: "OK"}
}

func Error(err string) model.WebResponse[any] {
	return model.WebResponse[any]{
		Error: err,
	}
}
