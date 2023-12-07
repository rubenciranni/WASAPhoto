package filesystem

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
)

type AppFileSystem interface {
	SavePhoto(file multipart.File, photoId string) error
}

type appfsimpl struct {
	root   string
	photos string
}

func New(root string, photos string) (AppFileSystem, error) {

	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.Mkdir(root, 0755)
		if err != nil {
			return nil, fmt.Errorf("error creating filesystem root: %w", err)
		}
	}
	if _, err := os.Stat(path.Join(root, photos)); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(root, photos), 0755)
		if err != nil {
			return nil, fmt.Errorf("error creating filesystem photos directory: %w", err)
		}
	}

	return &appfsimpl{
		root:   root,
		photos: photos,
	}, nil
}
