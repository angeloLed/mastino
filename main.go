package main

import (
    "SGraper/lib"
    "fmt"
)

func main() {

    sg := new(lib.Sgraper)
    //sg.SetUrl("https://google.com")
    sg.Go(`{"url":"https://google.com", "search": [
            {
                "tag": "a",
                "attributes": [{ "asd":"asd", "asd1":"asd1"}]
            },
            {
                "tag": "div",
                "attributes" : [{ "asd":"asd", "asd1":"asd1"}]
            }
        ]}`)

    fmt.Println("HTML:\n\n",)
}