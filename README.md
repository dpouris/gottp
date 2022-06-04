# gottp 🐹

## Example

## **get**
```shell
$ gottp get -u https://google.com -f google.html
```

## **post**
```shell
$ gottp post -u https://example.com/p/pjn87vi7h/post -p sample_data.json


-------------------- RESPONSE --------------------

{200 OK 200 HTTP/2.0 2 0 map[Access-Control-Allow-Origin:[*] Cache-Control:[private] Content-Type:[text/plain; charset=utf-8] Date:[Sat, 04 Jun 2022 20:44:51 GMT] Server:[Google Frontend] Vary:[Accept-Encoding] X-Cloud-Trace-Context:[fece1a06eab68ce6b99ae75e89741025]] 0x14000185320 -1 [] false true map[] 0x140000fe000 0x140001ea0b0}

```

## Installation

Probably the easiest way it to build, alias and run the binary. If you want to do it another way feel free.

```shell
$ go build main.go get.go
```
Get your current directory 
```shell
$ pwd

/directory/to/your/project
```
and add /main at the end

```shell
-> /directory/to/your/project/main
```
```shell
$ vi ~/.bashrc # or .zshrc whatever shell you use
```

On the file add the following line
```sh
alias gottp="/directory/to/your/project/main"
```

Save and exit.

Now, you can run:
```shell
$ gottp get -h

Must provide a url to fetch
  -f string
    	Specify the path to save the fetched data
  -h	Learn abou the commands
  -u string
    	Specify the url to fetch data from
```
```shell
$ gottp post -h

Please provide a url to post to.
  -f string
    	Specify the path to save the response
  -h	Learn about the commands
  -p string
    	Specify a path to a JSON/XML file to use as a payload for the request
  -u string
    	Specify the url to post data to

```

Give me a star if you like it!