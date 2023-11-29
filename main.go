package main

import (
	"embed"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// 下面这句很关键，用于将PDFtoPrinter.exe程序编译到toolBinary变量中。
// embed更多用法见https://taoshu.in/go/how-to-use-go-embed.html.
//
//go:embed PDFtoPrinter.exe
var toolBinary embed.FS

func main() {
	// 从嵌入的文件系统中提取工具的二进制
	data, err := fs.ReadFile(toolBinary, "PDFtoPrinter.exe")
	if err != nil {
		panic(err)
	}
	// 创建临时文件
	tmpDir, err := os.MkdirTemp("", "tool")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir) // 确保在退出前清理临时目录

	pdftoprinter := filepath.Join(tmpDir, "pdftoprinter.exe")
	err = os.WriteFile(pdftoprinter, data, 0700)
	if err != nil {
		panic(err)
	}

	// 执行嵌入的工具
	cmd := exec.Command(pdftoprinter, `C:\*****.pdf`, "pages=8,4-5") // 替换为您的参数
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
