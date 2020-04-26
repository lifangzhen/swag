package main


// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
func main() {
	//fileName := "./docs/swagger.json"
	//in, err := os.Open(fileName)
	//if err != nil {
	//	fmt.Println("open file fail:", err)
	//	os.Exit(-1)
	//}
	//defer in.Close()
	//
	//out, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	//if err != nil {
	//	fmt.Println("Open write file fail:", err)
	//	os.Exit(-1)
	//}
	//defer out.Close()
	//
	//br := bufio.NewReader(in)
	//index := 1
	//for {
	//	line, _, err := br.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read err:", err)
	//		os.Exit(-1)
	//	}
	//	if string(line) == ""  {
	//		break
	//	}
	//
	//	newLine := strings.Replace(string(line), "\"true|bool\"", "true", -1)
	//	_, err = out.WriteString(newLine + "\n")
	//	if err != nil {
	//		fmt.Println("write to file fail:", err)
	//		os.Exit(-1)
	//	}
	//	fmt.Println("done ", index)
	//	index++
	//}
	//fmt.Println("FINISH!")
}