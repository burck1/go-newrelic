package cmd

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const apiRoot = "https://api.newrelic.com"
const definitions = "/v2/definitions.json"

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
	definitionsPath := filepath.Join(rootPath, "definitions.json")

	client := initClient()

	url := apiRoot + definitions
	if err := downloadJSONFile(client, url, definitionsPath, raw); err != nil {
		return err
	}
	log.Printf("GET %s\n", url)

	data, err := readJSONObjectFile(definitionsPath)
	if err != nil {
		return err
	}

	apis, err := getApis(data)
	if err != nil {
		return err
	}

	paths := make([]string, len(apis), len(apis))
	for i, api := range apis {
		paths[i], err = getPath(i, api)
		if err != nil {
			return err
		}
	}

	// fetch all the files in parallel
	// create some channels so we can wait on the responses
	successc, errc := make(chan string), make(chan error)
	for _, path := range paths {
		go func(path string) {
			url := getURLFromPath(path)
			filename, err := getFilenameFromURL(url)
			if err != nil {
				errc <- err
				return
			}
			filePath := filepath.Join(rootPath, filename)
			if err := downloadJSONFile(client, url, filePath, raw); err != nil {
				errc <- err
				return
			}
			successc <- url
		}(path)
	}

	// wait for all the downloads to complete
	// collect the errors for later
	var errs []error
	for i := 0; i < len(paths); i++ {
		select {
		case url := <-successc:
			log.Printf("GET %s\n", url)
		case err := <-errc:
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		// just return the first error
		return errs[0]
	}

	// log.Println("End pull")

	return nil
}

func downloadJSONFile(client http.Client, url string, path string, raw bool) error {
	request, err := initRequest("GET", url, nil)
	if err != nil {
		return err
	}

	response, err := doWithRetry(client, request, 5, time.Second)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if raw {
		if err := writeToFile(response.Body, path); err != nil {
			return err
		}
	} else {
		if err := writeToFileIndent(response.Body, path); err != nil {
			return err
		}
	}

	return nil
}

func getURLFromPath(path string) string {
	return apiRoot + strings.Replace(path, "{format}", "json", 1)
}

func initPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	rootPath := filepath.Join(cwd, "definitions-1.1")

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
