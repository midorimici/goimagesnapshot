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

package snapshot

import (
	"image"

	"github.com/midorimici/goimagesnapshot/internal/matcher"
	"github.com/midorimici/goimagesnapshot/internal/option"
	"github.com/midorimici/goimagesnapshot/internal/testing"
)

var globalMatcher = NewMatcher()

// Match is matcher.Match with the default matcher options.
func Match(t testing.T, i image.Image, opts ...option.SnapshotOption) {
	globalMatcher.Match(t, i, opts...)
}

type Matcher interface {
	Match(t testing.T, i image.Image, opts ...option.SnapshotOption)
}

type mtchr struct {
	config *option.MatcherConfig
}

// NewMatcher returns a new matcher with matcher options.
func NewMatcher(opts ...option.MatcherOption) Matcher {
	const funcName = "snapshot.NewMatcher"

	m := &mtchr{}
	m.config = option.NewMatcherConfig(opts...)
	matcher.Setup(m.config)

	return m
}

// Match checks whether the image matches to the existing snapshot image.
//
// If no snapshot image is found, it creates a new snapshot image from the current image and fails the test.
func (m *mtchr) Match(t testing.T, i image.Image, opts ...option.SnapshotOption) {
	const funcName = "matcher.Match"

	t.Helper()
	if t.Failed() {
		return
	}

	sc := option.NewSnapshotConfig(opts...)

	matcher.Match(t, i, m.config, sc)
}
