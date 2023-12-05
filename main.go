package main

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
)

type Result struct {
	Name string
	Age  int
	Url  string
}

func main() {
	client := req.C().
		EnableDumpAllWithoutResponse().
		SetOutputDirectory("/path/to/download") //下载保存路径
	login(client)
	// json等其他请求和逻辑
	var result Result
	client.R().SetSuccessResult(&result).Get("http://www.baidu.com:8070/base/admin/user/")

	callback := func(info req.DownloadInfo) {
		if info.Response.Response != nil {
			fmt.Printf("downloaded %.2f%%\n", float64(info.DownloadedSize)/float64(info.Response.ContentLength)*100.0)
		}
	}
	pdfulr := ""
	// 临时指定绝对下载路径.
	client.R().SetOutputFile("/tmp/test.pdf").
		SetDownloadCallback(callback). //下载进度条
		Get(pdfulr)
	//打印PDF
	PrintPDF("", "8")
	os.Remove("/tmp/test.pdf")
}
