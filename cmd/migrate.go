// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate [direction]",
	Short: "perform a database migration.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASS"),
			os.Getenv("PGHOST"),
			os.Getenv("PGPORT"),
			os.Getenv("PGDATABASE"),
		)
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)

		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			viper.GetString("dir"),
			"postgres", driver)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)

		}
		direction := args[0]
		switch direction {
		case "":
		case "up":
			err = m.Up()
			break
		case "down":
			err = m.Down()
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)

		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migrateCmd.Flags().String("dir", "file:///migrations", "database migrations directory")
	viper.BindPFlags(migrateCmd.Flags())
}
