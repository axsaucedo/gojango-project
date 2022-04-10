package gojango

import (
	"os"
)

func (g *Gojango) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gojango) CreateFileIfNotExist(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	} else if err != nil {
		return err
	}

	return nil
}
