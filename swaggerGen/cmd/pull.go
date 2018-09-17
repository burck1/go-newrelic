package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the latest swagger files from New Relic",
	Long: `This command will get the latest swagger files from
the New Relic API Explorer and store them in the current
directory. It will use https://api.newrelic.com/v2/definitions.json
as the root swagger definition.`,
	RunE: pullLatestSwaggerDocs,
}

func init() {
	rand.Seed(time.Now().UnixNano())

	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolP("raw", "r", false, "Disable json formatting and indentation for the definition files")
}

func pullLatestSwaggerDocs(cmd *cobra.Command, args []string) error {
	// log.Println("Start pull")
	// cmd.DebugFlags()
	raw, err := cmd.Flags().GetBool("raw")
	if err != nil {
		return err
	}

	rootPath, err := initPath()
	if err != nil {
		return err
	}

	client := initClient()
	request, err := initRequest("GET", "https://api.newrelic.com/v2/definitions.json", nil)
	if err != nil {
		return err
	}

	response, err := doWithRetry(client, request, 5, time.Second)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	definitionsPath := filepath.Join(rootPath, "definitions.json")

	if raw {
		if err := writeToFile(response.Body, definitionsPath); err != nil {
			return err
		}
	} else {
		if err := writeToFileIndent(response.Body, definitionsPath); err != nil {
			return err
		}
	}

	// log.Println("End pull")

	return nil
}

func initPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootPath := filepath.Join(cwd, "definitions")

	if err := rmdir(rootPath); err != nil {
		return rootPath, err
	}

	if err := os.Mkdir(rootPath, os.ModePerm); err != nil {
		return rootPath, err
	}

	return rootPath, nil
}

func initClient() http.Client {
	return http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}

func initRequest(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return request, err
	}

	request.Header.Add("User-Agent", "GoNewRelic")
	request.Header.Add("Accept", "application/json")
	if body != nil {
		request.Header.Add("Content-Type", "application/json")
	}

	return request, nil
}

func writeToFile(body io.ReadCloser, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, body); err != nil {
		return err
	}

	return nil
}

func writeToFileIndent(body io.ReadCloser, path string) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var data interface{}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return err
	}

	json, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, json, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Remove a directory and all files in it if it exists
func rmdir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		dirRead, err := os.Open(path)
		if err != nil {
			return err
		}

		dirFiles, err := dirRead.Readdir(0)
		if err != nil {
			return err
		}

		for index := range dirFiles {
			fileHere := dirFiles[index]
			nameHere := fileHere.Name()
			fullPath := filepath.Join(path, nameHere)
			if err := os.Remove(fullPath); err != nil {
				return err
			}
		}

		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

// DoWithRetry makes a http request with automatic retries
// and exponential backoff
func doWithRetry(client http.Client, request *http.Request, attempts int, sleep time.Duration) (*http.Response, error) {
	response, shouldRetry, err := doForRetry(client, request)
	if shouldRetry {
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			log.Printf("Request failed. Waiting %.0f seconds\n", sleep.Seconds())
			time.Sleep(sleep)
			log.Println("Retrying request")
			return doWithRetry(client, request, attempts, 2*sleep)
		}
		return response, err
	}
	if stopErr, ok := err.(stop); ok {
		return response, stopErr.error
	}
	return response, err
}

func doForRetry(client http.Client, request *http.Request) (*http.Response, bool, error) {
	response, err := client.Do(request)
	if err != nil {
		return response, true, err
	}

	statusCode := response.StatusCode
	switch {
	case statusCode >= 500:
		return response, true, nil
	case statusCode >= 400:
		return response, false, stop{fmt.Errorf("client error: %v", statusCode)}
	default:
		return response, false, nil
	}
}

type stop struct {
	error
}
