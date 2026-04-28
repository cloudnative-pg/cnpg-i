/*
Copyright © contributors to CloudNativePG, established as
CloudNativePG a Series of LF Projects, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package operator //nolint:revive,stylecheck

import (
	"fmt"
	"strings"
)

// ValidationErrors is a list of validation errors.
type ValidationErrors []*ValidationError

func (v ValidationErrors) Error() string {
	messages := make([]string, len(v))

	for idx := range v {
		messages[idx] = v[idx].Error()
	}

	return strings.Join(messages, ";")
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"encountered a validation error, message: %s, value: %s, pathComponents: %s",
		v.GetMessage(),
		v.GetValue(),
		v.GetPathComponents(),
	)
}
