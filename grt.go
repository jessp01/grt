package grt

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
)

func ReplaceTokens(inputFilePath, outputFilePath, repoHost, repoOrg, repoName, newRepo string){
	parts := strings.Split(newRepo, "/")
	newRepoHost := parts[0]
	newRepoOrg := parts[1]
	newRepoName := parts[2]
	log.Printf("repoHost: %s, repoOrg: %s, repoName: %s, newRepoHost: %s, newRepoOrg: %s, newRepoName: %s\n", repoHost, repoOrg, repoName, newRepoHost, newRepoOrg, newRepoName)
	inplace := false

	if outputFilePath == ""{
		inplace = true
		tempFileH, err := ioutil.TempFile(filepath.Dir(inputFilePath),"")
		if err != nil{
		    log.Fatal(err.Error())
		}
		outputFilePath = tempFileH.Name()
	}
	_, err := script.File(inputFilePath).Replace(repoHost, newRepoHost).Replace(repoOrg, newRepoOrg).Replace(repoName, newRepoName).WriteFile(outputFilePath)

	if err != nil{
	    log.Fatal(err.Error())
	}

	if inplace {
	    script.File(outputFilePath).Stdout()
	    log.Printf("Moving %s to %s\n", outputFilePath, inputFilePath)
	    e :=os.Rename(outputFilePath, inputFilePath)
	    if e != nil {
		    log.Fatal(e)
	    }
	}
}
