package package_manager

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Download(filename, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET Fail: %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP GET Fail: %d\n", resp.StatusCode)
		return
	}
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("File Creaft Fail: %v\n", err)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Printf("File Download Fail: %v\n", err)
		return
	}
	err = Unzip(filename)
	if err != nil {
		return
	}
}

func Unzip(filename string) error {
	destDir := path.Join(".module", strings.Replace(filename, ".zip", "", 1))
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Printf("Can't open zip file: %v\n", err)
		return err
	}
	defer zipReader.Close()
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Can't create dir: %v\n", err)
		return err
	}
	for _, file := range zipReader.File {
		parts := strings.Split(file.Name, "/")
		if len(parts) > 1 {
			file.Name = strings.Join(parts[1:], "/")
		}
		destPath := filepath.Join(destDir, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(destPath, file.Mode()); err != nil {
				fmt.Printf("Can't create dir: %v\n", err)
				return err
			}
			continue
		}
		if err := ExtractFile(file, destPath); err != nil {
			fmt.Printf("Can't extract file: %v\n", err)
			return err
		}
	}
	return nil
}

func ExtractFile(file *zip.File, destPath string) error {
	sourceFile, err := file.Open()
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	if err := os.Chmod(destPath, file.Mode()); err != nil {
		return err
	}
	return nil
}
