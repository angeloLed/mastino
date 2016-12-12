package main

import (
    "SGraper/lib"
    "fmt"
)

func main() {

    sg := new(lib.Sgraper)
    //sg.SetUrl("https://google.com")
    result := sg.Go(`{"url":"http://www.tvapuntate.it/category/serie-tv-americane/", "tags": [
            {
                "type": "a",
                "attributes": [
                {
                    "name": "title",
                    "value": "Contatti"
                }
                ]
            }
        ]}`)

    fmt.Println("HTML:\n\n",result)
}