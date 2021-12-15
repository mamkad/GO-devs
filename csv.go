package main

import "fmt"
import "archive/zip"
import "strings"
import "encoding/csv"

func main() {
		const zipFileName string = "./task.zip" 
		
		zipArchive, err := zip.OpenReader(zipFileName)
		
		if err != nil {
				fmt.Println(err)
				return
		}
		defer zipArchive.Close()
		
		err = ReadFilesFromZip(zipArchive)
		
		if err != nil {
				fmt.Println(err)
				return
		}
}

func IsFile(fileName string) bool{
		if i := strings.Index(fileName, ".txt"); i != -1 {
				return true
		}
		return false
}

func IsCorrectData(data [][]string) bool{
	return len(data) == 10 && len(data[0]) == 10
}


func ReadFilesFromZip(zipArchive *zip.ReadCloser) error {
		for _, file := range zipArchive.File {
				if IsFile(file.Name) {
						buff, err := file.Open()   
						if err != nil {        
								return err
						}
						defer buff.Close()  
						
						data, err := csv.NewReader(buff).ReadAll()
						if err != nil {        
								return err
						}
						
						if IsCorrectData(data) {
								fmt.Println("Файл " + file.Name + " " + data[4][2])
								break;
						}
				}
		}
		return nil	
}
