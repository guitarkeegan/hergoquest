package announce

import "fmt"

func Log(format string, as ...any) {
	fmt.Printf(format+"\n", as...)
}
