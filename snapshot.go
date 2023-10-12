package snapshot

import (
	"image"
	"testing"

	"github.com/midorimici/goimagesnapshot/internal/matcher"
	"github.com/midorimici/goimagesnapshot/internal/option"
)

var globalMatcher = NewMatcher()

// Match is matcher.Match with the default matcher options.
func Match(t *testing.T, i image.Image, opts ...option.SnapshotOption) {
	globalMatcher.Match(t, i, opts...)
}

type Matcher interface {
	Match(t *testing.T, i image.Image, opts ...option.SnapshotOption)
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
func (m *mtchr) Match(t *testing.T, i image.Image, opts ...option.SnapshotOption) {
	const funcName = "matcher.Match"

	t.Helper()
	if t.Failed() {
		return
	}

	sc := option.NewSnapshotConfig(opts...)

	matcher.Match(t, i, m.config, sc)
}
