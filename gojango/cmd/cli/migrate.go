package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := gg.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := gg.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := gg.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := gg.MigrateDownAll(dsn)
		if err != nil {
			return err
		}

		err = gg.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "force":
		err := gg.MigrateForce(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
