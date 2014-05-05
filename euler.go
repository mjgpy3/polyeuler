package main

import (
	"fmt";
	"flag";
)


// Euler #1
// Answer: 233168
//
// If we list all the natural numbers below 10 that are multiples of 3
// or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
func euler1() int {
	var sum int;
	for i := 3; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i;
		}
	}
	return sum;
}


// Euler #2:
// Answer: 4613732
//
// Each new term in the Fibonacci sequence is generated by adding the
// previous two terms. By starting with 1 and 2, the first 10 terms
// will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// Find the sum of all the even-valued terms in the sequence which do
// not exceed four million.
func euler2() int {
	a := int(1);
	b := int(2);
	sum := 2;
	for b < 4000000 {
		c := a + b;
		if c % 2 == 0 {
			sum += c
		}
		a = b;
		b = c
	}
	return sum
}

func show(n string, r int) {
	fmt.Printf("%s: %d\n", n, r);
}

func main() {
	flag.Parse();
    if flag.NArg() <= 0 {
		show("1", euler1());
		show("2", euler2());
	} else {
		for i := 0; i < flag.NArg(); i++ {
			var r int;
			arg := flag.Arg(i);
			switch arg {
			case "1": r = euler1();
			case "2": r = euler2();
			default: r = -1;
			}
			show(arg, r);
		}
	}
}