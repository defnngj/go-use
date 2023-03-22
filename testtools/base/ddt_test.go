package base

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/voicera/tester/assert"
	"github.com/voicera/tester/ddt"
)

const vanillaContent = `{
  "testCases": [
    {
      "id": "The Ultimate Question",
      "input": {
        "question": "What do you get when you multiply six by nine?",
        "timeoutInHours": 65700000000,
        "config": {"base": 13}
      },
      "expected": {
        "answer": "42",
        "error": null
      }
    },
    {
      "id": "Ask Again",
      "input": {
        "question": "?",
        "timeoutInHours": 65700000000
      },
      "expected": {
        "answer": "42"
      }
    }
  ]
}`

func TestLoadTestCasesFromDerivedJSONFileWhenFileIsNotFound(t *testing.T) {
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActualError(err).IsNotNil().Passed() {
		expected := "open _ddt/TestLoadTestCasesFromDerivedJSONFileWhenFileIsNotFound.json: no such file or directory"
		assert.For(t).ThatActualError(err).Equals(assert.ErrorString(expected))
	}
}

func TestLoadTestCasesFromDerivedJSONFileWhenSchemaIsInvalid(t *testing.T) {
	mustWriteJSONFile("TestLoadTestCasesFromDerivedJSONFileWhenSchemaIsInvalid.json", "{")
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActualError(err).IsNotNil().Passed() {
		assert.For(t).ThatActualError(err).Equals(assert.ErrorString("unexpected end of JSON input"))
	}
}

func TestLoadTestCasesFromDerivedJSONFileWhenTestCasesAreMissing(t *testing.T) {
	mustWriteJSONFile("TestLoadTestCasesFromDerivedJSONFileWhenTestCasesAreMissing.json", "{}")
	err := ddt.LoadTestCasesFromDerivedJSONFile(nil)
	if assert.For(t).ThatActualError(err).IsNotNil().Passed() {
		expected := "ddt: cannot load test cases from TestLoadTestCasesFromDerivedJSONFileWhenTestCasesAreMissing.json"
		assert.For(t).ThatActualError(err).Equals(assert.ErrorString(expected))
	}
}

func mustWriteJSONFile(fileName string, content string) {
	if err := ioutil.WriteFile("_ddt/"+fileName, []byte(content), os.ModePerm); err != nil {
		panic(err)
	}
}
