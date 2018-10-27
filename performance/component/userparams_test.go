package component

import (
	"testing"
	"math/rand"
	"fmt"
)

func TestUserParams_Notify(t *testing.T) {

	s2 := rand.NewSource(20)
	r2 := rand.New(s2)
	flag := 0
	for {
		once.Do(func() {
			flag = 1
		})
		println(flag)
		fmt.Println(r2.Intn(100))
	}

}
