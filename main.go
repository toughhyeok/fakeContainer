package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		os.Exit(1)
	}
}

func run() {
	fmt.Printf("Running: %v as %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	// new UTS/PID/Mount namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child: %v as %d\n", os.Args[2:], os.Getpid())

	// hostname 설정
	must(syscall.Sethostname([]byte("fake-container")))

	syscall.Chroot("/tmp/ubuntu")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	defer syscall.Unmount("proc", 0)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
