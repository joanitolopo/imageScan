package api

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetImage_input(url string) {

	if file_name := checkFile(url); file_name {
		// remove file
		fmt.Println("File Exist and Need to remove...")
		exec.Command("rm", "result/", url).Run()
	}

	fmt.Println("Scanning image")

	err := exec.Command("trivy", "image", "-f", "json", "-o", "result/"+url+".json", url).Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finish scaning")

}

func GetImage_tar(path string) {
	fmt.Printf("Ini adalah functions get image by tar: %s", path)
}

func checkFile(name string) (value bool) {

	if _, err := os.Stat("result/" + name + ".json"); err == nil {
		value = true
	} else {
		value = false
	}

	return value

}
