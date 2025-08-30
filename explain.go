package roundly

import (
	"fmt"
	"runtime"
)

func writeExplanation(opts *RenderOptions) error {
	err := explainFunc()
	if err != nil {
		return err
	}

	return nil
}

func explainFunc() error {
	fmt.Println()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	for i := 0; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		frames := runtime.CallersFrames([]uintptr{pc})
		frame, _ := frames.Next()

		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)

		fmt.Printf("%s\n", fn.Name())
	}

	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println()

	return nil
}
