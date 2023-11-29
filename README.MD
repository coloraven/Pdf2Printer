
# PDFtoPrinter用法
下载地址：https://mendelson.org/PDFtoPrinter.exe，官方文档：https://mendelson.org/pdftoprinter.html
## 调用默认打印机
调用默认打印机打印`filename.pdf`文件:
```bash
PDFtoPrinter.exe filename.pdf
```
文件名支持绝对路径，但是如果路径中有空格，则需要用双引号包裹。
## 指定打印机打印
第二个非命名参数为指定打印机打印时的打印机名，需要双引号包裹:
```bash
PDFtoPrinter.exe filename.pdf "Name of Printer"
```
## 网络打印机
如果要打印到网络打印机，请使用Windows Print对话框中出现的名称，例如，请小心在名称开始时注意两个后斜线，以及Servername之后的单个后斜线）：
```bash
PDFtoPrinter.exe filename.pdf "\\SERVER\PrinterName"
```
如果您将打印机名称误认为或指定不存在的打印机，则不会打印，但是不会显示错误消息。
## 选择打印机
要从已安装的打印机列表中选择打印机，请将程序的名称更改为`PDFtoPrinterSelect.exe`。启动时，它将弹出已安装的打印机列表，除非您在命令行上指定打印机名称，否则在这种情况下，它将在未显示菜单的情况下打印到指定的打印机。

## 指定页面范围
无论您是否指定打印机，都可以通过添加可选的命名参数来指定要打印的页面范围（像这样没有空格）:
```bash
PDFtoPrinter.exe filename.pdf pages=#-#
```
单独的多个页面范围带有逗号（不能有空格），像这样: `2-4,7,12` 或者,打印指定页码及其之后的所有页: `7-` . 也可以使用如 `z-1` 一样来逆序指定页码打印最后一页。

## 指定打印份数 
要打印多个副本，请添加一个可选的参数，例如，在符号周围没有空格）：
```bash
PDFtoPrinter.exe filename.pdf copies=#
```
## 打印带密码的文件
要打印带密码的文件，请在命令行上加上`/p:[密码]`

## 静默打印
静静地运行，没有用户反馈或重点更改：在命令行上加上 `/s`。

## 打印多个PDF文件
在命令行上指定通配符（例如 `A*.pdf` ）;默认情况下，此命令递归搜索子目录。指定单个文件名时，您也可以使用 `/r`开关，或 `/Rn`来指定递归搜索子目录的深度。