package main

import (
	"testing"

	"bou.ke/monkey"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name      string
		mock      func()
		wantPanic bool
	}{
		{
			name: "Run Main",

			mock: func() {
				monkey.Patch(InitResources, func() *redis.Client {
					return nil
				})
				monkey.Patch(Run, func(address string, router *httprouter.Router) {})
			},
			wantPanic: false,
		},
		{
			name:      "Run Main without mock",
			mock:      func() {},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			assertPanicMain(t, main, tt.wantPanic)
			monkey.UnpatchAll()
		})
	}
}

func assertPanicMain(t *testing.T, f func(), wantPanic bool) {
	defer func() {
		r := recover()
		if (r != nil) != wantPanic {
			t.Errorf("get panic = %v, wawantPanic %v", r, wantPanic)
		}
	}()

	f()
}
