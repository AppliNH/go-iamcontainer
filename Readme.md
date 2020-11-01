# Go I am Container

Example of a container made from scratch with Go.

## Rootfs you can use

Have look [here](http://codelectron.com/how-to-get-a-linux-root-filesystem/) :)

## Run

After downloading the root file system of your choice, extract it to `./rootfs` and modify this line of code to target it :

```go
syscall.Chroot("./rootfs/alpine")
```

Then, you can run :

`sudo go run main.go run <command> <args> <opts>`

You can try

`sudo go run main.go run /bin/ls` 
for example.