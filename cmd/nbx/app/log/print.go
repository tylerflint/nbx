/*
Copyright 2017 Nanobox, Inc.

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

package log

import (
	"fmt"
	"io"
	"os"
)

type Printer interface {
	Print(format string, args ...interface{})
}

type fmtPrinter struct {
	Out io.Writer
}

// Create a new FmtPrinter to print to stderr
func newFmtPrinter() *fmtPrinter {
	return &fmtPrinter{
		Out: os.Stderr,
	}
}

// Print will use the 'fmt' package to print a message to Out
func (p fmtPrinter) Print(format string, args ...interface{}) {
	fmt.Fprintf(p.Out, format, args...)
}
