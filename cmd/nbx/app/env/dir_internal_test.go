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

package env

import (
	"os"
	"testing"
)

func TestSystemMkdir(t *testing.T) {
	creator := &SystemDirCreator{
		mkdirAllFunc: func(dir string, perms os.FileMode) error {
			return nil
		},
	}

	err := creator.Mkdir("/tmp", 0755)
	if err != nil {
		t.Errorf("Was expecting err to be nil, got: %v")
	}
}
