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
		panic("Woops idk what to do with that")
	}

}

func run() {

	// Running /proc/self/exec is basically running a /fork/exec
	// So we will run a /fork/exec in a child process, to get our container acting as a process, and isolated
	// In short, this will call again this program, put in another process
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// CLONE_NEWUTS : Clones hostname from host, and thus allows to have a special space to store it inside the container
	// CLONE_NEWPID : Clones PID
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {

	fmt.Printf("I am a container, and I execute %v \n", os.Args[2])
	fmt.Printf("PID : %v \n", os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Setting sub-process filesystem
	// This allows to have a truly isolated container, with its own filesystem
	// + keep in mind that ps looks at /proc. This way, the container while have its own /proc and thus its own ps :)
	syscall.Chroot("./rootfs/alpine") // <= Change this to point to the root filesystem you want in your container
	os.Chdir("/")

	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
