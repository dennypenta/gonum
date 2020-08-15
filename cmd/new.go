package cmd

import (
	"github.com/dennypenta/gnum/cmd/enums"
	"github.com/spf13/cobra"
	"log"
)

var _type, filepath, _package string
var values []string

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&_type, "type", "t", "", "type name (required)")
	newCmd.MarkFlagRequired("type")

	newCmd.Flags().StringVarP(&filepath, "output", "o", "", "file path where must be pasted generated code")
	newCmd.MarkFlagRequired("output")

	newCmd.Flags().StringVarP(&_package, "package", "p", "", "package name for generated file (required)")
	newCmd.MarkFlagRequired("package")

	newCmd.Flags().StringSliceVarP(&values, "values", "v", nil, "possible values for enum (required)")
	newCmd.MarkFlagRequired("values")

	// TODO: add optional withType option for names
	// directory
	// readme
	// install
}

type Params struct {
	Type string
	Values []string
	Filepath string
	Package string
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "generate new iota type",
	Run: func(cmd *cobra.Command, args []string) {
		enum := &enums.Enum{}
		renderArgs := enums.RenderArgs{
			Package: _package,
			Type: _type,
			Values: values,
			Path: filepath,
		}
		if err := enum.Render(renderArgs); err != nil {
			log.Fatal(err)
		}
	},
}
