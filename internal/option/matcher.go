package option

import "os"

type updateType int

const (
	UpdateTypeNone   updateType = iota // Does not update snapshots.
	UpdateTypeNormal                   // Updates snapshots with the same name, but does not delete obsolete files.
	UpdateTypeAll                      // Removes snapshot directory at first, thus obsolete files are deleted.
)

type MatcherOption interface {
	apply(c *MatcherConfig)
}

type MatcherConfig struct {
	directory  string
	updateType updateType
}

func NewMatcherConfig(opts ...MatcherOption) *MatcherConfig {
	c := defaultMatcherConfig()
	for _, o := range opts {
		o.apply(c)
	}
	return c
}

func defaultMatcherConfig() *MatcherConfig {
	v := os.Getenv("UPDATE_SNAPSHOTS")
	ut := UpdateTypeNone
	switch v {
	case "1":
		ut = UpdateTypeNormal

	case "2":
		ut = UpdateTypeAll
	}

	return &MatcherConfig{
		directory:  "__snapshots__",
		updateType: ut,
	}
}

func (c *MatcherConfig) Directory() string {
	return c.directory
}

func (c *MatcherConfig) UpdateType() updateType {
	return c.updateType
}

func WithDirectory(d string) MatcherOption {
	return withDirectory{d}
}

type withDirectory struct{ d string }

func (d withDirectory) apply(c *MatcherConfig) {
	c.directory = d.d
}
