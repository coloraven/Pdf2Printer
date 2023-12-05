package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// embed更多用法见https://taoshu.in/go/how-to-use-go-embed.html.
// 下面这句很关键，用于将PDFtoPrinter.exe程序编译到toolBinary变量中。
//
//go:embed PDFtoPrinter.exe
var toolBinary embed.FS

// pageFlags:"8,4-6,11-"
// 打印第8页以及第4到6页，还有11页以及之后所有页
func PrintPDF(pdf_file_path, pageFlags string) {
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
	cmd := exec.Command(pdftoprinter, pdf_file_path, "pages="+pageFlags) // 替换为您的参数
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
