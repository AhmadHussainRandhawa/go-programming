package main

import (
	"fmt"
	"os"
)

func main() {
	// plan: reading, writing, creating, open -Append - create- readwrite, stat, copy, rename, remove

	data := "This is the 1st line\nThis is the 2nd line\nThis is the 3rd line\n"

	writeFile("notes.txt", data)
	readFile("notes.txt")

	appendData("notes.txt", "append this in my data\n")
	readFile("notes.txt")

	infoFile("notes.txt")
	infoFile("rough.txt")

	// createFile("random.txt")
	// renameFile("random.txt", "rough.txt")
	// removeFile("rough.txt")

}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("\nRead content is in the following:\n", string(data))
}

func writeFile(fileName, data string) {
	err := os.WriteFile(fileName, []byte(data), 0o644)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func appendData(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(data)
}

func infoFile(fileName string) {
	info, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nName:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("Is Directory?:", info.IsDir())
	fmt.Println("Last updated:", info.ModTime())

}

func renameFile(oldFileName, newFileName string) {
	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func createFile(fileName string) {
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

}
