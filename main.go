package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	Conf.Carts.Sort()
	Conf.Carts.Fill()
	// Conf.Carts.String()
	fmt.Printf("%v", Conf.Carts)

	a := []string{"A", "B"}
	a = []string{"A", "B", "C"}
	a = []string{"A", "B", "C", "D"}
	a = []string{"A", "B", "C", "D", "E"}
	// a = []string{"A", "B", "C", "D", "E", "F"}

	allSgs := permutateElements(a, 2)

	// str := util.IndentedDump(allSgs)
	// lglc(str)

	for k, v := range allSgs {
		if len(v) > 0 {
			fmt.Printf("%v ", k)
			for _, v1 := range v {
				str := fmt.Sprintf("%v", v1)
				fmt.Printf("%v ", str)
			}
			fmt.Printf("\n")
		}
	}

}

func checkErr(err error, desc string) {
	if err != nil {
		log.Fatalf("err %v -  %v", desc, err)
	}
}
