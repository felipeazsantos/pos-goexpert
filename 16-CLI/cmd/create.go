/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/felipeazsantos/pos-goexpert/16-CLI/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	RunE:  runCreate(GetCategoryDB()),
}

func runCreate(categoryDB *database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDB.Create(name, description)

		return err
	}
}

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category to create")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
