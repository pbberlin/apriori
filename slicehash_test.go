package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)

func EncFmt(in []string) string {
	return fmt.Sprintf("%v", in)
}
func EncJson(in []string) string {
	b, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
func EncJoin(in []string) string {
	return strings.Join(in, "")
}

// var benchInp = []string{"A", "B", "C", "D", "E"}
var benchInp = []string{"Geh", "Du", "alter", "Esel", "Caffee", "trink", "nicht", "so", "viel", "Kaffee"}

func BenchmarkEncFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EncFmt(benchInp)
	}
}
func BenchmarkEncJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EncJson(benchInp)
	}
}
func BenchmarkEncJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EncJoin(benchInp)
	}
}
