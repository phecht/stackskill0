package main

import (
	"sync"
	"testing"
)

/* func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
} */
func BenchmarkContextSwitch(b *testing.B) {
	var wq sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wq.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wq.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wq.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wq.Wait()

}
