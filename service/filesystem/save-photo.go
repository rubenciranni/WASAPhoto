package filesystem

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (fs *appfsimpl) SavePhoto(file *multipart.File, photoId string) error {
	filename := fmt.Sprintf("%s.png", photoId)

	// Create the file on the filesystem
	dst, err := os.Create(filepath.Join(fs.root, fs.photos, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the contents of the file to the destination file
	_, err = io.Copy(dst, *file)
	return err
}
