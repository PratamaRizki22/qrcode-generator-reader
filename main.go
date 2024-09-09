package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
	goqr "github.com/tuotoo/qrcode"
)

const configFileName = "config.txt"

func readDefaultFolder() (string, error) {
	file, err := os.Open(configFileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", fmt.Errorf("failed to read default folder from config")
}

func writeDefaultFolder(folderPath string) error {
	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(folderPath)
	if err != nil {
		return err
	}
	return nil
}

func ensureFolderExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
	}
	return nil
}

func main() {
	defaultFolder, err := readDefaultFolder()
	if err != nil {
		fmt.Println("Enter default folder to save QR Codes:")
		fmt.Scanln(&defaultFolder)

		err = writeDefaultFolder(defaultFolder)
		if err != nil {
			log.Fatalf("Failed to save default folder: %v", err)
		}
		fmt.Println("Default folder saved successfully.")
	}

	var choice int
	fmt.Println("Choose an option:")
	fmt.Println("1. Create QR Code")
	fmt.Println("2. Read QR Code")
	fmt.Print("Enter choice (1/2): ")
	fmt.Scan(&choice)

	if choice == 1 {
		createQRCode(defaultFolder)
	} else if choice == 2 {
		readQRCode()
	} else {
		fmt.Println("Invalid choice.")
	}
}

func createQRCode(defaultFolder string) {
	var textInput string
	var fileName string
	var formatChoice int
	var format string

	fmt.Println("Enter text for the QR Code:")
	fmt.Scanln(&textInput)

	fmt.Println("Enter output file name (without extension):")
	fmt.Scanln(&fileName)

	fmt.Println("Choose file format:")
	fmt.Println("1. png")
	fmt.Println("2. jpg")
	fmt.Println("3. svg")
	fmt.Print("Enter format choice (1/2/3): ")
	fmt.Scan(&formatChoice)

	switch formatChoice {
	case 1:
		format = "png"
	case 2:
		format = "jpg"
	case 3:
		format = "svg"
	default:
		fmt.Println("Invalid format choice.")
		return
	}

	fileName = fileName + "." + format

	err := ensureFolderExists(defaultFolder)
	if err != nil {
		log.Fatalf("Failed to ensure folder exists: %v", err)
	}

	filePath := filepath.Join(defaultFolder, fileName)

	switch format {
	case "png", "jpg":
		err := qrcode.WriteFile(textInput, qrcode.Medium, 256, filePath)
		if err != nil {
			log.Fatal(err)
		}
	case "svg":
		svgStr, err := qrcode.New(textInput, qrcode.Medium)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(filePath, []byte(svgStr.ToString(false)), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("QR Code created and saved at: %s\n", filePath)

	qrASCII, err := qrcode.New(textInput, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nQR Code in ASCII format:")
	fmt.Println(qrASCII.ToSmallString(false))
}

func readQRCode() {
	var fileName string

	fmt.Println("Enter QR Code file path (e.g., /path/to/qrcode.png):")
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	qrMatrix, err := goqr.Decode(file)
	if err != nil {
		fmt.Println("Failed to read QR code:", err)
		return
	}

	fmt.Println("QR Code content:", qrMatrix.Content)

	qrASCII, err := qrcode.New(qrMatrix.Content, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qrASCII.ToSmallString(false))
}
