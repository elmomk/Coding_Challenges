package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// CountItemsInFile counts the number of items (lines, bytes, words, or characters) in a file based on the given split type.
func CountItemsInFile(file *bufio.Reader, splitType bufio.SplitFunc) int {
	scanner := bufio.NewScanner(file)
// func CountItemsInFile(scanner *bufio.Scanner, splitType bufio.SplitFunc) int {
//   fmt.Print(scanner)
	scanner.Split(splitType)

	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Panicln(err)
	}
	return count
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "ggwc <file>",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := cmd.Flags().GetBool("bytes")
		Check(err)
		l, err := cmd.Flags().GetBool("lines")
		Check(err)
		w, err := cmd.Flags().GetBool("words")
		Check(err)
		c, err := cmd.Flags().GetBool("chars")
		Check(err)

		fi, err := os.Stdin.Stat()
		Check(err)

		if fi.Mode()&os.ModeNamedPipe != 0 {
      fileopen := bufio.NewReader(os.Stdin)
			// filebuf := bufio.NewReader(os.Stdin)
      // fileopen := bufio.NewScanner(filebuf)
			if w {
				file_words := CountItemsInFile(fileopen, bufio.ScanWords)
				fmt.Println(file_words)
			} else if c {
				file_chars := CountItemsInFile(fileopen, bufio.ScanRunes)
				fmt.Println(file_chars)
			} else if l {
				file_lines := CountItemsInFile(fileopen, bufio.ScanLines)
				fmt.Println(file_lines)
			} else if b {
				file_bytes := CountItemsInFile(fileopen, bufio.ScanBytes)
				fmt.Println(file_bytes)
			} else {
				// bug file_lines and file_words are always 0
				file_bytes := CountItemsInFile(fileopen, bufio.ScanBytes)
        // fmt.Println(file_bytes)
				file_lines := CountItemsInFile(fileopen, bufio.ScanLines)
        // fmt.Println(file_lines)
				file_words := CountItemsInFile(fileopen, bufio.ScanWords)
        // fmt.Println(file_words)
				fmt.Println(file_lines, file_words, file_bytes)
			}
		}

		for _, file := range args {
			f, err := os.Open(file)
			Check(err)
			defer f.Close()
			fileopen := bufio.NewReader(f)
			// filebuf := bufio.NewReader(f)
      // fileopen := bufio.NewScanner(filebuf)

			if b {
				file_bytes := CountItemsInFile(fileopen, bufio.ScanBytes)
				fmt.Println(file_bytes, file)
			} else if l {
				file_lines := CountItemsInFile(fileopen, bufio.ScanLines)
				fmt.Println(file_lines, file)
			} else if w {
				file_words := CountItemsInFile(fileopen, bufio.ScanWords)
				fmt.Println(file_words, file)
			} else if c {
				file_chars := CountItemsInFile(fileopen, bufio.ScanRunes)
				fmt.Println(file_chars, file)
			} else {
				file_bytes := CountItemsInFile(fileopen, bufio.ScanBytes)
				file_lines := CountItemsInFile(fileopen, bufio.ScanLines)
				file_words := CountItemsInFile(fileopen, bufio.ScanWords)
				fmt.Println(file_lines, file_words, file_bytes, file)
			}
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
	rootCmd.Flags().BoolP("words", "w", false, "Show how many words are in a file")
	rootCmd.Flags().BoolP("chars", "m", false, "Show how many characters are in a file")
}
