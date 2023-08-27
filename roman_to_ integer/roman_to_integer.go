package main

import "fmt"

func roman_to_integer(s string)int{
	roman := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    res := 0

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
            res -= roman[s[i]]
        } else {
            res += roman[s[i]]
        }
    }

    return res
}

func main(){
	s:="MCMXCIV"
	fmt.Println(roman_to_integer(s))
}