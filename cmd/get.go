/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wailman24/cli-file-search.git/internal/service"
)

type data struct {
	File  string   `json:"file_path"`
	Numl  int      `json:"line_num"`
	Match []string `json:"matched"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get all files of a dir",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		infof := service.InfoFile{}
		//dirPath := "C:\\Users\\Asus\\OneDrive\\Desktop\\wscan"
		chtext := make(chan service.InfoFile)
		chfiles := make(chan []string)
		rgx, _ := cmd.Flags().GetString("regex")
		dir, err := cmd.Flags().GetString("dir")
		ext, _ := cmd.Flags().GetString("ext")
		ignore, _ := cmd.Flags().GetString("ignore")
		if err != nil {
			fmt.Printf("please provide directory to scan")
		}
		if !cmd.Flags().Lookup("regex").Changed {
			fmt.Print("please provide your regex flag --regex=\"exmpl\"")
			os.Exit(1)
		}

		if !cmd.Flags().Lookup("ext").Changed {
			ext = ""
		}

		go service.ListFiles(dir, chfiles, ext, ignore)
		go infof.ReadFiles(chfiles, chtext)
		//infof.ReadFiles
		r, err := regexp.Compile(rgx)
		if err != nil {
			fmt.Printf("please enter a valid regex")
			os.Exit(1)
		}
		for text := range chtext {
			//fmt.Printf("line: %s\n", text)
			if rgx != "" {
				if r.FindAllString(text.Line, -1) != nil {
					//fmt.Printf("file: %s  Line: %d ", text.File, text.NumL)
					//fmt.Println(r.FindAllString(text.Line, -1))
					result := data{
						File:  text.File,
						Numl:  text.NumL,
						Match: r.FindAllString(text.Line, -1),
					}
					printColored(result)
				}
				//fmt.Printf("hello %s", text)
			} else {
				fmt.Printf("the flag value is empty")
			}
		}

	},
}

func printColored(result data) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("{\n")
	fmt.Printf("  \"File\": \"%s\",\n", green(result.File))
	fmt.Printf("  \"Num of line\": %s,\n", yellow(result.Numl))
	fmt.Printf("  \"Match\": [")
	for i, match := range result.Match {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("\"%s\"", red(match))
	}
	fmt.Println("]\n}")
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().String("dir", "d", "enter the dir you to search in")
	getCmd.PersistentFlags().String("regex", "r", "enter the regex you are looking for")
	getCmd.PersistentFlags().String("ext", "e", "enter the extention you to scan")
	getCmd.PersistentFlags().String("ignore", "i", "enter the extention or dir you want to ignore")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
