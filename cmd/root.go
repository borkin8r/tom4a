/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"

	"strings"

	"io"

	"path/filepath"

	"io/fs"
)

// rootCmd represents the base command when called without any subcommands
var (
		Recursive bool
		FolderPath string

		rootCmd = &cobra.Command{
			Use:   "tom4a [Options]",
			Short: "Change MP4 to M4A",
			Long: `Creates a .M4A files by copying .MP4 files
			example:
			tom4a

			options:

			-r search subdirectories and make .M4A files there as well
			-p <path> the path to the folder containing the .mp4 files

			`,
			// Uncomment the following line if your bare application
			// has an action associated with it:
			Run: func(cmd *cobra.Command, args []string) {

				err := filepath.Walk(FolderPath, func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
						return err
					}
					if info.IsDir() && !Recursive {
						fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
						return filepath.SkipDir
					}

					files, err := os.ReadDir(path)

					if err != nil {
						fmt.Println(err)
					}

					for _, dirEntry := range files { //TODO use errgrp pkg to wait for all goroutines
						if len(dirEntry.Name()) > 4 && strings.Contains(dirEntry.Name(), ".mp4") {
							ToM4A(dirEntry)
						} else {
							fmt.Println(fmt.Sprintf("skipped %s\n", dirEntry.Name()))
						}
					}
					fmt.Printf("visited file or dir: %q\n", path)
					return nil
				})
				if err != nil {
					fmt.Printf("error walking the path: %v\n", err)
					return
				}
			},
		}
)

func ToM4A(dirEntry os.DirEntry) {
	splitIndex := len(dirEntry.Name()) - 4 //".mp4"
	name := dirEntry.Name()[:splitIndex]
	fin, err := os.Open(dirEntry.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fin.Close() //TODO refactor into function to simplify with defer?

	m4aFullName := fmt.Sprintf("%s.m4a", name)
	fout, err := os.Create(m4aFullName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	if _, err := io.Copy(fout, fin); err != nil { // check file sizes match?
		fmt.Println(err)
		return
	}
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
	rootCmd.Flags().StringVarP(&FolderPath, "folder path", "p", ".", "path to target folder to run in")
	rootCmd.Flags().BoolVarP(&Recursive, "recursive", "r", false, "iterate over subdirectories")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
