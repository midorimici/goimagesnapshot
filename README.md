# goimagesnapshot

<a href="https://godoc.org/github.com/midorimici/goimagesnapshot"><img src="https://godoc.org/github.com/midorimici/goimagesnapshot?status.svg" alt="GoDoc" /></a>
<a href="https://goreportcard.com/report/github.com/midorimici/goimagesnapshot"><img src="https://goreportcard.com/badge/github.com/midorimici/goimagesnapshot" alt="Go Report Card" /></a>

Snapshot testing for images in Go.

## Usage

### Using for a test

You can use `Match` function to test whether the given image is not changed as the previously generated snapshot image.
If no snapshot image is found, a new snapshot is created and the test fails.
By default, running a test creates `testdata/snapshots` directory beside a test file and snapshot images are put in the directory.

```go
import snap "github.com/midorimici/goimagesnapshot"

func TestSnapshot(t *testing.T) {
  t.Run("renders unchanged", func(t *testing.T) {
    snap.Match(t, img)
  })
}
```

### Updating a snapshot

You can update existing snapshots by running tests with the `UPDATE_SNAPSHOTS` environment variable set.

```
# Updates snapshots with the same name, but does not delete obsolete files.
UPDATE_SNAPSHOTS=1 go test ./...

# Removes the snapshot directory at first (i.e. obsolete files are deleted) and then snapshots are regenerated.
UPDATE_SNAPSHOTS=2 go test ./...
```

### Customization

You can specify some options to `Match` function or create a new snapshot matcher with custom options.

```go
func TestSnapshot(t *testing.T) {
  // Change snapshot output directory
  m := snap.NewMatcher(snap.WithDirectory("testdata"))
  
  t.Run("renders unchanged", func(t *testing.T) {
    m.Match(
      t,
      img,
      snap.WithName("custom_file_name"),
      snap.WithThreshold(0.2),
      snap.WithOnlyPixelComparison(),
    )
  })
}
```
