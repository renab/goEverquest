package everquest

import (
	"os"
	"path/filepath"
	"strings"
)

func getRecentRosterDump(path string, guildName string) string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(filepath.Base(path), guildName) {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files[len(files)-1] // return last file - should be latest
}
