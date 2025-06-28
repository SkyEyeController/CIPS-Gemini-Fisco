package internal

import (
	"bytes"
	"net/http"

	clog "github.com/kpango/glg"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

/******************************** Http Server ********************************/

// HttpEngine 实现了 ServeHTTP 方法，因此可以作为 http.Handler 使用
type HttpEngine struct {
	// router 记录了每种请求方式对应的所有路由规则
	router map[string]HandlerFunc
}

func NewHttpEngine() *HttpEngine {
	return &HttpEngine{router: make(map[string]HandlerFunc)}
}

// addRoute 用于注册路由规则
// 仅支持 GET 和 POST 方法
// 请求方式的记录形式为 method-pattern，例如 GET-/、POST-/hello
// 每种请求方式对应一个 HandlerFunc 函数
func (engine *HttpEngine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router[method+"-"+pattern] = handler
}

// 对 pattern 注册 GET 方法
func (engine *HttpEngine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 对 pattern 注册 POST 方法
func (engine *HttpEngine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *HttpEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND: " + req.URL.Path + "\n"))
	}
}

// 启动 HTTP 服务
func (engine *HttpEngine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

/******************************** Http Client ********************************/

// http 客户端
type HttpClient struct {
	server_url string
}

func NewHttpClient(server_url string) *HttpClient {
	return &HttpClient{server_url: server_url}
}

// http 客户端发送 GET 请求
func (client *HttpClient) GET(path string) []byte {
	req, err := http.NewRequest("GET", client.server_url+path, nil)
	if err != nil {
		clog.Fatalf("make GET request error: %v", err)
	}
	defer req.Body.Close()

	_ = req
	return nil
}

// http 客户端发送 POST 请求并获取响应
// resp 在使用后应该被关闭
func (client *HttpClient) POST(path string, data []byte) *http.Response {
	// httpclient := &http.Client{}
	// req, err := http.NewRequest("POST", client.server_url+path, bytes.NewReader(data))
	// // req, err := http.NewRequest("POST", client.server_url+path, bytes.NewReader([]byte("Hello world!")))
	// if err != nil {
	// 	clog.Fatalf("make POST request error: %v", err)
	// }
	// // 参见：https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi/23963271#23963271
	// req.Close = true

	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Connection", "keep-alive")
	// resp, err := httpclient.Do(req)
	// if err != nil {
	// 	clog.Fatalf("use POST to post data error: %v", err)
	// }

	resp, err := http.Post(client.server_url+path, "application/json", bytes.NewBuffer(data))
	if err != nil {
		clog.Fatalf("use POST to post data error: %v", err)
	}
	defer resp.Body.Close()

	return resp
}

// 测试使用
func (client *HttpClient) GetURL() string {
	return client.server_url
}
