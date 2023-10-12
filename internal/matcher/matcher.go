package matcher

import (
	"fmt"
	"image"
	"os"
	"path"
	"strings"
	"syscall"
	"testing"

	"github.com/midorimici/goimagesnapshot/internal/option"
	"github.com/midorimici/goimagesnapshot/internal/printer"
	"github.com/midorimici/goimagesnapshot/internal/snapshot"
)

func Setup(config *option.MatcherConfig) error {
	const funcName = "matcher.Setup"

	if config.UpdateType() == option.UpdateTypeAll {
		if err := snapshot.RemoveDirectory(config.Directory()); err != nil {
			return fmt.Errorf("%s: %w", funcName, err)
		}
	}

	return nil
}

const extension = ".png"

func Match(t *testing.T, i image.Image, mc *option.MatcherConfig, sc *option.SnapshotConfig) {
	const funcName = "matcher.Match"

	defaultUmask := syscall.Umask(0)
	defer func() { syscall.Umask(defaultUmask) }()

	if _, err := os.Stat(mc.Directory()); os.IsNotExist(err) {
		// Create snapshot directory
		if err := os.Mkdir(mc.Directory(), 0777); err != nil {
			t.Errorf("%s: %v", funcName, err)
			return
		}
	} else if err != nil {
		t.Errorf("%s: %v", funcName, err)
		return
	}

	var snapshotName string
	if sc.Name() == "" {
		snapshotName = strings.ReplaceAll(t.Name(), "/", "-")
	} else {
		snapshotName = sc.Name()
	}

	filePath := path.Join(mc.Directory(), fmt.Sprintf("%s%s", snapshotName, extension))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// The snapshot does not exist
		// Take a snapshot
		snapshot.Take(filePath, i)
		printer.Yellowf("📸 Snapshot %s is not found. A new snapshot is created: %s", snapshotName, filePath)
		t.Fail()
		return
	} else if err != nil {
		t.Errorf("%s: %v", funcName, err)
		return
	}

	// The snapshot exists

	if mc.UpdateType() == option.UpdateTypeNormal {
		// Remove existing snapshots first when updating
		if err := os.Remove(filePath); err != nil {
			t.Errorf("%s: %v", funcName, err)
			return
		}

		// Take a snapshot
		snapshot.Take(filePath, i)
		printer.Yellowf("♻️ Snapshot %s is updated: %s", snapshotName, filePath)
		t.Fail()
		return
	}

	// Compare snapshots
	ok, err := snapshot.Compare(filePath, i, sc)
	if err != nil {
		t.Errorf("%s: %v", funcName, err)
		return
	}

	if !ok {
		t.Errorf("%s: snapshot does not match", funcName)
	}
}
