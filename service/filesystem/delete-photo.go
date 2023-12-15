package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func (fs *appfsimpl) DeletePhoto(photoId string) error {
	filename := fmt.Sprintf("%s.png", photoId)
	path := filepath.Join(fs.root, fs.photos, filename)
	err := os.Remove(path)
	return err
}
