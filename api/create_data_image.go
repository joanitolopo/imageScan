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
	fmt.Println("Check folder result")
	if _, err := os.Stat("result/"); err != nil {
		fmt.Println("Creating folder result")

		exec.Command("mkdir", "result").Run()

		fmt.Println("Folder result has been created")
	} else {
		fmt.Println("Folder result is exist")
	}

	fmt.Println("Check file exist")

	if _, err := os.Stat("result/" + name + ".json"); err == nil {
		value = true
	} else {
		value = false
		fmt.Println("File is not found")
		fmt.Println("Default name is ", name+".json")
	}

	return value

}
