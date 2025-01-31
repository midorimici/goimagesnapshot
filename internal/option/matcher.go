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

package option

import (
	"github.com/midorimici/goimagesnapshot/updatetype"
)

type MatcherOption interface {
	apply(c *MatcherConfig)
}

type MatcherConfig struct {
	directory  string
	updateType updatetype.UpdateType
}

func NewMatcherConfig(opts ...MatcherOption) *MatcherConfig {
	c := defaultMatcherConfig()
	for _, o := range opts {
		o.apply(c)
	}
	return c
}

func defaultMatcherConfig() *MatcherConfig {
	ut := updatetype.Type()

	return &MatcherConfig{
		directory:  "testdata/snapshots",
		updateType: ut,
	}
}

func (c *MatcherConfig) Directory() string {
	return c.directory
}

func (c *MatcherConfig) UpdateType() updatetype.UpdateType {
	return c.updateType
}

func WithDirectory(d string) MatcherOption {
	return withDirectory{d}
}

type withDirectory struct{ d string }

func (d withDirectory) apply(c *MatcherConfig) {
	c.directory = d.d
}
