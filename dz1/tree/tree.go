package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// сюда писать функцию DirTree

func dirTree(out io.Writer, path string, printFiles bool) error {

	if err := InnerTree(out, path, printFiles, ""); err != nil {
		_ = fmt.Errorf("%s", err)
		return err
	}

	return nil
}

func InnerTree(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		_ = fmt.Errorf("can't read filenames")
	}

	nameToFile := make(map[string]os.FileInfo)
	fileNames := make([]string, 0)

	for _, file := range files {
		if file.IsDir() || printFiles {
			fileNames = append(fileNames, file.Name())
			nameToFile[file.Name()] = file
		}
	}

	sort.Strings(fileNames)

	sortedFiles := make([]os.FileInfo, 0)
	for _, name := range fileNames {
		sortedFiles = append(sortedFiles, nameToFile[name])
	}

	for idx, file := range sortedFiles {
		if file.IsDir() {
			newPrefix := ""
			if idx == len(sortedFiles) - 1 {
				_, _ = out.Write([]byte(prefix + "└───" + file.Name() + "\n"))
				newPrefix = prefix + "\t"
			} else {
				_, _ = out.Write([]byte(prefix + "├───" + file.Name() + "\n"))
				newPrefix = prefix + "│\t"
			}

			nextDir := filepath.Join(path, file.Name())
			_ = InnerTree(out, nextDir, printFiles, newPrefix)
		} else {
			if idx == len(sortedFiles) - 1 {
				formattedString := ""
				if file.Size() > 0 {
					formattedString = fmt.Sprintf("└───%s (%vb)\n", file.Name(), file.Size())
				} else {
					formattedString = fmt.Sprintf("└───%s (empty)\n", file.Name())
				}
				_, _ = out.Write([]byte(prefix + formattedString))
			} else {
				formattedString := ""
				if file.Size() > 0 {
					formattedString = fmt.Sprintf("├───%s (%vb)\n", file.Name(), file.Size())
				} else {
					formattedString = fmt.Sprintf("├───%s (empty)\n", file.Name())
				}
				_, _ = out.Write([]byte(prefix + formattedString))
			}
		}
	}

	return nil
}