// Package snapshot provides utility routines for image snapshot testing.
//
// With the package, you can ensure that an image look like the same as before: that there is no visual regression.
//
// A PNG file is generated when running a test with Match function, and used to compare with images generated in subsequent tests.
//
//	import snap "github.com/midorimici/goimagesnapshot"
//
//	func TestSnapshot(t *testing.T) {
//	  t.Run("renders unchanged", func(t *testing.T) {
//	    snap.Match(t, img)
//	  })
//	}
//
// To update snapshots, set UPDATE_SNAPSHOTS environment variable when running tests.
//
//	// Updates snapshots with the same name, but does not delete obsolete files.
//	UPDATE_SNAPSHOTS=1 go test ./...
//
//	// Removes the snapshot directory at first (i.e. obsolete files are deleted) and then snapshots are regenerated.
//	UPDATE_SNAPSHOTS=2 go test ./...
package snapshot
