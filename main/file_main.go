package main

import (
	"eshare-go-examples/io/file"
)

const TEST_FILE_NAME = "test.txt"

func main() {
	file.CreateFile(TEST_FILE_NAME)
	file.CreateFile(TEST_FILE_NAME)
	file.CreateMultiDir("test/demo")
	file.OpenFile(TEST_FILE_NAME)
	file.WriteData2File(TEST_FILE_NAME, "Hello,Evan")
	file.AppendData2File(TEST_FILE_NAME, "Hello World")
	file.ReadFile(TEST_FILE_NAME)
	file.ReadAllFile(".")
	file.DeleteFile(TEST_FILE_NAME)
}
