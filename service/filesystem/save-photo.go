package filesystem

import (
	"io"
	"os"
	"fmt"
	"path/filepath"
	"mime/multipart"
)

func SavePhoto(file multipart.File, photoId uint64) error {

	out, err := os.Create(filepath.Join(FileSystemPath, fmt.Sprintf("%d.png", photoId)))
	if err != nil { return err }
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil { return err }

	return nil
}