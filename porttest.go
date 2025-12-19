package main

import (
	"fmt"
	"net"
	"strconv" // 补充缺失的包：用于int转string
)

// 注意：函数参数名改为aport（和你保持一致），内部逻辑统一用aport
func CheckPortIsUsed(aport int, network string) (bool, error) {
	// 1. 修复：把port改为aport（参数名是aport）
	if aport < 0 || aport > 65535 {
		return false, fmt.Errorf("端口号%d非法，必须在0-65535之间", aport)
	}

	// 2. 拼接监听地址（":8090"格式）
	addr := ":" + strconv.Itoa(aport)

	// 3. 补充核心逻辑：尝试监听端口（声明err和listener）
	listener, err := net.Listen(network, addr)
	if err != nil {
		// 监听失败 = 端口被占用/权限不足
		return true, nil
	}

	// 4. 监听成功：关闭listener释放端口（避免占用）
	defer listener.Close()
	return false, nil
}

func main() {
	// 变量名：aport → ports（语义更清晰，可选）
	ports := []int{8080, 8091, 55555}
	var used bool
	var err error

	// 遍历每个端口（注意：循环变量用port，避免和切片名冲突）
	for _, port := range ports {
		// 调用函数，传循环变量port
		used, err = CheckPortIsUsed(port, "tcp")
		if err != nil {
			fmt.Printf("检查端口%d时出错，错误信息：%v\n", port, err)
			continue
		}
		if used {
			fmt.Printf("端口%d已被占用\n", port)
		} else {
			fmt.Printf("端口%d未被占用\n", port)
		}
	}
}
