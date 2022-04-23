package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := gg.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)

		upFile := gg.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := gg.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		fmt.Printf("%s %s\n", upFile, downFile)
	}

	return nil
}
