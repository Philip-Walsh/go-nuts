package main
func MaximumSubarraySum(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    if allNegative(numbers) {
        return 0
    }

    return kadane(numbers)
}

func allNegative(numbers []int) bool {
    for _, num := range numbers {
        if num > 0 {
            return false
        }
    }
    return true
}

func kadane(numbers []int) int {
    maxSum := numbers[0]
    currentSum := 0

    for _, num := range numbers {
        currentSum += num
        if currentSum > maxSum {
            maxSum = currentSum
        }
        if currentSum < 0 {
            currentSum = 0
        }
    }

    return maxSum
}
