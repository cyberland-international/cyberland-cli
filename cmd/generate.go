//Package cmd
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var (
	useCaseNumber    int
	projectName      string
	templateFilePath string
	TemplateFs       embed.FS
	renderedTemplate bytes.Buffer
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long:  `This is long generate command description`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Use case number: ", useCaseNumber)
		fmt.Println("Project name: ", projectName)

		// Craft the template file path and name
		templateFilePath = fmt.Sprintf("templates/use_case_%d_deploy.yaml", useCaseNumber)
		templateFileName := fmt.Sprintf("use_case_%d_deploy.yaml", useCaseNumber)

		// Reads the template file from the embedded file system
		tplFile, err := TemplateFs.ReadFile(templateFilePath)
		//fmt.Println(string(tplFile))

		// Prepare context (data) for the template
		type TemplateData struct {
			ProjectName string
		}
		td := TemplateData{ProjectName: projectName}

		// Render the template
		tpl, err := template.New("template").Delims("%%%", "%%%").Parse(string(tplFile))
		if err != nil {
			panic(err)
		}
		if err := tpl.Execute(&renderedTemplate, td); err != nil {
			panic(err)
		}

		// Get current working directory
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// Write the rendered template to a file
		renderedTemplateFilePath := fmt.Sprintf("%s/%s", wd, templateFileName)
		if err := os.WriteFile(
			renderedTemplateFilePath,
			renderedTemplate.Bytes(),
			0644); err != nil {
			panic(err)
		}

		// Print the location of the rendered template file
		fmt.Printf("Rendered template file: %s", renderedTemplateFilePath)

	},
}

func init() {
	generateCmd.Flags().IntVarP(&useCaseNumber, "use-case", "u", 0, "Use case number")
	generateCmd.Flags().StringVarP(&projectName, "project-name", "p", "",
		"Name of the project")

	if err := generateCmd.MarkFlagRequired("use-case"); err != nil {
		fmt.Println(err)
	}
	if err := generateCmd.MarkFlagRequired("project-name"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(generateCmd)
}
