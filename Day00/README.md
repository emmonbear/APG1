# Day 00 — Go Boot camp

## Statistics being handy

💡 [Tap here](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) **to leave your feedback on the project**. It's anonymous and will help our team make your educational experience better. We recommend completing the survey immediately after the project.

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Anscombe](#exercise-00-anscombe)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not exit unexpectedly (give an error on valid input). If this happens, your project will be considered non-functional and will receive a 0 in the evaluation.
- We encourage you to create test programs for your project, even though this work doesn't have to be submitted and won't be judged. This will allow you to easily test your work and the work of your peers. You will find these tests particularly useful during your defence. In fact, you are free to use your own tests and/or those of the peer you are evaluating during your defence.
- Submit your work to your allocated git repository. Only work in the git repository will be marked.
- If your code uses external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) to manage them.

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the day</h2>

- You should only submit `*.go` files and (in the case of external dependencies) `go.mod` + `go.sum`.
- Your programme should accept a sequence of numbers separated by newlines as standard input. A number is also a sequence.
- You can assume that it should only work with integers (but the output can be floats, rounded to 2 decimal places).
- Nevertheless, it should print a meaningful error message without panicking at runtime if it is fed some unexpected input, such as an out-of-bounds number, a letter, or an empty string.
- Your code for this task should be buildable with just `go build`.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

Go isn't generally thought of as a Data Science language. But that doesn't mean it can't crunch numbers. In fact, it's comparable to C for basic tasks. It can also be a lot easier to write, partly because GC handles memory management, and partly because Go's standard library is pretty good. We're constantly taught that it can be a bad idea to just trust your gut when dealing with important data. To make sense of a sample of numbers, it's usually better to use a statistical approach. Data can sometimes be deceptive, like [Anscombe's quartet](https://en.wikipedia.org/wiki/Anscombe%27s_quartet), but the more metrics we get, the more weighted decision we'll be able to make in the end, right?

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Anscombe</h3>


So let's say we have a bunch of integers, strictly between -100000 and 100000. It's probably a large set, so let's assume our application reads it from a standard input, separated by newlines. For now, let's think of four important statistical metrics that we can derive from this data, so that by default we can print all of them as a result, for example like this:

```
Mean: 8.2
Median: 9.0
Mode: 3
SD: 4.35
```

The order and actual format doesn't really matter, as long as we can understand which is which. \
A few notes, though:

1) The input data can be sorted or not. You don't need to write your own sorting algorithm, luckily Go already has one in the standard library, and it works for integers.
2) Median is a middle number of a sorted sequence if its size is odd, and an average between two middle numbers if its size is even.
3) Mode is the most frequent number, and if there are several, the smallest one is returned. You might think of using some structure to store number counts, and some standard Go container might help.
4) You can use both population and regular standard deviation, whichever you prefer.
5) Calling someone "average" can be mean.

It will also make sense for the user to be able to specifically choose which of these four parameters to print, so this needs to be implemented as well. By default, it's all of them, but there should be a way to specify whether to print only one, two or three specific metrics out of four when running the program (without recompiling). There is a built-in module in the standard library that allows you to parse additional parameters.
