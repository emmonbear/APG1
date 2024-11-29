# Day 02 — Go Boot camp

## Not Invented Here Syndrome

💡 [Tap here](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) **to leave your feedback on the project**. It's anonymous and will help our team make your educational experience better. We recommend completing the survey immediately after the project.

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Finding Things](#exercise-00-finding-things)
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01: Counting Things](#exercise-01-counting-things)
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02: Running Things](#exercise-02-running-things)
7. [Chapter VII](#chapter-vii) \
    7.1. [Exercise 03: Archiving Things](#exercise-03-archiving-things)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not exit unexpectedly (give an error on valid input). If this happens, your project will be considered non-functional and will receive a 0 in the evaluation.
- We encourage you to create test programs for your project, even though this work doesn't have to be submitted and won't be graded. This will allow you to easily test your work and the work of your peers. You will find these tests particularly useful during your defense. In fact, you are free to use your tests and/or the tests of the peer you are evaluating during your defense.
- Submit your work to your assigned git repository. Only the work in the git repository will be evaluated.
- If your code uses external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) to manage them.

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the day</h2>

- You should only turn in `*.go` files and (in case of external dependencies) `go.mod` + `go.sum`.
- Your code for this task should be buildable with just `go build`.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

It's really amazing how much you can do just by using command line utilities! Pretty much every OS, including embedded ones, has its own CLI and a bunch of little programs that do magical things. For example, you can read about [BusyBox](https://en.wikipedia.org/wiki/BusyBox), which is basically a Swiss army knife for a variety of systems, starting with Linux-based routers on OpenWRT and going all the way up to Android phones.

We're not trying to reinvent the wheel here, but knowing how to work with FS and do basic CLI stuff in Golang can really help, so let's spend some time on that.

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Finding Things</h3>

As a first step, let's implement a `find`-like utility using Go. It needs to take a path and a set of command line options to be able to find different types of entries. We are interested in three types of entries, which are directories, regular files and symbolic links. So we should be able to run our program as follows:

```
# Find all files/directories/symlinks recursively in directory /foo
~$ ./myFind /foo
/foo/bar
/foo/bar/baz
/foo/bar/baz/deep/directory
/foo/bar/test.txt
/foo/bar/buzz -> /foo/bar/baz
/foo/bar/broken_sl -> [broken]
```

or specify `-sl`, `-d` or `-f` to print only symlinks, only directories or only files. Note that the user should be able to explicitly specify one, two or all three of these options, e.g. `./myFind -f -sl /path/to/dir` or `./myFind -d /path/to/other/dir`.

You should also implement another option — `-ext` (works ONLY if -f is specified) to allow users to print only files with a certain extension. An extension in this task can be thought of as the last part of the filename when we split it by a dot. So,

```
# Find only *.go files and ignore all others.
~$ ./myFind -f -ext 'go' /go
/go/src/github.com/mycoolproject/main.go
/go/src/github.com/mycoolproject/magic.go
```

You'll also need to resolve symlinks. So if `/foo/bar/buzz` is a symlink pointing to another place in FS, like `/foo/bar/baz`, print both paths separated by `->`, like in the example above. 

Another thing about symlinks is that they can be broken (pointing to a non-existent file node). In this case, your code should print `[broken]` instead of the path of a symlink destination.

Files and directories that the current user doesn't have access to (permission errors) should be skipped in the output and not cause a runtime error.

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Exercise 01: Counting Things</h3>

Now we are able to find our files, but we may need more meta information about what is in these files. Let's implement a `wc`-like utility to collect basic statistics about our files.

First of all, let's assume that our files are utf-8 encoded text files, so your code should also work with text in Russian (forget about special cases like Arabic for now, only English and Russian are needed). You can also ignore punctuation and just use spaces as the only word delimiter.

You'll need to implement three mutually exclusive (only one can be specified at a time, otherwise you'll get an error message) flags for your code: `-l` for counting lines, `-m` for counting characters, and `-w` for counting words. Your program should be executable in this way:

```
# Count words in file input.txt
~$ ./myWc -w input.txt
777 input.txt
# Count lines in files input2.txt and input3.txt
~$ ./myWc -l input2.txt input3.txt
42 input2.txt
53 input3.txt
# Count characters in files input4.txt, input5.txt and input6.txt
~$ ./myWc -m input4.txt input5.txt input6.txt
1337 input4.txt
2664 input5.txt
3991 input6.txt
```

As you can see, the answer is always a calculated number and a filename separated by a tab (`\t`). If no flags are given, the `-w` behavior should be used.

**Important**: Since all files are independent, you should use goroutines to process them simultaneously. You can start as many goroutines as there are input files specified for the program.

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Running Things</h3>

Do you know what `xargs` is? You can read about it [here](https://shapeshed.com/unix-xargs/), for example. Let's implement a similar tool — in this exercise, you will have to write a utility that:

1) treats all parameters as a command, like `wc -l` or `ls -la`;
2) build a command by appending all the lines that are fed into the program's stdin as arguments to that command, and then run it. So if we run:

```
~$ echo -e "/a\n/b\n/c" | ./myXargs ls -la
```

it should be equivalent to running:

```
~$ ls -la /a /b /c
```

You can test this tool together with those from the previous exercises, so:

```
~$ ./myFind -f -ext 'log' /path/to/some/logs | ./myXargs ./myWc -l
```

will recursively calculate line counts for all ".log" files in the directory `/path/to/some/logs`.

<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="ex03">Exercise 03: Archiving Things</h3>

The last tool we'll implement today is the Log Rotation tool. "Log rotation" is a process by which the old log file is archived and stored so that logs don't pile up in a single file indefinitely. It should work like this:

```
# Will create file /path/to/logs/some_application_1600785299.tag.gz
# where 1600785299 is a UNIX timestamp from `some_application.log`'s [MTIME](https://linuxize.com/post/linux-touch-command/)
~$ ./myRotate /path/to/logs/some_application.log
```

```
# Will create two tar.gz files with timestamps (one for each log) 
# and put them in /data/archive directory
~$ ./myRotate -a /data/archive /path/to/logs/some_application.log /path/to/logs/other_application.log
```

As in Exercise 01, you should use goroutines to parallelize the archiving of multiple files.
