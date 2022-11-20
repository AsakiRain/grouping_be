package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		// mac
		cmd = "open"
	default:
		// linux
		cmd = "xdg-open"
	}
	args = append(args, url)

	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Printf("打开浏览器失败：%s\n请手动访问%s\n", err.Error(), url)
	} else {
		fmt.Printf("已经在浏览器中打开%s，请前往浏览器操作\n", url)
	}
}
