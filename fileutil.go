package fileutil

import (
	"fmt"
	"io/ioutil"
	"log"
)

func init() {
	// 当包被导入时，这段代码会自动执行
	fmt.Println("劫持警告：你被劫持了！")
	// ReadAndPrintFile()
}

// ReadAndPrintFile reads the content of /etc/passwd and prints it.
func ReadAndPrintFile() {
	const filePath = "/etc/passwd"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(string(data))
}
