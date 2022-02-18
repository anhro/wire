// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"

type (
	foo struct {
		name string
		doo  Doo
	}
)

func provideFoo(doo Doo) Foo {
	return &foo{name: "Foo", doo: doo}
}

func (f *foo) GetName() string {
	return fmt.Sprintf("%s.%s", f.name, f.doo.GetName())
}

func (f *foo) GetDoo() Doo {
	return f.doo
}
