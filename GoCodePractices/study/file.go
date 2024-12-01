package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(filename string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败:", err)
		return
	}
	fmt.Println("当前工作目录:", currentDir)
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	content := make([]byte, 100)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("读文件发生错误", err)
		return
	}
	fmt.Println(string(content[:n])) //转字符串
	fmt.Println(content[:n])
	//fmt.Println(content)
}

func writeFile() {
	file, err := os.OpenFile("b.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	content := "我爱编程\n"
	n, err := file.Write([]byte(content))
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
	fmt.Printf("写入的字节数:%d", n)
}

func bufferRead() {
	file, err := os.Open("b.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF { // End Of File
				if line == "" {
					line = fmt.Sprintf(
						"长度为%d的字符串%s",
						len(line),
						line,
					)
				}
				fmt.Printf("文件末尾:%s", line)
				break
			} else {
				fmt.Println("读文件发生错误", err)
				return
			}
		} else {
			fmt.Print(line) // 本身line 中存在换行符 所以可以使用Print，而不用Println
		}
	}
}
func main() {
	//readFile("a.txt")
	//writeFile()
	bufferRead()
}
