package lock

import (
	manifestType "github.com/KablamoOSS/kombustion/internal/manifest"
)

// UpdateLock - update and write out a new lock file
func UpdateLock(manifest *manifestType.Manifest, newLockFile *Lock) error {
	// TODO: reconcile the manifest with the lock file
	lockFile := FindAndLoadLock()
	return WriteLockToDisk(lockFile)
}
