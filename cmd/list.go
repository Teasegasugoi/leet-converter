/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.ReadFile("leet_table.json")
		if err != nil {
			log.Fatal(err)
		}
		var data interface{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Fatal(err)
		}
		jsonString, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonString))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
