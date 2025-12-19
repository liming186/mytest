package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if r.Method != http.MethodGet {
		http.Error(w, "只支持get请求", http.StatusMethodNotAllowed)
	}
}
func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}
func Divide(a int, b int) int {
	if b == 0 {
		return 0 // simple handling for division by zero
	}
	return a / b
}
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong（8090端口正常响应）"))
}
func main() {
	fmt.Println("Add(2,3) =", Add(2, 3))
	fmt.Println("Subtract(5,2) =", Subtract(5, 2))
	fmt.Println("Multiply(3,4) =", Multiply(3, 4))
	fmt.Println("Divide(10,2) =", Divide(10, 2))
	Nh := "计算结果："
	http.HandleFunc("/hello", HelloHandler) // /hello 路径走 HelloHandler
	http.HandleFunc("/ping", PingHandler)   // /ping 路径走 PingHandler
	fmt.Println(Nh, Add(7, 7), Divide(4, 2))
	port := ":809000"
	fmt.Printf("服务已启动（Mac）：http://localhost%s/hello\n", port)
	fmt.Printf("测试端口连通性：http://localhost%s/ping\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("启动服务失败，错误信息：", err)
	}

}
