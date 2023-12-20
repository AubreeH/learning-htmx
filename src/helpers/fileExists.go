package helpers

import (
	"errors"
	"os"
)

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return !info.IsDir(), nil
}
