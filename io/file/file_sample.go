package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//create specific file by file name
func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	check(err)
}

// write date to file
func WriteData2File(fileName, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), os.ModePerm)
	check(err)
}

//Write data to file but not cover existing content
func AppendData2File(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()
	check(err)
	file.Write([]byte(data))
}

//open specific file by file name
func OpenFile(fileName string) {
	file, err := os.Open(fileName)
	check(err)
	//buffer for each read operation
	bufsize := 1024
	//offset
	offset := 0
	buf := make([]byte, bufsize)
	for {
		len, _ := file.ReadAt(buf, int64(offset))
		offset = offset + len
		if len == 0 {
			break
		}
		log.Println(string(buf))
	}
	defer file.Close()
}

//if whether current filePath is a file , then read this file directly.
//if whether current filePath is a directory, then will ready all the files under this directory
func ReadAllFile(filePath string) {
	if IsDir(filePath) {
		//read a file
		file, _ := os.Open(filePath)
		log.Println(file.Name())
		defer file.Close()
		return
	}

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		//read all files
		log.Println(f.Name())
	}
}

func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	check(err)
	log.Println(data)
}

//create one folder
func CreateOneFolder(dirPath string) {
	err := os.Mkdir(dirPath, os.ModePerm)
	check(err)
	log.Println("Create folder successfully!,folder path:" + dirPath)
	os.RemoveAll(dirPath)

}

//create multiple folders
func CreateMultiDir(dirsPath string) {
	if !IsExist(dirsPath) {
		err := os.MkdirAll(dirsPath, os.ModePerm)
		check(err)
		log.Println("Create folder successfully!,folder path:" + dirsPath)
		os.RemoveAll(strings.Split(dirsPath, "/")[0])

	}

}

func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	check(err)
	log.Println("Remove file successfully!,file name:" + fileName)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//check whether current file path is directory
func IsDir(dirPath string) bool {
	f, _ := os.Stat(dirPath)
	return f.IsDir()
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
