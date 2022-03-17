/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"

	"log"

	"strings"

	"io"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ToM4A",
	Short: "Change MP4 to M4A",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		files, err := os.ReadDir(".");

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			if len(file.Name()) > 4 && strings.Contains(file.Name(), ".mp4") {
				splitIndex := len(file.Name()) - 4 //".mp4"
				name := file.Name()[:splitIndex]
				fin, err := os.Open(file.Name())
				if err != nil {
					//TODO: log error
					continue
				}
				m4aFullName := fmt.Sprintf("%s.m4a", name)
				fout, err := os.Create(m4aFullName)
				if err != nil {
					//TODO: log error
					fin.Close() //TODO refactor into function to simplify with defer?
					continue
				}

				if _, err := io.Copy(fout, fin); err != nil { // check file sizes match?
					//TODO: log error
					fin.Close() //TODO refactor into function to simplify with defer?
					fout.Close()
					continue
				}

				fin.Close() //TODO err check?
				fout.Close()
			}
			//TODO: log skip if verbose?
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ToM4A.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


