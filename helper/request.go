package helper

import (
	validator2 "github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var tagTranslations = map[string]string{
	"gte": "min",
	"lte": "max",
}

func Validate(body reflect.Type, errors error) interface{} {
	errorMessages := make(map[string][]string)

	for _, err := range errors.(validator2.ValidationErrors) {
		field, _ := body.FieldByName(err.Field())

		fieldTag := field.Tag.Get("json")
		if fieldTag == "" {
			fieldTag = field.Tag.Get("form")
			if fieldTag == "" {
				fieldTag = field.Tag.Get("query")
				if fieldTag == "" {
					fieldTag = err.Field()
				}
			}
		}

		if idx := strings.Index(fieldTag, "["); idx != -1 {
			fieldTag = strings.ToLower(fieldTag[:idx])
		}

		errTag := err.Tag()

		if translatedTag, found := tagTranslations[errTag]; found {
			errTag = translatedTag
		}

		errorMessages[fieldTag] = append(errorMessages[fieldTag], errTag)
	}

	return errorMessages
}
