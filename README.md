# Daily Coding Problem: Problem #816 [Medium] 

This problem was asked by Facebook.

Given three 32-bit integers x, y, and b,
return x if b is 1 and y if b is 0,
using only mathematical or bit operations.
You can assume b can only be 1 or 0.

## Analysis

Solving this hinges on knowing that an unsigned integer
in almost all programming languages can "roll over".
For a 64-bit unsigned integer variable,
setting that variable to 0xFFFFFFFFFFFFFFFF and incrementing
it gives it the value 0.

You also have to know that bitwise-and of 0xFFFFFFFFFFFFFFFF with
any value gives back that value.
A bitwise-and of 0x0000000000000000 with any value gives you
a value of zero.

So the "y if b is 0" is available like this:

```go
var zork uint64 
zork = 0xFFFFFFFFFFFFFFFF 
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

## Interview Analysis

There's actually minimal programming involved with this problem.
Solving this involves knowing some particulars:
if the machine/programming language uses 2s-complement arithmetic,
what values cause over- or under-flow,
and the bitwise-and operation.

This may be what the interviewer wants to find out.
If so, this is a good problem.
If the interviewer wants to see some programming,
this is not a good problem for a job interview.
