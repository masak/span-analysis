package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    buf := bufio.NewReader(os.Stdin)
    for {
        s, err := buf.ReadString('\n')
        if err != nil {
            return
        }
        for i := 0; i < 16; i++ { // XXX: should be 30 in the long run
            new_board := s[ 0 : i ] + "w" + s[ i + 1 : len(s) ]
            fmt.Printf("%s", new_board);
        }
    }
}
