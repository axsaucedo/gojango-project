package gjango

import "path/filepath"

const version = "1.0.0"

type Gjango struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Gjango) New(rootPath string) error {
	pathConfig := InitPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (c *Gjango) Init(p InitPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		err := c.CreateDirIfNotExist(filepath.Join(root + "/" + path))
		if err != nil {
			return err
		}
	}
	return nil
}
