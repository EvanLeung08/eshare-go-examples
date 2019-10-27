package file

import "testing"

const TEST_FILE_NAME = "test.txt"

func TestCreateFile(t *testing.T) {
	CreateFile(TEST_FILE_NAME)
}

func TestCreateMultiDir(t *testing.T) {
	CreateMultiDir("test/demo")
}

func TestOpenFile(t *testing.T) {
	OpenFile(TEST_FILE_NAME)
}

func TestWriteData2File(t *testing.T) {
	WriteData2File(TEST_FILE_NAME, "Hello,Evan")
}

func TestAppendData2File(t *testing.T) {
	AppendData2File(TEST_FILE_NAME, "Hello World")
}

func TestReadFile(t *testing.T) {
	ReadFile(TEST_FILE_NAME)
}

func TestReadAllFile(t *testing.T) {
	ReadAllFile(".")
}

func TestDeleteFile(t *testing.T) {
	DeleteFile(TEST_FILE_NAME)
}
