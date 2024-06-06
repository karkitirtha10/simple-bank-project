package systemerror

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/karkitirtha10/simplebank/app/localization"
)

type ValidationMessage struct {
	Field   string
	Tag     string
	Message string
}

func ToValidationErrorBag(
	original validator.ValidationErrors,
	localization localization.ILocalization,
	inputStruct interface{},
) []ValidationMessage {

	var (
		validationMessages []ValidationMessage
	)

	for _, fieldError := range original {
		//fieldError.Field() geves struct name instead of json name. hence getting json name using reflect package
		t := reflect.TypeOf(inputStruct)
		field, _ := t.FieldByName(fieldError.Field())
		jsonName := field.Tag.Get("json")
		localizedFieldName := localization.Translate(jsonName)
		normalizedTagName := fieldError.Tag() + "Validation"

		message := localization.TranslateTemplate(
			normalizedTagName,
			map[string]interface{}{
				"Field": localizedFieldName,
				"Param": fieldError.Param(), //all tags may not need Param
				"Value": fieldError.Value(), //all tags may not need value
			},
		)

		validationMessages = append(
			validationMessages,
			ValidationMessage{
				Field:   jsonName,
				Tag:     fieldError.Tag(),
				Message: message,
			},
		)
	}

	return validationMessages
}

/*
// ValidationError do not log ValidationError
type ValidationError struct {
	message  string //use for future
	original validator.ValidationErrors
	source   ErrorSource
}

func (yo *ValidationError) Error() string {
	return yo.original.Error() // or first error
}

func (yo ValidationError) response() {

	localization := services.NewLocalization("np")

	var (
		validationMessages []ValidationMessage
	)
	for _, fieldError := range yo.original {
		localizedField := localization.Translate(fieldError.Field())
		normalizedTagName := fieldError.Tag() + "Validation"

		message := localization.TranslateTemplate(
			normalizedTagName,
			map[string]interface{}{
				"Field": localizedField,
				"Value": nil, //all tags may not need value
			},
		)

		validationMessages = append(
			validationMessages,
			ValidationMessage{
				Field:   fieldError.Field(),
				Tag:     fieldError.Tag(),
				Message: message,
			},
		)

	}
}
*/
