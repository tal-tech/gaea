package version

import (
	"fmt"
	"os"
)

var (
	TAG        string = ""
	VERSION    string = ""
	BUILD_DATE string = ""
	AUTHOR     string = ""
	BUILD_INFO string = ""
)

func Version() {
	fmt.Fprintf(os.Stderr, "tag\t%s\n", TAG)
	fmt.Fprintf(os.Stderr, "version\t%s\n", VERSION)
	fmt.Fprintf(os.Stderr, "build\t%s\n", BUILD_DATE)
	fmt.Fprintf(os.Stderr, "author\t%s\n", AUTHOR)
	fmt.Fprintf(os.Stderr, "info\t%s\n", BUILD_INFO)
}
