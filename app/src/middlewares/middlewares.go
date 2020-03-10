package middlewares

import (
	"os"
	"strings"
)

// 指定した環境変数から値を取り出し返します。
// 環境変数が存在しない場合はerrorが発生します。
func GetEnv(name string) (v string) {
	v = os.Getenv(name)
	if v == "" {
		err := []string{"missing required environment variable: ", name}
		panic(strings.Join(err, ""))
	}
	return
}
