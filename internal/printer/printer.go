package printer

import "github.com/fatih/color"

func Print(violations []string) {
	if len(violations) == 0 {
		successPrint := color.New(color.Bold, color.FgGreen).PrintlnFunc()
		successPrint("PASS")
		return
	}

	failPrint := color.New(color.Bold, color.FgRed).PrintlnFunc()
	for _, v := range violations {
		failPrint(v)
	}
}
