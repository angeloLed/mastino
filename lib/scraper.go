package lib

import (
   // "os"
    "io/ioutil"
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "encoding/json"
    "strings"
    "log"
)

type Config struct {
    Url string `json:"url"`
    Tags []Tag
}

type Tag struct {
    Type string `json:"type"`
    Attributes []Attribute
    Count int
}

type Attribute struct {
    Name string `json:"name"`
    Value  string `json:"value"`
}

//---

type Sgraper struct {
    strem *http.Response
    Config Config
}

func (sg *Sgraper) Go(jsonConf string) string {

    sg.ResolveConfiguration(jsonConf)

    fmt.Println("opening : " + sg.Config.Url)

    resp := sg.getStream()
    z := html.NewTokenizer(resp.Body)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:

            sg.closeStrem()
            result, err := json.Marshal(sg.Config)
            if err != nil {
                log.Fatal(err)
            }
            return string(result)

        case tt == html.StartTagToken:
            sg.analyzeToken(z.Token())
        }
    }
}

func (sg *Sgraper) ResolveConfiguration(jsonConf string) {
    reader := strings.NewReader(jsonConf)
    err := json.NewDecoder(reader).Decode(&sg.Config)
    if err != nil {
        log.Fatal(err)
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
    }
    */

fmt.Println(t)
    // for every tags check the type
    for x,tag := range sg.Config.Tags {
        if (tag.Type == t.Data) {

            htmlAttrs := t.Attr
            structAttrs := tag.Attributes
            metch := true

            if(len(structAttrs) > 0) {
                for _,strAtt := range structAttrs {
                    foundAttr := false

                    for _,htmlAtt := range htmlAttrs {
                        if(strAtt.Name == htmlAtt.Key) {

                            // fmt.Println(strAtt.Name, strAtt.Value, htmlAtt.Key, htmlAtt.Val)
                            foundAttr = true
                            if(! (strAtt.Value == "" || strAtt.Value == htmlAtt.Val)) {
                                metch = false
                            }
                        }

                    }

                    if(! foundAttr) {
                        metch = false
                    }
                }
            }

            if(metch) {
                 sg.Config.Tags[x].Count ++
            }

        }
    }

    // fmt.Println(sg.Config.Tags[0].Type,sg.Config.Tags[0].Count)
    // fmt.Println(sg.Config.Tags[1].Type,sg.Config.Tags[1].Count)
}

// STREAM
func (sg *Sgraper) getStream() *http.Response {

    resp, _ := http.Get(sg.Config.Url)
    sg.strem = resp
    return resp
}
func (sg *Sgraper) closeStrem() {
    sg.strem.Body.Close()
}