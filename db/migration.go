package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/karkitirtha10/simplebank/config"
)

type DBMigration struct{}

func (d DBMigration) MigrateUp(step int, c config.Config) {
	m, err := migrate.New("file://"+c.MigrationUrl, c.DbUrl)
	if err != nil {
		log.Fatal("cannot create new migration instance: ", err)
	}
	defer m.Close()

	// ///////
	if step != 0 {
		err = m.Steps(step)
	} else {
		err = m.Up()
	}

	if err == nil {
		log.Println("migrated successfully")
	} else if err == migrate.ErrNoChange {
		log.Println("nothing to migrate")
	} else if e, ok := err.(*migrate.ErrShortLimit); ok {
		log.Println("migrated " + fmt.Sprint(step-int(e.Short)) + " files")
	} else {
		log.Fatal("failed to migrate: ", err)
	}

	////////
	// if err = m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal("failed to run migrateion: ", err)
	// }

	// if step != 0 {
	// 	m.Steps(step)
	// }

	// if err == migrate.ErrNoChange {
	// 	log.Println("nothing to migrate")
	// } else {
	// 	log.Println("migrated successfully")
	// }
}

func (d DBMigration) MigrateDown(step int, c config.Config) {
	m, err := migrate.New("file://"+c.MigrationUrl, c.DbUrl)
	if err != nil {
		log.Fatal("cannot create new migration instance: ", err)
	}
	defer m.Close()

	fmt.Println(step)
	if step != 0 {
		err = m.Steps(-step)
	} else {
		err = m.Down()
	}

	if err == nil {
		log.Println("migration rolledback successfully")
	} else if err == migrate.ErrNoChange {
		log.Println("nothing to rollback")
	} else if e, ok := err.(*migrate.ErrShortLimit); ok {
		log.Println("migrated " + fmt.Sprint(step-int(e.Short)) + " files")
	} else {
		log.Fatal("failed to rollback migration: ", err)
	}
}

func (d DBMigration) CreateMigration(dir string, name string) error {
	var version string
	var err error

	dir = filepath.Clean(dir)
	ext := ".sql"

	startTime := time.Now()
	version = strconv.FormatInt(startTime.UnixNano(), 10)

	versionGlob := filepath.Join(dir, version+"_*"+ext)
	matches, err := filepath.Glob(versionGlob)

	if err != nil {
		return err
	}

	if len(matches) > 0 {
		return fmt.Errorf("duplicate migration version: %s", version)
	}

	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%s_%s.%s%s", version, name, direction, ext)
		filename := filepath.Join(dir, basename)

		if err = d.createFile(filename); err != nil {
			return err
		}

		if print := true; print {
			absPath, _ := filepath.Abs(filename)
			log.Println(absPath)
		}
	}

	return nil
}

func (d DBMigration) createFile(filename string) error {
	// create exclusive (fails if file already exists)
	// os.Create() specifies 0666 as the FileMode, so we're doing the same
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}
