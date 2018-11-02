
package main 


import(
	"fmt"
	"time"
	"groot/imooc/11-接口/mock"
	"groot/imooc/11-接口/real"
)

// duck typing
// 像鸭子走路，像鸭子叫（长得像鸭子），那么就是鸭子
// 描述事物的外部行为而非内部结构
// 严格说go属于结构化类型系统，类似duck typing
func main(){

	// 使用者 download   实现者retriever

	var r Retriever

	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	print(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	print(r)
	// fmt.Println(download(r))

	// fmt.Println(session(r)) //real.Retriever并没有实现Post
	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}

const url = "http://www.imooc.com"


type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, 
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}

// 组合接口 
type RetrieverPoster interface {
	Retriever
	Poster 
}

func session(s RetrieverPoster) string{

	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}


// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口指针
// 指针接受者 
func print(r Retriever) {
	fmt.Println("检查",r)
	fmt.Printf(">type:%T value:%v\n", r, r)
	switch v := r.(type){
		case *mock.Retriever:
			fmt.Println("contents:",v.Contents)
		case *real.Retriever:
			fmt.Println("useragent:",v.UserAgent)
	}	
}




















