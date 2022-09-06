package api

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetImage_input(url string) {

	checkFile(url)

	fmt.Println("Scanning image...")

	err := exec.Command("trivy", "image", "-f", "json", "-o", "result/"+url+".json", url).Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finish scaning")

}

func GetImage_tar(name string) {

	checkFile(name)

	fmt.Println("Scanning image...")

	path := "tar_files/" + name

	err := exec.Command("trivy", "image", "-f", "json", "-o", "result/"+name+".json", "--input", path).Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finish scaning")

}

func checkFile(name string) {

	checkFolder()

	fmt.Println("\nCheck file exist...")
	if _, err := os.Stat("result/" + name + ".json"); err == nil {
		// remove file
		fmt.Println("File Exist! Removing...")
		exec.Command("rm", "result/", name).Run()
		fmt.Println("File removed")
		fmt.Printf("\n")
	} else {
		fmt.Println("File is not found.")
		fmt.Println("Default name is ", name+".json \n")
	}

}

func checkFolder() {

	fmt.Println("Check folder result...")
	if _, err := os.Stat("result/"); err != nil {
		fmt.Println("Creating folder result")

		exec.Command("mkdir", "result").Run()

		fmt.Println("Folder result has been created")
	} else {
		fmt.Println("Folder result is exist")
	}
}
