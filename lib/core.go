package lib

import (
    "io/ioutil"
    //"fmt"
    "net/http"
    "golang.org/x/net/html"
    "encoding/json"
    "strings"
    "log"
    "bytes"
)


type Config struct {
    Url string `json:"url"`
    Tags []Tag `json:"tags"`
    Error bool `json:"error"`
    Message string `json:"message"`
}

type Tag struct {
    Type string `json:"type"`
    Attributes []Attribute `json:"attributes"`
    Matches []string `json:"matches"`
}

type Attribute struct {
    Key string `json:"key"`
    Value  string `json:"value"`
}

type Mastino struct {
    strem *http.Response
    Config Config
}


// ---------------------------------

func (sg *Mastino) endJob() {
    if r := recover(); r != nil && ! sg.Config.Error {
        
        sg.Config.Error = true
        switch x := r.(type) {
        case string:
            sg.Config.Message = x
        case error:
            sg.Config.Message = r.(error).Error()
        default:
            sg.Config.Message = "unknown error"
        }
    }
}

func (sg *Mastino) Go(jsonConf string) {

    defer sg.endJob()

    if( ! sg.ResolveConfiguration(jsonConf)) {
        sg.Config.Error = true
        sg.Config.Message = "error on load configuration manifest"
    }

    resp := sg.getStream()
    z := html.NewTokenizer(resp.Body)

    for {
        tt := z.Next()

        switch {
            case tt == html.ErrorToken:
                sg.closeStrem()
                return

            case tt == html.StartTagToken:
                sg.analyzeToken(z.Token())
        }
    }
}

func (sg *Mastino) ResolveConfiguration(jsonConf string) bool {
    reader := strings.NewReader(jsonConf)
    err := json.NewDecoder(reader).Decode(&sg.Config)
    if err != nil {
        return false
    }
    return true
}

func (sg *Mastino) getBody() string {
    bytes, _ := ioutil.ReadAll(sg.strem.Body)
    return string(bytes)
}

func (sg *Mastino) analyzeToken(t html.Token) {

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
                        if(strAtt.Key == htmlAtt.Key) {
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
                
                //fmt.Println(t.String())
                sg.Config.Tags[x].Matches = append(sg.Config.Tags[x].Matches, t.String())
            }

        }
    }
}

func (sg *Mastino) GetResults() string {

    result, err := json.Marshal(sg.Config)
    if err != nil {
        log.Fatal(err)
    }
    return string(unscapeMarshal(result))
}
func unscapeMarshal(b []byte) []byte {

    b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
    b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
    b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
    
    return b
}

// STREAM
func (sg *Mastino) getStream() *http.Response {

    resp, _ := http.Get(sg.Config.Url)
    sg.strem = resp
    return resp
}
func (sg *Mastino) closeStrem() {
    sg.strem.Body.Close()
}