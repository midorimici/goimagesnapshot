// Copyright 2023 midorimici
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package printer

import "fmt"

const (
	ylw = "\x1b[33m"
	end = "\x1b[0m"
)

func Yellow(s string) {
	fmt.Printf("%s%s%s\n", ylw, s, end)
}

func Yellowf(format string, a ...any) {
	Yellow(fmt.Sprintf(format, a...))
}
