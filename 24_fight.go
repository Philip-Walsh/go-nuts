package kata

import (
	"codewarrior/kata"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)
type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
  var first Fighter;
  var second Fighter;
  if fighter1.Name == firstAttacker {
    first = fighter1
    second = fighter2
  } else {
    first = fighter2
    second = fighter1
  }
  for first.Health > 0 && second.Health > 0 {
    second.Health -= first.DamagePerAttack
    if second.Health > 0{
      first.Health -= second.DamagePerAttack
    }
  }
  if first.Health <= 0{
      return second.Name
  }
  return first.Name
  
}

var _ = Describe("Sample Tests", func() {
	It("should return the correct values", func() {
		Expect(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew")).To(Equal("Lew"))
		Expect(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry")).To(Equal("Harry"))
		Expect(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry")).To(Equal("Harald"))
		Expect(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald")).To(Equal("Harald"))
		Expect(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry")).To(Equal("Harald"))
		Expect(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald")).To(Equal("Harald"))
	})