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
	"os"
	"testing"

	"github.com/midorimici/goimagesnapshot/internal/option"
	"github.com/midorimici/goimagesnapshot/updatetype"
)

func Test_NewMatcherConfig(t *testing.T) {
	type want struct {
		directory  string
		updateType updatetype.UpdateType
	}
	tests := []struct {
		name               string
		envUpdateSnapshots string
		opts               []option.MatcherOption
		want               want
	}{
		{
			name: "returns a correct matcher config with no option and environment variable",
			want: want{
				directory:  "testdata/snapshots",
				updateType: updatetype.UpdateTypeNone,
			},
		},
		{
			name:               "returns a correct matcher config with no option and UPDATE_SNAPSHOTS=1",
			envUpdateSnapshots: "1",
			want: want{
				directory:  "testdata/snapshots",
				updateType: updatetype.UpdateTypeNormal,
			},
		},
		{
			name:               "returns a correct matcher config with no option and UPDATE_SNAPSHOTS=2",
			envUpdateSnapshots: "2",
			want: want{
				directory:  "testdata/snapshots",
				updateType: updatetype.UpdateTypeAll,
			},
		},
		{
			name: "returns a correct matcher config with WithDirectory option and no environment variable",
			opts: []option.MatcherOption{option.WithDirectory("test_dir")},
			want: want{
				directory:  "test_dir",
				updateType: updatetype.UpdateTypeNone,
			},
		},
		{
			name:               "returns a correct matcher config with WithDirectory option and UPDATE_SNAPSHOTS=1",
			envUpdateSnapshots: "1",
			opts:               []option.MatcherOption{option.WithDirectory("test_dir")},
			want: want{
				directory:  "test_dir",
				updateType: updatetype.UpdateTypeNormal,
			},
		},
		{
			name:               "returns a correct matcher config with WithDirectory option and UPDATE_SNAPSHOTS=2",
			envUpdateSnapshots: "2",
			opts:               []option.MatcherOption{option.WithDirectory("test_dir")},
			want: want{
				directory:  "test_dir",
				updateType: updatetype.UpdateTypeAll,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("UPDATE_SNAPSHOTS", tt.envUpdateSnapshots)

			c := option.NewMatcherConfig(tt.opts...)

			if got := c.Directory(); got != tt.want.directory {
				t.Errorf("option.NewMatcherConfig() Directory() = %v, want %v", got, tt.want.directory)
			}

			if got := c.UpdateType(); got != tt.want.updateType {
				t.Errorf("option.NewMatcherConfig() UpdateType() = %v, want %v", got, tt.want.updateType)
			}
		})
	}
}
