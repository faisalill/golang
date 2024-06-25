/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// conNum is the connection number
var (
	conNum              int8
	connectionCount     int
	connectionNames     []string = []string{}
	connectionDatabases []string = []string{}
	connectionHosts     []string = []string{}
	connectionPorts     []int    = []int{}
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

				var connectionIter int = 1

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

	selectCmd := &cobra.Command{
		Use:   "select",
		Short: "Select a connection",
		Long:  `Select a connection from the list of connections.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Selecting connection number: ", conNum)
			s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
			s.Start()                                                    // Start the spinner
			viper.SetConfigName("connections.yaml")
			viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
			viper.AddConfigPath(".")    // optionally look for config in the working directory
			err := viper.ReadInConfig() // Find and read the config file
			if err != nil {             // Handle errors reading the config file
				panic(fmt.Errorf("fatal error config file: %w", err))
			}

			viper.SetConfigFile("pgsql.yaml")
			viper.SetConfigType("yaml")
			viper.AddConfigPath("#HOME/.psql/")
			viper.AddConfigPath(".")
			errr := viper.ReadInConfig()
			if errr != nil {
				panic(errr)
			}
			// yamlExample := []byte(`
			//       username: name
			//       password: pass
			//       database: dbname
			//       `)

			// viper.ReadConfig(bytes.NewBuffer(yamlExample))

			var connectionIter int = 1

			for viper.IsSet("username_"+strconv.Itoa(connectionIter)) && viper.IsSet("password_"+strconv.Itoa(connectionIter)) && viper.IsSet("database_"+strconv.Itoa(connectionIter)) && viper.IsSet("host_"+strconv.Itoa(connectionIter)) && viper.IsSet("port_"+strconv.Itoa(connectionIter)) {
				connectionNames = append(connectionNames, viper.Get("username_"+strconv.Itoa(connectionIter)).(string))
				connectionDatabases = append(connectionDatabases, viper.Get("database_"+strconv.Itoa(connectionIter)).(string))
				connectionHosts = append(connectionHosts, viper.Get("host_"+strconv.Itoa(connectionIter)).(string))
				connectionPorts = append(connectionPorts, viper.Get("port_"+strconv.Itoa(connectionIter)).(int))
				connectionCount += 1
				connectionIter += 1
			}

			if conNum > int8(connectionCount) {
				fmt.Println("Connection number does not exist")
			} else if conNum <= 0 {
				fmt.Println("Connection number must be greater than 0")
			} else {
				viper.Set("selected_connection", conNum)
			}
			s.Stop()
		},
	}

	selectCmd.Flags().Int8VarP(&conNum, "connection", "c", 0, "Connection Number")
	selectCmd.MarkFlagRequired("connection")

	connectionsCmd.AddCommand(listCmd)
	connectionsCmd.AddCommand(selectCmd)
}
