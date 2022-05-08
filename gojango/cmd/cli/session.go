package main

import (
	"fmt"
	"time"
)

func doSessionTable() error {
	dbType := gg.DB.DataType

	if dbType == "mariadb" {
		dbType = "mysql"
	} else if dbType == "postgresql" {
		dbType = "postgres"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro)

	upFile := gg.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := gg.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table sessions"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}
