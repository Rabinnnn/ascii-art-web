package asciiart

import (
	"fmt"
	"os"
)

// Checks the properties of banner file with focus on size and non-exist error
func CheckStatus(s string) {
	_, err := os.Stat(s)
	if os.IsNotExist(err) {
		var choice string
		fmt.Println("Banner file missing! Would you wish to download and continue? yes/no")
		fmt.Scanf("%s", &choice)
		switch choice {
		case "yes", "y", "YES", "Yes", "Y":
			err := Download(s)
			if err != nil {
				fmt.Printf("error donloading file: %s", err)
			}
		case "no", "No", "NO", "n", "N":
			fmt.Println("Can't download without user authorization. Bye.")
			os.Exit(1)
		default:
			fmt.Println("Wrong USER input. Try Again later.")
			os.Exit(1)
		}

	}
	switch s {
	case "ascii-art/banners/standard.txt":
		if info, _ := os.Stat("ascii-art/banners/standard.txt"); info.Size() != 6623 {
			fmt.Println("File content in 'standard.txt' changed. Please confirm!")
			os.Exit(1)
		}

	case "ascii-art/banners/shadow.txt":
		if info, _ := os.Stat("ascii-art/banners/shadow.txt"); info.Size() != 7463 {
			fmt.Println("File content in 'shadow.txt' changed. Please confirm!")
			os.Exit(1)
		}
	case "ascii-art/banners/thinkertoy.txt":
		if info, _ := os.Stat("ascii-art/banners/thinkertoy.txt"); info.Size() != 5558 {
			fmt.Println("File content in 'thinkertoy.txt' changed. Please confirm!")
			os.Exit(1)
		}

	}

}
