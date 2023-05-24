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
	"a": "4",
	"b": "",
	"c": "",
	"d": "",
	"e": "3",
	"f": "",
	"g": "6",
	"h": "",
	"i": "1",
	"j": "",
	"k": "",
	"l": "1",
	"m": "",
	"n": "",
	"o": "0",
	"p": "9",
	"q": "",
	"r": "",
	"s": "5",
	"t": "7",
	"u": "",
	"v": "",
	"w": "",
	"x": "",
	"y": "",
	"z": "2",
}

// Used for flags
var (
	min int
	max int
)

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
		} else if max < 1 || min < 1 {
			return fmt.Errorf("min and max must be greater than 0")
		} else if max < min {
			return fmt.Errorf("max must be greater than min")
		} else if len(args[0]) < min {
			return fmt.Errorf("min must be smaller than arg's length")
		}
		create(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&min, "min", "m", 1, "min")
	createCmd.Flags().IntVarP(&max, "max", "M", 100, "max")
}

func create(name string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 変換できる箇所を探す
	var c []int
	l := strings.ToLower(name)
	for i := 0; i < len(l); i++ {
		if l := leetTable[string(l[i])]; l != "" {
			c = append(c, i)
		}
	}

	// shuffle slice
	r.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	// 変更する文字数設定
	if max > len(c) {
		max = len(c)
	}
	if min > len(c) {
		min = len(c)
	}
	var n int = r.Intn(max-min+1) + min
	c = c[:n]

	// Convert to Leet
	s := strings.Split(name, "")
	for i := 0; i < len(c); i++ {
		s[c[i]] = leetTable[string(l[c[i]])]
	}

	fmt.Println(strings.Join(s, ""))
}
