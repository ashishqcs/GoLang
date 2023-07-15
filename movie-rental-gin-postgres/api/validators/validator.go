package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var ValidYear validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if year, ok := fieldLevel.Field().Interface().(int); ok {
		if year == 0 {
			return true
		}
		return year >= 1950 && year <= time.Now().Year()
	}
	return false
}

var ValidDate2 validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if date, ok := fieldLevel.Field().Interface().(string); ok {
		_, err := time.Parse("2006-01-02", date)
		if err == nil {
			return true
		}
	}
	return false
}
