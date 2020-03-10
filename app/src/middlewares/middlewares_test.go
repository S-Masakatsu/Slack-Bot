package middlewares

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("-------- before test --------\n\n")
	code := m.Run()
	fmt.Printf("\n-------- after test ---------\n")
	os.Exit(code)
}

func TestGetEnv(t *testing.T) {
	t.Helper()

	// 存在しない環境変数を取りに行った際正常にpanicが発生するかどうかのサブテストです。
	t.Run("Data Not exists", func(t *testing.T) {
		n := "SAMPLE"
		defer func() {
			err := recover()
			msg := []string{"missing required environment variable: ", n}
			if err != strings.Join(msg, "") {
				t.Errorf("got %v\nwant %v", err, "illegal processing")
			}
		}()
		t.Logf("%v:%v / pattern: %v / expected: %v\n", 0, "DataNotExists", n, "panic()")
		GetEnv(n)
	})

	// 環境変数の値が正常に取れるかどうかのサブテストです。
	t.Run("Data exists", func(t *testing.T) {
		t.Helper()
		patterns := []struct {
			name     string
			data     string
			expected string
		}{
			{"SAMPLE", "sample", "sample"},
		}

		for i, k := range patterns {
			os.Setenv(k.name, k.data)
			t.Logf("%v:%v / pattern: %v / expected: %v\n", i, "DataExists", k.name, k.data)
			v := GetEnv(k.name)
			if k.expected != v {
				t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", i, k.expected, v, "GetEnv")
			}
		}
	})
}
