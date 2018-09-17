package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a swagger configuration from local JSON files",
	Long: `This command will read swagger files from the New Relic API Explorer
and generate a single swagger.json file from them.`,
	RunE: generateSwaggerDocs,
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func generateSwaggerDocs(cmd *cobra.Command, args []string) error {
	log.Println("Start gen")
	cmd.DebugFlags()
	log.Println("End gen")
	return nil
}
