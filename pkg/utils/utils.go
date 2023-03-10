package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func WriteContentToFile(filename string, outdir string, content string) error {
	filePath := filepath.Join(outdir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(filePath), 0700)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, _ = w.WriteString(content)
	return nil
}

func Filter(arr []string, fn func(s string) bool) []string {
	result := make([]string, 0)
	for _, el := range arr {
		if fn(el) {
			result = append(result, el)
		}
	}
	return result
}
