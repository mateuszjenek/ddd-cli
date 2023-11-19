/*
Copyright © 2023 Mateusz Jenek <mateusz.jenek@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/mateuszjenek/ddd-cli/internal/application"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure/port"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dbPort, err := port.NewSQLiteDatabasePort("./notes.db")
		if err != nil {
			return fmt.Errorf("failed to create a database port: %w", err)
		}
		defer dbPort.Close()

		repository := infrastructure.NewNoteLocalRepository(dbPort)
		useCase := application.NewListNotesUseCase(repository)
		result, err := useCase.ListNotes()
		if err != nil {
			return fmt.Errorf("create note use case returned an error: %v", err)
		}

		fmt.Println("-- List of notes --")
		for i, note := range result.Notes {
			if i != 0 {
				fmt.Println("-------------------")
			}
			fmt.Println("ID: ", note.Id)
			fmt.Println("Author: ", note.Author)
			fmt.Println("Message: ", note.Message)
		}
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
