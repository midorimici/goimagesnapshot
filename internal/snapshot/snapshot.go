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
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	diff "github.com/olegfedoseev/image-diff"

	"github.com/midorimici/goimagesnapshot/internal/option"
	"github.com/midorimici/goimagesnapshot/internal/printer"
)

func RemoveDirectory(d string) error {
	const funcName = "snapshot.RemoveDirectory"

	// Remove snapshot directory
	if err := os.RemoveAll(d); err != nil {
		return fmt.Errorf("%s: %v", funcName, err)
	}
	printer.Yellow("🔥 Snapshot directory is removed.")

	return nil
}

func Compare(filePath string, i image.Image, config *option.SnapshotConfig) (bool, error) {
	const funcName = "snapshot.Compare"

	f, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}
	defer f.Close()

	eb := &bytes.Buffer{}
	if err := png.Encode(eb, i); err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}

	sb, err := io.ReadAll(f)
	if err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}

	snapshotImage, err := png.Decode(bytes.NewReader(sb))
	if err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}

	if !config.ShouldSkipByteComparison() && reflect.DeepEqual(eb.Bytes(), sb) {
		return true, nil
	}

	d, percent, err := diff.CompareImages(snapshotImage, i)
	if err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}

	isBelowThreshold := percent <= config.Threshold()
	if isBelowThreshold {
		return true, nil
	}

	// Create diff image file for visual comparison
	ext := path.Ext(filePath)

	fileNameWithoutExt := strings.TrimSuffix(filePath, ext)
	diffPath := fmt.Sprintf("%s_diff%s", fileNameWithoutExt, ext)
	diffFile, err := os.Create(diffPath)
	if err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}
	defer diffFile.Close()

	if err := png.Encode(diffFile, d); err != nil {
		return false, fmt.Errorf("%s: %w", funcName, err)
	}

	t := time.Now().UTC().Format("20060102150405")
	imgPath := fmt.Sprintf("%s_%s%s", fileNameWithoutExt, t, ext)
	Take(imgPath, i)

	printer.Yellowf("⚠️ Snapshots did not match (diff %.4f%%). Files %s and %s were created\n", percent, imgPath, diffPath)

	return false, nil
}

func Take(filePath string, i image.Image) error {
	const funcName = "snapshot.Take"

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%s: %w", funcName, err)
	}
	defer f.Close()

	if err := png.Encode(f, i); err != nil {
		return fmt.Errorf("%s: %w", funcName, err)
	}

	return nil
}
