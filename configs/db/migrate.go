package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var MigrateUp = &cobra.Command{
	Use:   "migrate",
	Short: "Running Database Migrate",
	RunE:  migrationUp,
}

var MigrateDown = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback Database Migration",
	RunE:  migrationDown,
}

func migrationUp(cmd *cobra.Command, args []string) error {
	db, _ := New()
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://configs/db/migrate",
		"postgres", driver,
	)

	if err != nil {
		return err
	}

	errMig := mig.Up()

	if errMig != nil {
		return errMig
	}

	log.Println("Database migrate up success")
	return nil
}

func migrationDown(cmd *cobra.Command, args []string) error {
	db, _ := New()
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://configs/db/migrate",
		"postgres", driver,
	)

	if err != nil {
		return err
	}

	errMig := mig.Steps(-1)

	if errMig != nil {
		return errMig
	}

	log.Println("Database rollback success")
	return nil
}
