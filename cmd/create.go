/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var leetTable map[string][]string

// Used for flags
var (
	min int
	max int
	num int
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
		} else if num < 1 {
			return fmt.Errorf("num must be greater than 0")
		}
		for i := 0; i < num; i++ {
			create(args[0])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&min, "min", "m", 1, "min")
	createCmd.Flags().IntVarP(&max, "max", "M", 100, "max")
	createCmd.Flags().IntVarP(&num, "num", "n", 1, "num")

	// JSONファイルの読み込み
	file, err := os.ReadFile("leet_table.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &leetTable)
	if err != nil {
		log.Fatal(err)
	}
}

func create(name string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 変換できる箇所を探す
	var c []int
	lowName := strings.ToLower(name)
	for i := 0; i < len(lowName); i++ {
		if l, ok := leetTable[string(lowName[i])]; ok && len(l) > 0 {
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
		rn := r.Intn(len(leetTable[string(lowName[c[i]])]))
		s[c[i]] = leetTable[string(lowName[c[i]])][rn]
	}

	fmt.Println(strings.Join(s, ""))
}
