package fakeUserAgent

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
	"embed"
)

//go:embed src/*
var embeddedFiles embed.FS

func getPackageFilePath(filename string) string {
	return filepath.Join("src", filename)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Filter[E any](s *[]E, f func(E) bool) []E {
	s2 := make([]E, 0, len(*s))
	for _, e := range *s {
		if f(e) {
			s2 = append(s2, e)
		}
	}
	return s2
}

func randFromLen(n int) int {
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)
	return r.Intn(n)
}

func ExtractMajorVersion(version string) int {
	parts := strings.Split(version, ".")
	if len(parts) > 0 {
		ver, _ := strconv.Atoi(parts[0])
		return ver
	}
	return 0
}

func loadFile() (io.Reader, error) {
	return embeddedFiles.Open(getPackageFilePath(userAgentsFile))
}

func getUserAgents(r io.Reader) (*[]UserAgents, error) {
	var userAgents []UserAgents
	if err := json.NewDecoder(r).Decode(&userAgents); err != nil {
		return nil, err
	}
	return &userAgents, nil
}
