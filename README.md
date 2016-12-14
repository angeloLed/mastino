**Mastino**
=======

A library writed in Golang for scraping web page. It's can be use with any language ( php, java, c# ... ) because **Mastino** use json as input/output comunication between the `manifest configuration`.

Usage
-----
run commands
```bash
git clone https://github.com/angeloLed/mastino.git mastino
cd mastino
go get -u
```

Mastino can be run like:
```bash
#compiled
./mastino "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}"

# or not compiled
go run main.go "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}"
```

**The manifest configuration**
Example of manifest configuration
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
				}, {
					"type": "div",
					"attributes": [{
						"key": "css"
				}]
	}]
}
```

| key | type | description |
|---|---|---|
| url |string| url target |
|tags |array|each tag that Mastino must try to `bite`. This key contain all the roule of matching. Contains `attributes` `type`|
|type|string|DOM element |
|attributes| array| each attributes of tag to verify. Contais `key` `value`|
|key| string|the key of attribute |
|value|string|the value of attribute|

**Possibile Matching**

 - `tag` is specified without `attribute`
 - `attribute` `key` is specified without `value`
 - `attribute` `key` is specified and `value` metch with value of DOM

**Json Output**
The json output is the same of input, but whenever mastino found a match, adds a new key `matches` inside `tags` element. The `matches` is array of string ( the DOM ).
Example:
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

**the other language...**
example php:
```php
$jsonManifest = "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}";
$output = shell_exec("/path/to/mastino {$jsonManifest}");
var_dump(json_decode($output, true));
```

example c# :
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

example nodejs:
```js
var sys = require('sys')
var exec = require('child_process').exec;
function puts(error, stdout, stderr) { sys.puts(stdout) }
var manifest = "{\"url\":\"https://blog.golang.org/error-handling-and-go\", \"tags\": [{\"type\": \"div\",\"attributes\": [{\"key\": \"class\",\"value\": \"code\"}]}]}";
exec("/path/to/mastino " + manifest, puts);
```

License
-------
MIT
