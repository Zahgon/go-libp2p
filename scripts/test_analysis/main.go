package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "modernc.org/sqlite"
)

const dbPath = "./test_results.db"
const retryCount = 4

var coverRegex = regexp.MustCompile(`-cover`)

func main() {
	var t tester
	if len(os.Args) >= 2 {
		if os.Args[1] == "summarize" {
			md, err := t.summarize()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(md)
			return
		}
	}

	passThruFlags := os.Args[1:]
	err := t.runTests(passThruFlags)
	if err != nil {
		log.Fatal(err)
	}
}

type tester struct {
	Dir string
}

func (t *tester) runTests(passThruFlags []string) error { _ = "STUB: not implemented"; return nil }

func (t *tester) goTestAll(extraFlags []string) error { _ = "STUB: not implemented"; return nil }

func (t *tester) goTestPkgTest(pkg, testname string, extraFlags []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tester) goTest(extraFlags []string) error { _ = "STUB: not implemented"; return nil }

type failedTest struct {
	Package string
	Test    string
}

type timedOutPackage struct {
	Package string
	Outputs string
}

func (t *tester) findFailedTests(ctx context.Context) ([]failedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tester) findTimedoutTests(ctx context.Context) ([]timedOutPackage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterOutFlags(flags []string, exclude *regexp.Regexp) []string {
	_ = "STUB: not implemented"
	return nil
}

func (t *tester) summarize() (string, error) { _ = "STUB: not implemented"; return "", nil }
