package util

import (
	"log"
	"os"
)

func MakeDir(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Fatalf("create dir %s is error: %s\n", dirName, err)
		}
	} else if err != nil {
		log.Fatalf("get dir is error: %s\n", err)
	}
}
