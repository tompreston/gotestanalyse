package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type TestEvent struct {
	Time    time.Time `json:"Time"` // encodes as an RFC3339-format string
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"` // seconds
	Output  string    `json:"Output"`
}

type TestResult struct {
	Package string
	Test    string
	Result  string
}

func (t *TestResult) String() string {
	return fmt.Sprintf("%v %v.%v", t.Result, t.Package, t.Test)
}

func (t *TestResult) Name() string {
	return fmt.Sprintf("%v.%v", t.Package, t.Test)
}

func parseTestEvents(filename string) ([]TestEvent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	testEvents := []TestEvent{}

	dec := json.NewDecoder(bufio.NewReader(file))
	for {
		var testEvent TestEvent

		err := dec.Decode(&testEvent)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		testEvents = append(testEvents, testEvent)
	}

	return testEvents, nil
}

// filterTestEvents returns a list of test events we want to use
func filterTestEvents(testEvents []TestEvent) []TestEvent {
	filtered := []TestEvent{}

	for _, t := range testEvents {
		// skip package actions
		if t.Test == "" {
			continue
		}
		switch t.Action {
		case "pass", "skip", "fail":
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func testEventsToTestResults(testEvents []TestEvent) []TestResult {
	testResultsByName := map[string]TestResult{}

	for _, t := range filterTestEvents(testEvents) {
		testResult := TestResult{
			Package: t.Package,
			Test:    t.Test,
			Result:  t.Action,
		}

		// If a test result changes, mark it as flaky
		if oldTestResult, exists := testResultsByName[testResult.Name()]; exists {
			if oldTestResult.Result == "flaky" {
				continue
			}
			if oldTestResult.Result != testResult.Result {
				testResult.Result = "flaky"
				testResultsByName[testResult.Name()] = testResult
				continue
			}
		}

		testResultsByName[testResult.Name()] = testResult
	}

	testResults := []TestResult{}
	for _, v := range testResultsByName {
		testResults = append(testResults, v)
	}
	return testResults
}

func countResults(testResults []TestResult, result string) int {
	count := 0
	for _, tr := range testResults {
		if tr.Result == result {
			count++
		}
	}
	return count
}

func main() {
	testEvents, err := parseTestEvents("test-output.log")
	if err != nil {
		log.Fatalln(err)
	}

	testResults := testEventsToTestResults(testEvents)

	numTests := len(testResults)
	numPass := countResults(testResults, "pass")
	numSkip := countResults(testResults, "skip")
	numFlaky := countResults(testResults, "flaky")
	numFail := countResults(testResults, "fail")

	// TODO report to backend

	fmt.Printf(
		"%v tests, %v pass, %v skip, %v flaky, %v fail\n",
		numTests,
		numPass,
		numSkip,
		numFlaky,
		numFail,
	)

	if numFail > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
