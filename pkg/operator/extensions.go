package operator

import (
	"fmt"
	"strings"
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	messages := make([]string, len(v))

	for idx := range v {
		validationError := &v[idx]
		messages[idx] = validationError.Error()
	}

	return strings.Join(messages, ";")
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"encountered a validation error, message: %s, value: %s, pathComponents: %s",
		v.Message,
		v.Value,
		v.PathComponents,
	)
}
