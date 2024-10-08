# Bitwise Operations Interview Questions

* [Choice of two numbers](#daily-coding-problem-problem-816-medium)
* [Swap bits of 8-bit integer](#daily-coding-problem-problem-1780-medium)
* [Maximum of two numbers](#daily-coding-problem-problem-1011-hard)

---

# Daily Coding Problem: Problem #816 [Medium] 

This problem was asked by Facebook.

Given three 32-bit integers x, y, and b,
return x if b is 1 and y if b is 0,
using only mathematical or bit operations.
You can assume b can only be 1 or 0.

## Build and run

I wrote a command line program which accepts arguments in hexadecimal.
The arguments in order are `x y b`

```sh
$ go build a1.go
$ ./a1 fe ef 0
fe      ef      0
result: ef
$ ./a1 fe ef 1
fe      ef      1
result: fe
$
```

## Analysis

Solving this hinges on knowing that an unsigned integer
in almost all programming languages can "roll over".
For a 64-bit unsigned integer variable,
setting that variable to 0xFFFFFFFFFFFFFFFF and incrementing
it gives it the value 0.
The value 0xFFFFFFFF works the same way for 32-bit unsigned integers.

You also have to know that bitwise-and of 0xFFFFFFFF with
any value gives back that value.
A bitwise-and of 0x00000000 with any value gives you
a value of zero.

So the "y if b is 0" is available like this:

```go
var zork uint32
zork = 0xFFFFFFFF 
result := y & (zork + b)
``` 

Variable `result` will end up with the value of `y`,
when `b` contains 0.
`result` will end up with value of 0 when `b` contains 1.

The other part, "x if b is 1", hinges on knowing that
most programming languages use 2s-complement representation
of negative numbers.
A variable with value -1 actually has all 1-bits.
So `0-b` will be all-1-bits when b has value 1 and all-0-bits
when b has value 0.

You can bitwise manipulate x and y values into either 0 or that value,
then add the two to get the desired result.

The hard parts here are knowing a few bitwise operations,
2s-complement arithmetic, and the insight that 0 + x equals x.

Other people have solved this in much the same way,
except [they usually](https://github.com/iamvictorli/Daily-Coding-Problem/blob/master/solutions/81-90/Problem85.js)
thought of creating the and-mask not by under or overflow,
but rather by knowing that setting `b = -b` gives you all 1-bits
if b originally has the value 1, and all 0-bits if b is originally zero.

## Interview Analysis

There's actually minimal programming involved with this problem.
Solving this involves knowing some particulars:
if the machine/programming language uses 2s-complement arithmetic,
what values cause over- or under-flow,
and the bitwise-and operation.

This may be what the interviewer wants to find out.
If the interviewer wants to see some programming,
this is not a good problem for a job interview.

I still think that wanting a job candidate to remember bitwise stunts
in a stressful job interview is too much.
Also, the "medium" rating isn't deserved if you know the trick,
but if you don't know the trick, this is a very difficult problem.

---

# Daily Coding Problem: Problem #1780 [Medium]

This problem was asked by Cisco.

Given an unsigned 8-bit integer, swap its even and odd bits.
The 1st and 2nd bit should be swapped, the 3rd and 4th bit should be swapped, and so on.

For example, `10101010` should be `01010101`. `11100010` should be `11010001`.

Bonus: Can you do this in one line?

## Build and run

I wrote a command line program which accepts arguments in text represention of 8-bit binary numbers.

```sh
$ go build c1.go
$ ./c1 10101010
original: 10101010
swapped:  01010101
swapped:  01010101
```

There's two "swapped" 8-bit integer string representations
because one of them is the result of the bonus "one-liner".

## Analysis

Again, minimal actual programming involved with this problem.
Solving this involves knowing how to mask bits with binary-and,
put bits together with binary-or.
That's it.

I did the code above off the top of my head, and it worked the first compile.
Not to brag, this is a simple problem.

## Interview Analysis

I can't imagine that this gets an interviewer any knowledge
about a candidate's coding skill, except whether they can do bitwise operations.
I suppose the interviewer could get operation priority knowledge out of
a candidate by getting them to remove parentheses in the on-liner.

Maybe a candidate could suggest other test cases,
perhaps to see if every position swaps correctly.

---

# Daily Coding Problem: Problem #1011 [Hard]

This problem was asked by Nvidia.

Find the maximum of two numbers without using any if-else statements,
branching, or direct comparisons.

# Analysis

# Interview Analysis
