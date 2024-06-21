/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
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

				var connectionNames []string = []string{}
				var connectionDatabases []string = []string{}
				var connectionHosts []string = []string{}
				var connectionPorts []int = []int{}

				for viper.IsSet("username_"+strconv.Itoa(connectionIter)) && viper.IsSet("password_"+strconv.Itoa(connectionIter)) && viper.IsSet("database_"+strconv.Itoa(connectionIter)) && viper.IsSet("host_"+strconv.Itoa(connectionIter)) && viper.IsSet("port_"+strconv.Itoa(connectionIter)) {
					connectionNames = append(connectionNames, viper.Get("username_"+strconv.Itoa(connectionIter)).(string))
					connectionDatabases = append(connectionDatabases, viper.Get("database_"+strconv.Itoa(connectionIter)).(string))
					connectionHosts = append(connectionHosts, viper.Get("host_"+strconv.Itoa(connectionIter)).(string))
					connectionPorts = append(connectionPorts, viper.Get("port_"+strconv.Itoa(connectionIter)).(int))
					connectionCount += 1
					connectionIter += 1
				}

				connectionsTable := table.NewWriter()
				connectionsTable.SetOutputMirror(os.Stdout)
				connectionsTable.AppendHeader(table.Row{"#", "Username", "Database", "Host", "Port"})

				for index, val := range connectionNames {
					connectionsTable.AppendRows([]table.Row{
						{index + 1, val, connectionDatabases[index], connectionHosts[index], connectionPorts[index]},
					})
					connectionsTable.AppendSeparator()
				}

				connectionsTable.SetStyle(table.StyleLight)
				connectionsTable.Render()

			}
		},
	}

	connectionsCmd.AddCommand(listCmd)
}
