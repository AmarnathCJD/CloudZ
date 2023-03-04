package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
)

func HandleCtrlZ() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Ctrl+C pressed in Terminal, shutting down!")
			os.Exit(0)
		}
	}()
}

func SendPortAndIP(port int) {
	ip := GetOutboundIP()
	fmt.Println("Server running on: " + ip + ":" + fmt.Sprint(port))
	fmt.Println("Press Ctrl+C to stop")
}

func GetOutboundIP() string {
	ipcmd := "hostname -I | awk '{print $1}'"
	if runtime.GOOS == "windows" {
		ipcmd = "ipconfig | findstr IPv4"
	}
	ip, err := RunCommand(ipcmd)
	if err != nil {
		log.Fatal("Error getting IP address: ", err)
		return ""
	}
	ipregex := "([0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+)"
	ip = RegexFirstMatch(ipregex, ip)
	return ip
}

func RegexFirstMatch(regex string, str string) string {
	re := regexp.MustCompile(regex)
	if re.MatchString(str) {
		return re.FindStringSubmatch(str)[1]
	}
	return ""
}

func RunCommand(cmd string) (string, error) {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = exec.Command("cmd", "/C", cmd).Output()

	} else {
		out, err = exec.Command("sh", "-c", cmd).Output()
	}
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func JoinPathWithDir(dir string, path string) string {
	if runtime.GOOS == "windows" {
		dir = filepath.FromSlash(dir)
		return filepath.Join(dir, path)
	}
	return filepath.Join(dir, path)
}
