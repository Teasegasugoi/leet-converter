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
	Short: "Display of leet table",
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
