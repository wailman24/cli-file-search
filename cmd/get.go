/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		dirPath := "C:\\Users\\Asus\\OneDrive\\Desktop\\wscan"
		chtext := make(chan string)
		chfiles := make(chan []string)

		go service.ListFiles(dirPath, chfiles)
		go service.ReadFiles(chfiles, chtext)
		for text := range chtext {
			fmt.Printf("line: %s\n", text)
		}

		jokeTerm, _ := cmd.Flags().GetString("regex")

		if jokeTerm != "" {
			fmt.Printf("hello man %s", jokeTerm)
		} else {
			fmt.Printf("hello man 2 %s", jokeTerm)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().String("regex", "r", "A search term for a dad joke.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
