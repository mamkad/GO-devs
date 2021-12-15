package main

import "bufio"
import "fmt"
import "os"
import "io"

func main() {
		var fileName string = "task.data"
		
		file, err := os.Open(fileName)
		if err != nil {
				fmt.Println(err)
				return
		}
		defer file.Close()
		
		reader := bufio.NewReader(file)
		
		var delim byte = ';'
		var line string
		var required_line string = "0;"
		
		line, err = reader.ReadString(delim);
		for i := 1; err != io.EOF; i++ {
				if line == required_line {
						fmt.Println("pos is ",i)
						break;
				}
				line, err = reader.ReadString(delim)
		}
}
