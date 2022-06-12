package validator

import (
	"encoding/json"
	"fmt"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Validatable interface {
	GetFields() map[string]FieldSetter
	GetErrors() pkg.ResponseErrors
	AddError(fieldName string, errorMessage string)
	HasErrors() bool
	ToDto() any
	ValidateFields()
	Reset()
}

type FieldSetter func(v any)

func Validate(reqBody Validatable) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		values := map[string]any{}
		if err := json.Unmarshal(data, &values); err != nil {
			c.Error(pkg.BadRequestError(err, "unexpected JSON format"))
			c.Abort()
			return
		}

		fields := reqBody.GetFields()

		for name, value := range values {
			if setValue, ok := fields[name]; ok {
				setValue(value)
			}
		}

		reqBody.ValidateFields()

		if reqBody.HasErrors() {
			c.Error(pkg.CustomUnprocessableEntityError("validator error", reqBody.GetErrors()))
			c.Abort()
			reqBody.Reset()
			return
		}

		c.Set("body", reqBody.ToDto())
		reqBody.Reset()
		c.Next()
	}
}

func MakeWrongTypeErrorMessage(expected string, current string) string {
	return fmt.Sprintf("Unexpected type '%s' instead of '%s'", current, expected)
}
