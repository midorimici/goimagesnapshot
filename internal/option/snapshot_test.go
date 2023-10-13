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

package option_test

import (
	"testing"

	"github.com/midorimici/goimagesnapshot/internal/option"
)

func Test_NewSnapshotConfig(t *testing.T) {
	type want struct {
		name                     string
		threshold                float64
		shouldSkipByteComparison bool
	}
	tests := []struct {
		name string
		opts []option.SnapshotOption
		want want
	}{
		{
			name: "returns a correct snapshot config with no option",
		},
		{
			name: "returns a correct snapshot config with WithName option",
			opts: []option.SnapshotOption{
				option.WithName("name"),
			},
			want: want{
				name: "name",
			},
		},
		{
			name: "returns a correct snapshot config with WithThreshold option",
			opts: []option.SnapshotOption{
				option.WithThreshold(0.1),
			},
			want: want{
				threshold: 0.1,
			},
		},
		{
			name: "returns a correct snapshot config with WithOnlyPixelComparison option",
			opts: []option.SnapshotOption{
				option.WithOnlyPixelComparison(),
			},
			want: want{
				shouldSkipByteComparison: true,
			},
		},
		{
			name: "returns a correct snapshot config with WithName and WithThreshold option",
			opts: []option.SnapshotOption{
				option.WithName("name"),
				option.WithThreshold(0.1),
			},
			want: want{
				name:      "name",
				threshold: 0.1,
			},
		},
		{
			name: "returns a correct snapshot config with WithName and WithOnlyPixelComparison option",
			opts: []option.SnapshotOption{
				option.WithName("name"),
				option.WithOnlyPixelComparison(),
			},
			want: want{
				name:                     "name",
				shouldSkipByteComparison: true,
			},
		},
		{
			name: "returns a correct snapshot config with WithThreshold and WithOnlyPixelComparison option",
			opts: []option.SnapshotOption{
				option.WithThreshold(0.1),
				option.WithOnlyPixelComparison(),
			},
			want: want{
				threshold:                0.1,
				shouldSkipByteComparison: true,
			},
		},
		{
			name: "returns a correct snapshot config with WithName and WithThreshold and WithOnlyPixelComparison option",
			opts: []option.SnapshotOption{
				option.WithName("name"),
				option.WithThreshold(0.1),
				option.WithOnlyPixelComparison(),
			},
			want: want{
				name:                     "name",
				threshold:                0.1,
				shouldSkipByteComparison: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := option.NewSnapshotConfig(tt.opts...)

			if got := c.Name(); got != tt.want.name {
				t.Errorf("option.NewSnapshotConfig() Name() = %v, want %v", got, tt.want.name)
			}

			if got := c.Threshold(); got != tt.want.threshold {
				t.Errorf("option.NewSnapshotConfig() Threshold() = %v, want %v", got, tt.want.threshold)
			}

			if got := c.ShouldSkipByteComparison(); got != tt.want.shouldSkipByteComparison {
				t.Errorf("option.NewSnapshotConfig() ShouldSkipByteComparison() = %v, want %v", got, tt.want.shouldSkipByteComparison)
			}
		})
	}
}
