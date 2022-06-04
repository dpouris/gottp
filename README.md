# gottp 🐹

## Example

```shell
$ gottp get -u https://google.com -f google.html
```

```shell
$ ls

Applications                          Postman Agent
Creative Cloud Files                  Public
Desktop                               Sites
Downloads                             bin
Library                               test.py
Movies                                go
Music                            ~~>  google.html
Pictures                              test.js
Postman                               
```

## Installation

Probably the easiest way it to build, alias and run the binary. If you want to do it another way feel free.

```shell
$ go build main.go get.go
```
```shell
$ vi ~/.bashrc # or .zshrc whatever shell you use
```
Get your current directory 
```shell
$ pwd

/Users/<your-username>
```
and add /main at the end
```shell
-> /Users/<your-username>/main
```

On the file add the following line
```sh
alias gottp="/Users/<your-username>/main"
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

Give me a star if you like it!