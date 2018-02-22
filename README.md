# go-wc

This is my attempt to write a go clone of the BSD `wc` (word count) command ([as shipped by Apple](https://developer.apple.com/legacy/library/documentation/Darwin/Reference/ManPages/man1/wc.1.html)), just as a project for learning go.

It behaves pretty much the same way as `wc`, with some caveats:

- In `wc`, you can pass all of your arguments together (like `wc -wlc` vs `wc -w -l -c`), but that isn't supported here yet
- the only supported arguments are `-w`, `-l`, and `-c`, and `-m`
- `-m` and `-c` are exclusive. the `wc` manpage says that the *last* one of them specified takes precendence. Here, instead, the program simply fails if you try to specify both.

# Compile

`go build cmd/wc/wc.go`

# Usage


Get the lines, words, and bytes of a file:
```sh
$ ./wc foo.txt   
 3 4 27 foo.txt
```

or multiple files...
```sh
$ ./wc foo.txt bar.txt 
 3 4 27 foo.txt
 3 4 27 bar.txt
 6 8 54 total
```
or stdin...
```sh
$ ./wc < foo.txt       
 4 3 27 
$ ls | ./wc 
 5 5 30 
```
You can request just the lines, words, bytes, or characters (really, [runes](https://golang.org/pkg/builtin/#rune)) count with -l, -w, -c or -m

```sh
$ ls | ./wc -l
 5 
$ ./wc -w foo.txt   
 4 foo.txt
$ ./wc -c foo.txt
 27 foo.txt
```
