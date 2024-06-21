/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// connectionsCmd represents the connections command
var connectionsCmd = &cobra.Command{
	Use:   "connections",
	Short: "Manage database connections",
	Long:  `Manage database connections like creating a connection, listing connections, updating a connection, and deleting a connection.`,
}

func init() {
	rootCmd.AddCommand(connectionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add subcommands
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all connections",
		Long:  `List all connections.`,
		Run: func(cmd *cobra.Command, args []string) {
			viper.SetConfigFile("pgsql.yaml")
			viper.SetConfigType("yaml")
			viper.AddConfigPath("#HOME/.psql/")
			viper.AddConfigPath(".")
			err := viper.ReadInConfig()
			if err != nil {
				panic(err)
			} else {
				// yamlExample := []byte(`
				//       username: name
				//       password: pass
				//       database: dbname
				//       `)

				// viper.ReadConfig(bytes.NewBuffer(yamlExample))

				var connectionCount int
				var connectionIter int = 1

				fmt.Println(viper.IsSet("username_" + strconv.Itoa(connectionIter)))
				fmt.Println(viper.IsSet("password_" + strconv.Itoa(connectionIter)))
				fmt.Println(viper.IsSet("database_" + strconv.Itoa(connectionIter)))

				for viper.IsSet("username_"+strconv.Itoa(connectionIter)) && viper.IsSet("password_"+strconv.Itoa(connectionIter)) && viper.IsSet("database_"+strconv.Itoa(connectionIter)) {
					connectionCount += 1
					connectionIter += 1
				}

				fmt.Println("Connection count: ", connectionCount)

			}
		},
	}

	connectionsCmd.AddCommand(listCmd)
}
