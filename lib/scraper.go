package lib

import (
   // "os"
    "io/ioutil"
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "encoding/json"
)

type Sgraper struct {
    strem *http.Response
    config map[string]interface{}
}

func (sg *Sgraper) Go(jsonConf string) {

    sg.ResolveConfiguration(jsonConf)

    fmt.Println("opening : " + sg.config["url"].(string))

    resp := sg.getStrem()
    z := html.NewTokenizer(resp.Body)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            return
        case tt == html.StartTagToken:
            sg.analyzeToken(z.Token())
        }
    }

    sg.closeStrem()
    return
}

func (sg *Sgraper) ResolveConfiguration(jsonConf string) {
    b := []byte(jsonConf)
    var f interface{}
    err := json.Unmarshal(b, &f)
    sg.config = f.(map[string]interface{})

    if(err != nil) {
        panic(err)
    }
}

func (sg *Sgraper) getBody() string {
    bytes, _ := ioutil.ReadAll(sg.strem.Body)
    return string(bytes)
}

func (sg *Sgraper) analyzeToken(t html.Token) {

/*
            isAnchor := t.Data == "a"
            if isAnchor {
                fmt.Println("We found a link!")
            }*/

    fmt.Println(t.Data)
}


// STREAM
func (sg *Sgraper) getStrem() *http.Response {

    resp, _ := http.Get(sg.config["url"].(string))
    sg.strem = resp
    return resp
}
func (sg *Sgraper) closeStrem() {
    sg.strem.Body.Close()
}