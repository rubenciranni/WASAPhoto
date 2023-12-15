package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func (fs *appfsimpl) GetPhotoPath(photoID string) (string, error) {
	filename := fmt.Sprintf("%s.png", photoID)
	path := filepath.Join(fs.root, fs.photos, filename)
	_, err := os.Stat(path)
	return path, err
}
