package grt

import (
	"testing"

	"github.com/bitfield/script"
)

func TestTokenReplacement(t *testing.T) {
	inputFilePath := "README.tmpl.md"
	outputFilePath := "README_test.md"
	newRepoName := "grt"
	expectedMatches := 8
	ReplaceTokens(inputFilePath, outputFilePath, "@REPO_HOST@", "@REPO_ORG@", "@REPO_NAME@", "github.com/jessp01/"+newRepoName)
	n, _ := script.File(outputFilePath).Match(newRepoName).CountLines()
	if n != expectedMatches {
		t.Errorf("\nInput [%#v]\nExpected to find %#v %#v times\nGot %#v matches\n",
			inputFilePath, newRepoName, expectedMatches, n)
	}
}
