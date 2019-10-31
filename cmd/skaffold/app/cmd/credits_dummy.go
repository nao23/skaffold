// +build !release

/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"errors"
	"io"
)

// exportCredits with !release build tag is just here for compilation purposes
// this file does not depend on the generated statik.go file that is not checked
// in by default to git
func exportCredits(out io.Writer) error {
	return errors.New("not implemented, skaffold should be built with make ('release' build tag)")
}