package network

import (
	"fmt"
	"os"
	"syscall"
)

func SyscallNet() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println("Error in syscall.AF_INET, syscal.SOCK_RAW, syscall.IPROTO_ICMP")
		return
	}

	f := os.NewFile(uintptr(fd), "captureICMP")
	if f == nil {
		fmt.Println("Error in os.NewFile:", err)
		return
	}
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 256)
	if err != nil {
		fmt.Println("Error in syscal.Socket:", err)
		return
	}
	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("% X\n", buf[:numRead])
	}
}