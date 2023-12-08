package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func (fs *appfsimpl) GetPhotoPath(photoId string) (string, error) {
	filename := fmt.Sprintf("%s.png", photoId)
	path := filepath.Join(fs.root, fs.photos, filename)
	_, err := os.Stat(path)
	return path, err
}
