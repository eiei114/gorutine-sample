package runtime_trace

import (
	"os"
	"runtime/trace"
)

func main() {
	// トレースの開始
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// トレースしたくないコード
	Work()
}

// Work は何かの作業を表現する関数
func Work() {
	for i := 0; i < 10000; i++ {
		_ = i * i
	}
}
