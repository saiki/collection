package command

import (
	"os/user"
	"path/filepath"
)

const BUCKET_NAME string = "storage"

func home() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func storePath(path string) (string, error) {
	if path == "" {
		homeDir, err := home()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, ".collection"), nil
	}
	return path, nil
}
