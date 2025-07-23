/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/wailman24/cli-file-search.git/internal/service"
)

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
		dirPath := "C:\\Users\\Asus\\OneDrive\\Desktop\\wscan"
		chtext := make(chan service.InfoFile)
		chfiles := make(chan []string)
		rgx, _ := cmd.Flags().GetString("regex")

		go service.ListFiles(dirPath, chfiles)
		go infof.ReadFiles(chfiles, chtext)
		//infof.ReadFiles
		for text := range chtext {
			//fmt.Printf("line: %s\n", text)
			r, _ := regexp.Compile(rgx)
			if rgx != "" {
				fmt.Printf("file: %s ", text.File)
				fmt.Println(r.FindAllString(text.Line, -1))

				//fmt.Printf("hello man %s", text)
			} else {
				fmt.Printf("the flag value is empty")
			}
		}

		/* 	line, _ := cmd.Flags().GetString("regex")
		r, _ := regexp.Compile("w([a-z]+)ch")
		if line != "" {
			fmt.Println(r.FindAllString(line, -1))
			fmt.Printf("hello man %s", line)
		} else {
			fmt.Printf("the flag value is empty")
		} */
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	//getCmd.PersistentFlags().String("dir", "d", "A search term for a dad joke.")
	getCmd.PersistentFlags().String("regex", "r", "enter the regex you are looking for")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
