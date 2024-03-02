package utils

import (
	"net/http"
	"os"
	"strings"
)

// Found this idea in a wonderful blog post about preventing . and .. file access
// https://crazcalm.github.io/blog/post/custom_file_server/

func isIllegal(name string) bool {
	for _, substring := range strings.Split(name, "/") {
		if strings.HasPrefix(substring, ".") {
			return true
		}
	}
	return false
}

type SecureFileSystem struct {
	http.FileSystem
}

func (fs SecureFileSystem) Open(name string) (http.File, error) {
	file, err := fs.FileSystem.Open(name)

	if isIllegal(name) {
		return nil, os.ErrPermission
	}

	return file, err
}
