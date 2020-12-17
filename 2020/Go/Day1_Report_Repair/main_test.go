package main

import "testing"

//short report of expenses
var withExpenses2 = `
pkg: Day1_Report_Repair
BenchmarkWithMapPart1
BenchmarkWithMapPart1-16       	   68224	     17235 ns/op	    9709 B/op	      28 allocs/op
BenchmarkWithoutMapPart1
BenchmarkWithoutMapPart1-16    	   77900	     15576 ns/op	    5008 B/op	       3 allocs/op
BenchmarkWithoutMapPart2
BenchmarkWithoutMapPart2-16    	   72426	     16413 ns/op	    5024 B/op	       3 allocs/op
BenchmarkWithMapPart2
BenchmarkWithMapPart2-16       	    4536	    251268 ns/op	    9724 B/op	      28 allocs/op
PASS
`

//longer report of expenses
var withExpenses = `
pkg: Day1_Report_Repair
BenchmarkWithMapPart1
BenchmarkWithMapPart1-16       	   12412	     97507 ns/op	   47469 B/op	      28 allocs/op
BenchmarkWithoutMapPart1
BenchmarkWithoutMapPart1-16    	    7113	    176287 ns/op	   59408 B/op	       3 allocs/op
BenchmarkWithoutMapPart2
BenchmarkWithoutMapPart2-16    	     996	   1148998 ns/op	   59424 B/op	       3 allocs/op
BenchmarkWithMapPart2
BenchmarkWithMapPart2-16       	    3799	    325517 ns/op	   47488 B/op	      28 allocs/op
PASS
`

func BenchmarkWithMapPart1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		withMapPart1()
	}
}

func BenchmarkWithoutMapPart1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		withoutMapPart1()
	}
}

func BenchmarkWithoutMapPart2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		withoutMapPart2()
	}
}

func BenchmarkWithMapPart2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		withMapPart2()
	}
}
