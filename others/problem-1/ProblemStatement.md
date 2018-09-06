Start with 2 integers L and R. Their values are initially set to L = 0 and R = 1.

You can change their values by performing following commands
Command 'L' will change the value of L to 2*L-R
Command 'R" will change the value of R to 2*R-L

You are given an integer N. Your goal is to find the shortest sequence of
commands after which L=N or R=N

For example consider N =-11. Shortest sequence of commands required to set
either L or R to -11 is ('L', 'L', 'R', 'L). After each command, the value
of Land R are as follows
Initial value L= 0, R = 1
After command 'L': L = -1, R = 1
After command 'L: L = -3, R = 1
After command 'R': L = -3, R = 5
After command 'L': L =-11, R = 5

Hence, at 4th command, we get L=N

Write a function solution(N) which will return a string describing the
shortest sequence of commands required in order to set L = N or R = N. Each
character in the string is either 'L' or 'R' and corresponds the command.
If there is more than 1 correct answer, solution function can return any
one of those values. The function should return string 'impossible' if the
conditions L = N or R = N can not be met.
