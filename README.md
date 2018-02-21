# go-wc

This is my attempt to write a go clone of the BSD `wc` (word count) command ([as shipped by Apple](https://developer.apple.com/legacy/library/documentation/Darwin/Reference/ManPages/man1/wc.1.html)), just as a project for learning go.

It behaves pretty much the same way as `wc`, with some caveats:

- There is not (yet) a differentiation between "bytes" and "characters", and no -m argument
- In `wc`, you can pass all of your arguments together (like `wc -wlc` vs `wc -w -l -c`), but that isn't supported here yet
- the only supported arguments are -w, -l, and -c

# Compile

`go build -o mywc`

# Usage


Get the lines, words, and bytes of a file:
```sh
$ ./mywc foo.txt   
 3 4 27 foo.txt
```

or multiple files...
```sh
$ ./mywc foo.txt bar.txt 
 3 4 27 foo.txt
 3 4 27 bar.txt
 6 8 54 total
```
or stdin...
```sh
$ ./mywc < foo.txt       
 4 3 27 
$ ls | ./mywc 
 5 5 30 
```
You can request just the lines, word, or bytes count with -l, -w, or -c

```sh
$ ls | ./mywc -l
 5 
$ ./mywc -w foo.txt   
 4 foo.txt
$ ./mywc -c foo.txt
 27 foo.txt
```
