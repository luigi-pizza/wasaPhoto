package filesystem

import (
	"os"
	"fmt"
	"path/filepath"
)

func DeletePhoto(photoId uint64) error {

	err := os.Remove(filepath.Join(FileSystemPath, fmt.Sprintf("%d.png", photoId)))
	if err != nil { return err }

	return nil
}