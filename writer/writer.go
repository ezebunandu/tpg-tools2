package writer

import (
	"flag"
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
    err := os.WriteFile(path, data, 0o600)
    if err != nil {
        return  err
    }
    return  os.Chmod(path, 0o600)
}

const Usage = `Usage: writefile -size <FILE SIZE> <FILENAME>`

func Main(){
    if len(os.Args) < 2 {
        fmt.Println(Usage)
        return
    }

    size := flag.Int("size", 0, "size in bytes")
    flag.Parse()

    if len(flag.Args()) < 1 {
        fmt.Fprintln(os.Stderr, Usage)
        os.Exit(1)
    }
    
    err := WriteToFile(flag.Args()[0], make([]byte, *size))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}