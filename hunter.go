package main

import "fmt"

func primeFactor(n int) {
    for ; n % 2 == 0; n /= 2 {
        fmt.Print(2, " ")
    }
    for i := 3; i * i <= n; i += 2 {
        for ; n % i == 0; n /= i {
            fmt.Print(i, " ")
        }
    }
    if n > 1 {
        fmt.Println(n)
    }
}

func main() {
    primeFactor(1234567890)
}
