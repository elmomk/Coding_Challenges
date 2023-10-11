/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LinesInFile(fileName string) int {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	// for index, lines := range result {
	// fmt.Println(index, lines)
	// }
	return len(result)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "ggwc -l <file>",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		b, err := cmd.Flags().GetBool("bytes")
		check(err)
		l, err := cmd.Flags().GetBool("lines")
		check(err)
		file := args[0]
		if b {
			fileinfo, err := os.Stat(file)
			check(err)
			fmt.Println(fileinfo.Size(), file)
		}
		if l {
			lines := LinesInFile(file)
			fmt.Println(lines, file)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("bytes", "c", false, "Show how many bytes are in a file")
	rootCmd.Flags().BoolP("lines", "l", false, "Show how many lines are in a file")
}
