package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// PrintJSON 格式化输出 JSON 格式
func PrintJSON(any ...interface{}) {
	fmt.Println(SprintJSON(any))
}

// SprintJSON 格式化输出 JSON 格式并返回字符串
// NOTICE: 包含换行符
func SprintJSON(any ...interface{}) string {
	strSlice := make([]string, 0)
	for _, v := range any {
		b, _ := json.MarshalIndent(v, "", "  ")
		strSlice = append(strSlice, string(b))
	}
	return strings.Join(strSlice, "\n")
}

//创建日志文件并写入运行日志信息
func LogGenerator(dst string, SprintJSON string) {
	File, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0766)

	if err != nil {
		fmt.Print("open file err=", err, "\n")
		return
	}

	defer File.Close()

	Writer := bufio.NewWriter(File)
	Writer.WriteString(SprintJSON)

	Writer.Flush()
}
