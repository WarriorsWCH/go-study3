
package main

import (
	"net/http"
	"os"
	// "io/ioutil"
	"log"
	"groot/imooc/13-错误处理/errforserver/fileserve"
)


type appHandler func(writer http.ResponseWriter,
request *http.Request) error

// 统一处理错误
func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request){

	return func(writer http.ResponseWriter, request *http.Request){
		
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// panic：不要经常使用
		// 停止当前函数执行
		// 一直向上返回，执行每一层的defer
		// 如果没有遇见recover，程序退出
		// recover:
		// 仅在defer调用中使用
		// 获取panic的值
		// 如果无法处理，可重新panic
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())

			// user error 如path /abc must start with /list/
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
				// os提供的函数  IsNotExist文件不存在
				case os.IsNotExist(err):

					// 报403错误StatusNotFound
					// 第一个参数 向谁回报错误
					// 第二个参数 是一个字符串，不能写err.Error()，否则错误就暴露出去了
					// 第三个参数 状态码
					// http.Error(
					// 	writer,
					// 	http.StatusText(http.StatusNotFound),
					// 	http.StatusNotFound)

					code = http.StatusNotFound
					// 没权限
				case os.IsPermission(err):
					code = http.StatusForbidden//403
				default:
					code = http.StatusInternalServerError//500未知错误
			}
			http.Error(
				writer,
				http.StatusText(code),
				code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(fileserve.HandFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}