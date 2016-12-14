package main

import (
    "mastino/lib"
    "fmt"
    "os"
)

func main() {

    // configString := `{"url":"https://blog.golang.org/error-handling-and-go", "tags": [
    //         {
    //             "type": "div",
    //             "attributes": [
    //             {
    //                 "key": "class",
    //                 "value": "code"
    //             }
    //             ]
    //         }
    //     ]}`

    configString := ""
    if(len(os.Args) > 1) {
        configString = os.Args[1]
    }


    sg := new(lib.Mastino)
    sg.Go(configString)

    result := sg.GetResults()

    fmt.Println(result)
}