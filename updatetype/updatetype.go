// Copyright 2024 midorimici
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

package updatetype

import "os"

// UpdateType represents the type of updating snapshots.
type UpdateType int

const (
	UpdateTypeNone   UpdateType = iota // Does not update snapshots.
	UpdateTypeNormal                   // Updates snapshots with the same name, but does not delete obsolete files.
	UpdateTypeAll                      // Removes snapshot directory at first, thus obsolete files are deleted.
)

// Type returns the specified update type.
//
// If the environment variable "UPDATE_SNAPSHOTS" is set to "1", it returns UpdateTypeNormal.
// If the environment variable "UPDATE_SNAPSHOTS" is set to "2", it returns UpdateTypeAll.
// Otherwise, it returns UpdateTypeNone.
func Type() UpdateType {
	v := os.Getenv("UPDATE_SNAPSHOTS")
	ut := UpdateTypeNone
	switch v {
	case "1":
		ut = UpdateTypeNormal

	case "2":
		ut = UpdateTypeAll
	}

	return ut
}
