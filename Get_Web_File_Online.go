package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {

	Init()

	Benner()

	fileUrl := "http://10.1.171.100/Data/7z1900-x64.exe" // Download File Url
	fileName := "7z1900-x64.exe"                         // Download File Name
	localFileName := fileName
	oldFileHash := "d7b20f933be6cdae41efbe75548eba5f" // 7zip 7z1900-x64.exe Hash

	for range time.Tick(time.Minute * 1) { // Timing For loop Format:time.'time unit' * Quantity
		err := DownloadFile(fileName, fileUrl)
		if err != nil {
			panic(err)
		}

		fmt.Println("Downloaded：" + fileUrl) // File Download Source

		filehash, err := hash_file_md5(localFileName) // Download File Hash
		if err == nil {
			fmt.Println("FileHash：" + filehash)
		}

		fmt.Println(time.Now().Format("FileTime：2006,01,02,15:04:05")) // Now File Download Time

		if filehash != oldFileHash {
			cmd := "start " + localFileName //Start Download File
			out, err := exec.Command("powershell.exe", cmd).Output()
			if err != nil {
				log.Fatalf("cmd.Start() failed with %s\n", err)
			}

			fmt.Printf(string(out))
		}
	}
}

func Benner() {
	fmt.Println("         #        ##   ")
	fmt.Println("        ##       #  #  ")
	fmt.Println(" #  #    #          #  ")
	fmt.Println(" #  #    #        ##   ")
	fmt.Println("  ##     #   ##  #     ")
	fmt.Println("  ##    ###  ##  ####  ")

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// Source:https://blog.csdn.net/wade3015/article/details/84878511
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Source:https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
func hash_file_md5(filePath string) (string, error) {
	// Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	// Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	// Tell the program to call the following function when the current function returns
	defer file.Close()

	// Open a new hash interface to write to
	hash := md5.New()

	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	// Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

func Init() {
	out, err := exec.Command("powershell.exe", "chcp 65001").Output()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)

	}
	if false {
		fmt.Printf(string(out))
	}

	out, err = exec.Command("powershell.exe", "ping -n 1 10.0.171.200").Output()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)

	}
	if false {
		fmt.Printf(string(out))
	}

	out, err = exec.Command("powershell.exe", "ping -n 1 10.0.171.201").Output()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)

	}
	if false {
		fmt.Printf(string(out))
	}

	out, err = exec.Command("powershell.exe", "ping -n 1 10.0.171.254").Output()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)

	}
	if false {
		fmt.Printf(string(out))
	}

	out, err = exec.Command("powershell.exe", "arp -a").Output()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)

	}

	if false {
		fmt.Println(string(out))
	}
}
