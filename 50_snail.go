package kata

func Snail(array [][]int) []int {
    if len(array) == 0 || len(array[0]) == 0 {
        return []int{}
    }

    var result []int

    for len(array) > 0 {
        result = append(result, array[0]...)
        array = array[1:]

        if len(array) == 0 {
            break
        }

        for i := 0; i < len(array); i++ {
            result = append(result, array[i][len(array[i])-1])
            array[i] = array[i][:len(array[i])-1]
        }

        if len(array) == 0 {
            break
        }

        for i := len(array[len(array)-1]) - 1; i >= 0; i-- {
            result = append(result, array[len(array)-1][i])
        }
        array = array[:len(array)-1]

        if len(array) == 0 {
            break
        }

        for i := len(array) - 1; i >= 0; i-- {
            result = append(result, array[i][0])
            array[i] = array[i][1:]
        }
    }

    return result
}
