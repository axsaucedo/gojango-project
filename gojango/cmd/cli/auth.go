package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func doAuth() error {
	// migrations
	dbType := gg.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := gg.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := gg.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	dropTableCmd := []string{
		"drop table if exists users cascade;",
		"drop table if exists tokens cascade;",
		"drop table if exists remember_tokens cascade;",
	}

	err = copyDataToFile([]byte(strings.Join(dropTableCmd, " ")), downFile)
	if err != nil {
		exitGracefully(err)
	}

	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/user.go.txt", gg.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", gg.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}
