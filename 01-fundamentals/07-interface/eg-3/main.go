package main

import "fmt"

type LocalStorage struct{}
type CloudStorage struct{}

func (local LocalStorage) Upload(filename string) {
	fmt.Println(filename,"is uploaded to Local Storage successfully")
}

func (local LocalStorage) Delete(filename string) {
	fmt.Println(filename,"is deleted from Local Storage successfully")
}

func (cloud CloudStorage) Upload(filename string) {
	fmt.Println(filename,"is uploaded to Cloud Storage successfully")
}

func (cloud CloudStorage) Delete(filename string) {
	fmt.Println(filename,"is deleted from Cloud Storage successfully")
}

type Storage interface{
	Upload(filename string)
	Delete(filename string)
}

func UploadFile(storage Storage,filename string) {
	storage.Upload(filename)
}

func DeleteFile(storage Storage,filename string) {
	storage.Delete(filename)
}

func main() {

	filename := "file.txt"

	local := LocalStorage{}
	cloud := CloudStorage{}

	UploadFile(local,filename)
	UploadFile(cloud,filename)

	DeleteFile(local,filename)
	DeleteFile(cloud,filename)

}