package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid number of args!")
		fmt.Fprintln(os.Stderr, "usage: <path to format>")
		os.Exit(1)
	}

	result := formatPath(os.Args[1])

	fmt.Print(result)
}

func formatPath(path string) string {
  // this is a special case, as we cannot really improve the formatting here
  if path == "/" {
    return "/"
  }

	pathToFormat := strings.TrimSuffix(path, string(filepath.Separator))

	var resultBuilder strings.Builder
	resultBuilder.Grow(len(pathToFormat))

	homePath, found := os.LookupEnv("HOME")

	if found {
		prefix, rest := splitPrefix(homePath, pathToFormat)

		if prefix != "" {
			resultBuilder.WriteString(prefix)

			// we need to prevent the prefix from being written twice
			if prefix != string(filepath.Separator) {
				resultBuilder.WriteRune(filepath.Separator)
			}
		}

		pathToFormat = rest
	}

	parts := strings.Split(pathToFormat, string(filepath.Separator))

	if len(parts) == 1 {
		resultBuilder.WriteString(parts[0])
	} else {
		for _, p := range parts[0 : len(parts)-1] {
			// an empty part always means a bug
			if len(p) == 0 {
				println("unexpected empty element in path:", path)
				os.Exit(1)
			}

			// ".." will stay the same, as it is kinda special
			if len(p) == 1 || p == ".." {
				resultBuilder.WriteString(p)
			} else {
				// we just want to add the first char from a file path part, to make it shorter
				c := []rune(p)[0]
				resultBuilder.WriteRune(c)
			}
			resultBuilder.WriteRune(filepath.Separator)
		}

		// the final path part of the input will be added as a whole
		resultBuilder.WriteString(parts[len(parts)-1])
	}

	return resultBuilder.String()
}

func splitPrefix(home, target string) (string, string) {
	if strings.HasPrefix(target, home) {
		path := strings.TrimPrefix(target, home)

		if filepath.IsAbs(path) {
			path = strings.TrimPrefix(path, string(filepath.Separator))
		}

		return "~", path
	}

	if filepath.IsAbs(target) {
		return string(filepath.Separator), strings.TrimPrefix(target, string(filepath.Separator))
	}

	return "", target
}
