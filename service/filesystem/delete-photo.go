package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func (fs *appfsimpl) DeletePhoto(photoID string) error {
	filename := fmt.Sprintf("%s.png", photoID)
	path := filepath.Join(fs.root, fs.photos, filename)
	err := os.Remove(path)
	return err
}
