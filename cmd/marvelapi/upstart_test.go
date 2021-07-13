package main

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/go-redis/redis"
)

func TestInitResources(t *testing.T) {
	tests := []struct {
		name      string
		mock      func()
		wantPanic bool
	}{
		{
			name: "Success Init",

			mock: func() {
				monkey.Patch(loadEnv, func() {})
				monkey.Patch(initRedis, func() *redis.Client {
					return nil
				})
			},
			wantPanic: false,
		},
		{
			name: "Fail Init - env",

			mock: func() {
				monkey.Patch(initRedis, func() *redis.Client {
					return nil
				})
			},
			wantPanic: true,
		},
		{
			name: "Fail Init - redis",

			mock: func() {
				monkey.Patch(loadEnv, func() {})
				monkey.Patch(os.Getenv, func(key string) string {
					return "mock"
				})
			},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			assertPanicUpstart(t, InitResources, tt.wantPanic)
			monkey.UnpatchAll()
		})
	}
}

func assertPanicUpstart(t *testing.T, f func() *redis.Client, wantPanic bool) {
	defer func() {
		r := recover()
		if (r != nil) != wantPanic {
			t.Errorf("get panic = %v, wawantPanic %v", r, wantPanic)
		}
	}()

	f()
}
