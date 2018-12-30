package walker

import (
	"os"
	"path/filepath"
)

func resolveFilePaths(path string) ([]string, error) {
	allPaths := make([]string, 0, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		allPaths = append(allPaths, path)
		return nil
	})
	return allPaths, err
}
