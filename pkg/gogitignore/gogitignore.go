package gogitignore

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	c "github.com/rexsimiloluwah/gogitignorecli/pkg/cache"
	"github.com/rexsimiloluwah/gogitignorecli/pkg/utils"

	"github.com/schollz/closestmatch"
)

var (
	cache                                = c.NewCache()
	GITIGNORE_API_BASE_URL               = "https://gitignore.io/api"
	GITIGNORE_INPUT_TYPES_LIST_CACHE_KEY = "gitignorelist"
	GITIGNORE_FILE_CACHE_KEY_PREFIX      = "gitignorefile"
	GITIGNORE_INPUT_TYPES_MAP_CACHE_KEY  = "gitignorecharmap"
)

// Fetches the list of the operating systems, programming languages and IDE input types
// from the gitignore API
func fetchGitignoreListFromServer() []string {
	listEndpoint := fmt.Sprintf("%s/list", GITIGNORE_API_BASE_URL)
	response, err := http.Get(listEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := strings.Replace(string(responseBytes), "\n", ",", -1)
	responseList := strings.Split(strings.Trim(responseString, "\n"), ",")

	return responseList
}

// Fetches the gitignore list of input types from the cache
func fetchGitignoreListFromCache() ([]string, bool) {
	data, found := cache.Get(GITIGNORE_INPUT_TYPES_LIST_CACHE_KEY)
	if !found {
		return nil, found
	}
	l := make([]string, len(data.([]interface{})))
	for i, v := range data.([]interface{}) {
		l[i] = fmt.Sprint(v)
	}

	return l, found
}

// Returns the list of gitignore input types
func FetchGitignoreList() []string {
	// fetch the list from the cache if it exists
	d, found := fetchGitignoreListFromCache()
	if found {
		return d
	}
	// fetch the list from the server if it does not exist in the cache
	gitignoreList := fetchGitignoreListFromServer()
	// add to cache
	_ = cache.Set(GITIGNORE_INPUT_TYPES_LIST_CACHE_KEY, gitignoreList)
	return gitignoreList
}

// Check if a gitignore input type exists
func CheckGitignoreInputTypeExists(inputType string, gitignoreList []string) bool {
	gitignoreListMap := make(map[string]bool)
	for _, v := range gitignoreList {
		gitignoreListMap[v] = true
	}

	return gitignoreListMap[inputType]
}

// // Fetch the gitignore file content for a specific input type from server
func fetchGitignoreFileContentFromServer(inputType string) string {
	listEndpoint := fmt.Sprintf("%s/%s", GITIGNORE_API_BASE_URL, inputType)
	response, err := http.Get(listEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseBytes)
}

// // Fetch the gitignore file content for a specific input type from cache
func fetchGitignoreFileContentFromCache(inputType string) (string, bool) {
	cacheKey := fmt.Sprintf("%s:%s", GITIGNORE_FILE_CACHE_KEY_PREFIX, inputType)
	data, found := cache.Get(cacheKey)
	if !found {
		return "", found
	}

	return data.(string), found
}

// Fetch the gitignore file content for a specific input type
func FetchGitignoreFileContent(inputType string) string {
	// fetch the list from the cache if it exists
	d, found := fetchGitignoreFileContentFromCache(inputType)
	if found {
		return d
	}
	// fetch the list from the server if it does not exist in the cache
	content := fetchGitignoreFileContentFromServer(inputType)
	cacheKey := fmt.Sprintf("%s:%s", GITIGNORE_FILE_CACHE_KEY_PREFIX, inputType)
	_ = cache.Set(cacheKey, content)
	return content
}

// Merge gitignore file content for multiple input types
func MergeGitignoreFileContent(inputTypes ...string) string {
	var result []string
	for _, inputType := range inputTypes {
		gitignoreFileContent := FetchGitignoreFileContent(inputType)
		result = append(result, fmt.Sprintf("#%s", inputType))
		result = append(result, gitignoreFileContent)
	}
	return strings.Join(result, "\n")
}

// Returns the closest match to the passed input type using fuzzy matching
func GetClosestInputTypeMatch(inputType string, gitignoreList []string) []string {
	bagSizes := []int{2}
	cm := closestmatch.New(gitignoreList, bagSizes)
	closestMatches := cm.ClosestN(inputType, 3)
	return closestMatches
}

// Returns the input types arranged based on their first character
func GetInputTypesCharMap() map[string][]string {
	charMap := make(map[string][]string)
	// check if it exists in cache
	data, found := cache.Get(GITIGNORE_INPUT_TYPES_MAP_CACHE_KEY)
	if found {
		for k, v := range data.(map[string]interface{}) {
			t := v.([]interface{})
			data := make([]string, len(t))
			for i, s := range t {
				data[i] = fmt.Sprint(s)
			}
			charMap[k] = data
		}
		return charMap
	}
	gitignoreInputTypes := FetchGitignoreList()
	alphabetLetters := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, char := range alphabetLetters {
		charMap[char] = make([]string, 0)
	}

	for _, char := range alphabetLetters {
		for _, inputType := range gitignoreInputTypes {
			c := string(inputType[0]) // first letter
			if c == char {
				charMap[c] = append(charMap[c], inputType)
			}
		}
	}

	//cache the result
	_ = cache.Set(GITIGNORE_INPUT_TYPES_MAP_CACHE_KEY, charMap)
	return charMap
}

// Returns the input types that start with a specific character
func GetInputTypesStartsWith(inputTypes []string, char string) []string {
	firstLetter := string(char[0])
	return utils.Filter(inputTypes, func(s string) bool {
		return string(s[0]) == firstLetter
	})
}
