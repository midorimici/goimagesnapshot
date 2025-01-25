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

import "github.com/midorimici/goimagesnapshot/internal/option"

// WithDirectory returns a snapshot matcher option
// which specifies the directory name where the output snapshot images are put.
//
// The default value is "testdata/snapshots".
func WithDirectory(d string) option.MatcherOption {
	return option.WithDirectory(d)
}

// WithName returns a snapshot option
// which specifies the output snapshot file name.
//
// A file name uses the test name by default.
func WithName(n string) option.SnapshotOption {
	return option.WithName(n)
}

// WithThreshold returns a snapshot option
// which specifies a value which a test fails when the differences by percent in compared images exceeds.
// For example, when compared images have 0.11% of differences, a test fails with threshold of 0.1,
// whereas it passes with threshold of 0.2.
//
// The default value is 0, i.e. two images must match perfectly.
func WithThreshold(t float64) option.SnapshotOption {
	return option.WithThreshold(t)
}

// WithOnlyPixelComparison returns a snapshot option
// which specifies to compare only image pixels between two snapshots,
// which generally makes the test more strict.
//
// By default, the matcher first compares the two image byte slices.
// If they differ, then their image pixels are tested one by one.
//
// Specifying this option skips the first byte slice comparison.
func WithOnlyPixelComparison() option.SnapshotOption {
	return option.WithOnlyPixelComparison()
}
