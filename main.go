package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gcodetec/gogcsclient/gcs"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("stating...")
	walkDir(viper.GetString("root_dir"))
}

func uploadFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if !info.IsDir() {
		dest := destPath(path)
		dest = strings.Replace(dest, "\\", "/", -1)
		uploadToGCS(path, dest)
	}
	return nil
}

func walkDir(dir string) {
	err := filepath.Walk(dir, uploadFile)
	if err != nil {
		log.Fatal(err)
	}
}

func uploadToGCS(origin string, dest string) {
	msg, err := gcs.Upload(origin, dest, viper.GetString("bucket"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func destPath(path string) (dest string) {
	baseDest := viper.GetString("base_dest")
	return strings.Replace(path, viper.GetString("root_dir"), baseDest, -1)
}
