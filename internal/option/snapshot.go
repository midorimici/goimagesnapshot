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

type SnapshotOption interface {
	apply(c *SnapshotConfig)
}

type SnapshotConfig struct {
	name                     string
	threshold                float64
	shouldSkipByteComparison bool
}

func NewSnapshotConfig(opts ...SnapshotOption) *SnapshotConfig {
	c := defaultSnapshotConfig()
	for _, o := range opts {
		o.apply(c)
	}
	return c
}

func defaultSnapshotConfig() *SnapshotConfig {
	return &SnapshotConfig{}
}

func (c *SnapshotConfig) Name() string {
	return c.name
}

func (c *SnapshotConfig) Threshold() float64 {
	return c.threshold
}

func (c *SnapshotConfig) ShouldSkipByteComparison() bool {
	return c.shouldSkipByteComparison
}

func WithName(n string) SnapshotOption {
	return withName{n}
}

type withName struct{ n string }

func (n withName) apply(c *SnapshotConfig) {
	c.name = n.n
}

func WithThreshold(t float64) SnapshotOption {
	return withThreshold{t}
}

type withThreshold struct{ t float64 }

func (t withThreshold) apply(c *SnapshotConfig) {
	c.threshold = t.t
}

func WithOnlyPixelComparison() SnapshotOption {
	return withOnlyPixelComparison{}
}

type withOnlyPixelComparison struct{}

func (t withOnlyPixelComparison) apply(c *SnapshotConfig) {
	c.shouldSkipByteComparison = true
}
