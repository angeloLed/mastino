**Mastino**
=======

A Golang library for scraping web page. It can be used with any language ( php, java, c# ... ) since **Mastino** uses json as input/output interchange format.

Usage
-----
run commands
```bash
git clone https://github.com/angeloLed/mastino.git mastino
cd mastino
go get -u
```

You can also run Mastino like this:
```bash
#compiled
./mastino "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}"

# or not compiled
go run main.go "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}"
```

**Manifest configuration**
Example of manifest configuration:
```json
{
	"url": "https://google.com",
	"tags": [{
			"type": "a",
			"attributes": [{
				"key": "class",
				"value": "bg1"
			}, {
				"type": "span"
			}]
		},
		{
			"type": "div",
			"attributes": [{
			    "key": "css"
	        }]
    }]
}
```

| key | type | description |
|---|---|---|
| url |string| target url |
|tags |array|each tag that Mastino must try to `bite`. The key contains all the matching rules. Contains `attributes` `type`|
|type|string|DOM element |
|attributes| array| each attributes of tag to verify. Contais `key` `value`|
|key| string|the key of attribute |
|value|string|the value of attribute|

**Possibile matching result**
 - `tag` is specified without `attribute`
 - `attribute` `key` is specified without `value`
 - `attribute` `key` is specified and `value` match with value of DOM

**Json Output**
Mastino returns a JSON with the same structure of the provided input JSON. Each  match found is returned us a `matches` entry inside `tags` element. `matches` is an array of strings ( the DOM ).
```json
{
	"url": "https://blog.golang.org/error-handling-and-go",
	"tags": [{
		"type": "div",
		"attributes": [{
			"key": "class",
			"value": "code"
		}],
		"matches": ["<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">", "<div class=\"code\">"]
	}],
	"error": false,
	"message": ""
}
```

**Using Mastino with Other languages**
PHP:
```php
$jsonManifest = "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}";
$output = shell_exec("/path/to/mastino {$jsonManifest}");
var_dump(json_decode($output, true));
```

C# :
```cs
using System;
using System.Diagnostics;

class Runshell
{
  static void Main()
  {
    ProcessStartInfo psi = new ProcessStartInfo();
    psi.FileName = "/path/to/mastino";
    psi.UseShellExecute = false;
    psi.RedirectStandardOutput = true;

    psi.Arguments = "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}";
    Process p = Process.Start(psi);
    string strOutput = p.StandardOutput.ReadToEnd();
    p.WaitForExit();
    Console.WriteLine(strOutput);
  }
}
```

Nodejs:
```js
var sys = require('sys')
var exec = require('child_process').exec;
function puts(error, stdout, stderr) { sys.puts(stdout) }
var manifest = "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}";
exec("/path/to/mastino " + manifest, puts);
```

Golang:
```go
package main

import "os/exec"

func main() {
    app := "/path/to/mastino"

    manifest := "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}"

    cmd := exec.Command(app, manifest, nil, nil, nil)
    stdout, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}
```


License
-------
MIT
