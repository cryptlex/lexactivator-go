package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"log"
)

func main() {
	baseURL := "https://dl.cryptlex.com/downloads/"
      libVersion :=  "v3.30.3";
	basePath := "./libs/"
	fmt.Println("Downloading LexActivator libs " + libVersion + " ...")
	url := baseURL + libVersion + "/LexActivator-Static-Mac.zip"
	err := downloadFile(url, "libs/clang/universal/libLexActivator.a", basePath+"darwin_amd64/libLexActivator.a")
	if err != nil {
		panic(err)
	}
	url = baseURL + libVersion + "/LexActivator-Win.zip"
	err = downloadFile(url, "libs/vc14/x64/LexActivator.dll", basePath+"windows_amd64/LexActivator.dll")
	if err != nil {
		panic(err)
	}
	url = baseURL + libVersion + "/LexActivator-Static-Linux.zip"
	err = downloadFile(url, "libs/gcc/amd64/libLexActivator.a", basePath+"linux_amd64/libLexActivator.a")
	if err != nil {
		panic(err)
	}
	err = downloadFile(url, "libs/gcc/arm64/libLexActivator.a", basePath+"linux_arm64/libLexActivator.a")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Lexactivator libs downloaded successfully!")
}

func downloadFile(url string, packagePath string, targetpath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = unzip(body, packagePath, targetpath)

	if err != nil {
		return err
	}
	return nil
}

func unzip(body []byte, packagePath string, targetpath string) error {
	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		if file.Name == packagePath {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()
			targetFile, err := os.Create(targetpath) // OpenFile(targetpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		}		
	}

	return nil
}
