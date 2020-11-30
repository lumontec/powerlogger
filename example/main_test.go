package main

import (
	"fmt"
	"runtime"
	"testing"
)

var fmap = make(map[location]string, 0)

type location struct {
	file string
	line int
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate()
	}
}

func Calculate() int {
	i := 1 + 1
	return i
}

// func BenchmarkLogPackage(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Println("ciao")
// 	}
// }

// ~638 ns/op
//func BenchmarkLogZap(b *testing.B) {
//
//	logger, _ := zap.NewProduction()
//	defer logger.Sync()
//	for i := 0; i < b.N; i++ {
//		logger.Info("failed to fetch URL",
//			// Structured context as strongly typed Field values.
//			zap.String("url", "ciao"),
//			zap.Int("attempt", 3),
//		)
//	}
//}

func BenchmarkStackTrace(b *testing.B) {
	// fmap[string[]{"/home/crash/Documents/local/pers/radlog/main_test.go","2"}] = "ciao"
	fmap[location{"/home/crash/Documents/local/pers/radlog/main_test.go", 2}] = "ciao"
	for i := 0; i < 1; i++ {
		sub3()
	}
}

func sub3() {
	sub4()
}

func sub4() {
	sub5()
}

func sub5() {
	sub6()
}

func sub6() {
	sub7()

}

func sub7() {
	fmt.Println("MAIN FUNCTION")
	printCaller()
	go func() {
		fmt.Println("GOROUTINE")
		printCaller()
	}()
}

func printCaller() {
	counter := make([]uintptr, 4)
	runtime.Callers(2, counter)
	// n := runtime.Callers(0, counter)
	frames := runtime.CallersFrames(counter)

	// i := 0
	frame, _ := frames.Next()

	// for {
	// 	i++
	fmt.Println("frame:", frame.Function)
	// 	frame, _ = frames.Next()
	// 	if i > n {
	// 		break
	// 	}
	// }

	// level := fmap[location{frame.File, frame.Line}]
}
