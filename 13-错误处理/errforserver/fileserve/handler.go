

package fileserve

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}
func HandFileList(
	writer http.ResponseWriter,request *http.Request) error{

	// 判断路径是不是/list/ 如/abc
	if strings.Index(request.URL.Path, prefix) != 0 {
		// path /abc must start with /list/
		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, prefix))
	}



	// 拿到具体目录
	path := request.URL.Path[len("/list/"):]//list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		// panic(err)

		// http.Error(writer,
		// 	err.Error(),
		// 	// http.StatusText(http.StatusInternalServerError),
		// 	http.StatusInternalServerError)
		// return
		return err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		// panic(err)
		return err
	}

	writer.Write(all)
	return nil
}









