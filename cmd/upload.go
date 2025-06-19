/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	filePathOptionName       = "file-path"
	filePathShortOptionName  = "p"
	chunkSizeOptionName      = "chunk-size"
	chunkSizeShortOptionName = "c"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file and split the file into chunks",
	Long: `Upload command is going to get the uploaded file
	and split the file into chunks. In recent, the chunked data
	can only save into disk files.
	You can use options like
	--file-path to define specific file path.
	Or 
	--chunk-size to decide the size to split the file. ex. 1M, 1k, 256k,...`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")

		// Get flags
		filepath, err := cmd.Flags().GetString(filePathOptionName)
		if err != nil {
			fmt.Printf("Error getting %s option: %s", filePathOptionName, err.Error())
			return
		}

		fmt.Println("filepath:", filepath)

		chunkSize, err := cmd.Flags().GetString(chunkSizeOptionName)
		if err != nil {
			fmt.Printf("Error getting %s option: %s", chunkSizeOptionName, err.Error())
			return
		}

		fmt.Println("chunkSize:", chunkSize)

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	uploadCmd.Flags().StringP(filePathOptionName, filePathShortOptionName, "", "File path to upload.")
	uploadCmd.MarkFlagRequired(filePathOptionName)
	uploadCmd.Flags().StringP(chunkSizeOptionName, chunkSizeShortOptionName, "256k", "Chunk size to split file. Default is 256kB.")
}
