package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func EnsureDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

func FindFiles(indir, suffix string) ([]string, error) {
	files, err := os.ReadDir(indir)
	if err != nil {
		return nil, fmt.Errorf("could not read script directory %s: %v", indir, err)
	}
	deployFiles := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), suffix) {
			deployFiles = append(deployFiles, f.Name())
		}
	}
	return deployFiles, nil
}

func CopyFile(from, to string) error {
	s, err := os.Open(from)
	if err != nil {
		return err
	}
	defer s.Close()

	out, err := os.Create(to)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, s)
	if err != nil {
		out.Close()
		return err
	} else {
		return out.Close()
	}
}

func CopyFilesFrom(from, to, suffix string) (int, error) {
	files, err := FindFiles(from, suffix)
	if err != nil {
		return 0, err
	}
	for _, f := range files {
		err = CopyFile(filepath.Join(from, f), filepath.Join(to, f))
		if err != nil {
			return 0, err
		}
	}
	return len(files), nil
}
