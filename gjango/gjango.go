package gjango

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Gjango struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
}

type config struct {
	port     string
	renderer string
}

func (c *Gjango) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(filepath.Join(rootPath, "/.env"))
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath

	return nil
}

func (c *Gjango) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (c *Gjango) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		err := c.CreateDirIfNotExist(filepath.Join(root + "/" + path))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Gjango) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
