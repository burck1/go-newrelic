package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// https://rpm.newrelic.com/api/explore/api/explore
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md

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
	// cmd.DebugFlags()

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootPath := filepath.Join(cwd, "definitions")
	definitionsPath := filepath.Join(rootPath, "definitions.json")

	data, err := readJSONObjectFile(definitionsPath)
	if err != nil {
		return err
	}

	apis, err := getApis(data)
	if err != nil {
		return err
	}

	for i, api := range apis {
		path, err := getPath(i, api)
		if err != nil {
			return err
		}
		path = strings.Replace(path, "{format}", "json", 1)
		filename, err := getFilenameFromURL(path)
		if err != nil {
			return err
		}
		filePath := filepath.Join(rootPath, filename)

		_, err = readJSONObjectFile(filePath)
		if err != nil {
			return err
		}

		log.Println(filePath)
	}

	log.Println("End gen")
	return nil
}
