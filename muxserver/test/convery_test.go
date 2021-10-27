package test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvey(t *testing.T) {
	Convey("test convey", t, func() {
		a := map[string]string{
			"first": "111",
			"two":   "222",
		}
		for _, value := range a {
			go func(value string) {
				fmt.Println(value)
			}(value)
		}
		for val := range a {
			fmt.Println(val)
		}
		time.Sleep(3 * time.Second)
	})
}

type Slice []int

func NewSlice() Slice {

	return make(Slice, 0)

}

func (s *Slice) Add(elem int) *Slice {

	*s = append(*s, elem)

	fmt.Print(elem)

	return s

}

func TestTwo(t *testing.T) {

	s := NewSlice()

	defer s.Add(1)

	s.Add(3)
	fmt.Println("===")

	d := time.Second * 3
	ti := time.NewTimer(d)

	for {
		ti.Reset(d)
		select {
		case <-ti.C:
			fmt.Println("123213123")
		}
	}
}
