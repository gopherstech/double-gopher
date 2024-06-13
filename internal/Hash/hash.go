package hash

import (
	"crypto/md5"
	file "double-gopher/internal/File"
	storage "double-gopher/internal/Storage"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// Takes the file system interface as argument and pointer to a storage to keep hashes
func HashFiles(folder fs.FS, storage *storage.Storage) {

	// fs.Walkdir from standard Go lang library fs.
	result := fs.WalkDir(folder, ".", func(path string, d fs.DirEntry, err error) error {
		//Check if the file is Directory, if it is we are not going to hash it.
		if d.IsDir() {
			fmt.Printf("%s, is directory so we can not hash it \n", path)
			//If the file is not Directory, we can hash it.
		} else if !d.IsDir() {
			//Read the file
			data, err := os.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			//Save path and hash
			resultFile := file.File{
				Name: path,
				Hash: md5.Sum(data),
			}
			//Upload and give us a notice if it is save or deleted
			flag := storage.UploadFileData(resultFile)
			if !flag {
				//If path and hash is not uploaded to the storage we get the notification
				fmt.Printf("File %s was not uploaded as it is a duplicate and has to be deleted \n", path)
				//Remove the file
				os.Remove(path)

			} else {
				//Notify if the file is not a duplicate and save path and hash
				fmt.Printf("%s was successfully stored \n", path)
			}
		}
		return nil
	})
	//In case smth went wrong, we better know it.
	if result != nil {
		log.Fatal("There were an error")
	}

}
