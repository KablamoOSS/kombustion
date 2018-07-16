package lock

import (
	manifestType "github.com/KablamoOSS/kombustion/internal/manifest"
)

// UpdateLock - update and write out a new lock file
func UpdateLock(manifest *manifestType.Manifest, newLockFile *Lock) error {
	// First load the lock file
	lockFile := FindAndLoadLock()

	err := updateLock(manifest, lockFile, newLockFile)

	return err
}

func updateLock(manifest *manifestType.Manifest, lockFile *Lock, newLockFile *Lock) error {
	// TODO: reconcile the manifest with the lock file

	err := WriteLockToDisk(newLockFile)
	return err
}
