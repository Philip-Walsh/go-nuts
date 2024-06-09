package main

func solve(str string) int {
  var highestScore int = 0
  var rollingScore int = 0
  for _, char := range str {
    var charIndex = int(char) - 96
    if isVowel(charIndex) {
      if highestScore < rollingScore {
         highestScore = rollingScore
      }
      rollingScore = 0
    } else {
        rollingScore += charIndex
      }
  }
  if highestScore < rollingScore {
     highestScore = rollingScore
  }
  return highestScore
}

func isVowel(index int) bool{
  switch index {
    case
        1,
        5,
        9,
        15,
        21:
        return true
    }
    return false
}