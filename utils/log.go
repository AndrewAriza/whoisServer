package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Preload() {
	success := color.New(color.FgGreen).PrintlnFunc()
	success("__          ___           _        _____")
	success("\\ \\        / / |         (_)      / ____|")
	success(" \\ \\  /\\  / /| |__   ___  _ ___  | (___   ___ _ ____   _____ _ __")
	success("  \\ \\/  \\/ / | '_ \\ / _ \\| / __|  \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|")
	success("   \\  /\\  /  | | | | (_) | \\__ \\  ____) |  __/ |   \\ V /  __/ |")
	success("    \\/  \\/   |_| |_|\\___/|_|___/ |_____/ \\___|_|    \\_/ \\___|_|")
	fmt.Println("                       Listening at port", os.Getenv("port"))
}
