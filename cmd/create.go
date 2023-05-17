/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var leetTable = map[string]string{
	"A": "4",
	"B": "",
	"C": "",
	"D": "",
	"E": "3",
	"F": "",
	"G": "6",
	"H": "",
	"I": "1",
	"J": "",
	"K": "",
	"L": "1",
	"M": "",
	"N": "",
	"O": "0",
	"P": "9",
	"Q": "",
	"R": "",
	"S": "5",
	"T": "7",
	"U": "",
	"V": "",
	"W": "",
	"X": "",
	"Y": "",
	"Z": "2",
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing argument")
		} else if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}
		create(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(name string) {
	// 変換できる箇所を探す
	var c []int
	for i := 0; i < len(name); i++ {
		if l := leetTable[string(name[i])]; l != "" {
			c = append(c, i)
		}
	}
	// 変更する文字数設定
	var num int
	if len(c) > 5 {
		num = 5
	} else {
		num = len(c)
	}
	// shuffle slice
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(c); i > 0; i-- {
		j := rand.Intn(i)
		c[i-1], c[j] = c[j], c[i-1]
	}
	c = c[:num]

	// Convert to Leet
	s := strings.Split(name, "")
	for i := 0; i < len(c); i++ {
		s[c[i]] = leetTable[string(name[c[i]])]
	}

	fmt.Println(strings.Join(s, ""))
}
