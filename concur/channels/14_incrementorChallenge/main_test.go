package main

import (
	"reflect"
	"testing"
)

func Test_incrementor(t *testing.T) {
	type args struct {
		incQty int
	}
	tests := []struct {
		name string
		args args
		want <-chan string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementor(tt.args.incQty); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incrementor() = %v, want %v", got, tt.want)
			}
		})
	}
}
