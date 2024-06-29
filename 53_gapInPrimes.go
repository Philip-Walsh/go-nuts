package main

func Gap(g, m, n int) []int {
    previousPrime := 0

    for i := m; i <= n; i++ {
        if isPrime(i) {
            if previousPrime != 0 && i-previousPrime == g {
                return []int{previousPrime, i}
            }
            previousPrime = i
        }
    }

    return nil
}

func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
