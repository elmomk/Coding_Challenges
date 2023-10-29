https://dev.to/napicella/linux-pipes-in-golang-2e8j


```go
func runCommand() error {
    if isInputFromPipe() {
        // if input is from a pipe, upper case the
        // content of stdin
        print("data is from pipe")
        return toUppercase(os.Stdin, os.Stdout)
    } else {
        // ...otherwise get the file
        file, e := getFile()
        if e != nil {
            return e
        }
        defer file.Close()
        return toUppercase(file, os.Stdout)
    }
}

func isInputFromPipe() bool {
    fileInfo, _ := os.Stdin.Stat()
    return fileInfo.Mode() & os.ModeCharDevice == 0
}
```


