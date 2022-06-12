package validator

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg/entity"
)

type CreateProductValidator struct {
	Name   string
	Price  float32
	Errors map[string][]string
}

func (b *CreateProductValidator) GetFields() map[string]FieldSetter {
	return map[string]FieldSetter{
		"name": func(reqVal any) {
			if val, ok := reqVal.(string); ok {
				b.Name = val
				return
			}

			b.AddError(
				"name",
				MakeWrongTypeErrorMessage("string", pkg.GetType(reqVal)),
			)
		},
		"price": func(reqVal any) {
			if val, ok := reqVal.(float64); ok {
				b.Price = float32(val)
				return
			}

			b.AddError(
				"price",
				MakeWrongTypeErrorMessage("number", pkg.GetType(reqVal)),
			)
		},
	}
}

func (b CreateProductValidator) GetErrors() pkg.ResponseErrors {
	return &b.Errors
}

func (b *CreateProductValidator) AddError(fieldName string, errorMessage string) {
	if b.Errors == nil {
		b.Errors = map[string][]string{
			fieldName: {errorMessage},
		}
		return
	}

	b.Errors[fieldName] = append(b.Errors[fieldName], errorMessage)
}

func (b *CreateProductValidator) ValidateFields() {
	if b.Name == "" {
		b.AddError(
			"name",
			"The field is required",
		)
	}

	if b.Price == 0 {
		b.AddError(
			"price",
			"The field is required and must be > 0",
		)
	}
}

func (b CreateProductValidator) HasErrors() bool {
	return len(b.Errors) != 0
}

func (b *CreateProductValidator) Reset() {
	*b = CreateProductValidator{}
}

func (b CreateProductValidator) ToDto() any {
	return entity.ProductDto{
		Name:  b.Name,
		Price: b.Price,
	}
}
