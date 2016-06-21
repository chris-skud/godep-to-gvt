package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
)

var fp = flag.String("fp", "", "path to manifest file")
var latest = flag.Bool("latest", false, "fetch latest version of dep")

func init() {
	flag.Parse()
}

type G struct {
	Deps []Dep
}
type Dep struct {
	ImportPath string
	Comment    string
	Rev        string
}

func main() {
	fpath := fmt.Sprintf("%s/Godeps/Godeps.json", *fp)

	fbytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err)
	}

	var g G
	err = json.Unmarshal(fbytes, &g)
	if err != nil {
		panic(err)
	}

	for _, dep := range g.Deps {
		var gvt *exec.Cmd
		if *latest {
			fmt.Printf("fetching latest from %s. ", dep.ImportPath)
			gvt = exec.Command("gvt", "fetch", dep.ImportPath)
		} else {
			fmt.Printf("fetching revision %s from %s. ", dep.Rev, dep.ImportPath)
			revParam := fmt.Sprintf("-revision=%s", dep.Rev)
			gvt = exec.Command("gvt", "fetch", revParam, dep.ImportPath)
		}

		stdout, err := gvt.StderrPipe()
		if err != nil {
			fmt.Printf(err.Error())
			//log.Fatal(err)
		}
		if err := gvt.Start(); err != nil {
			fmt.Printf(err.Error())
			// log.Fatal(err)
		}

		b, _ := ioutil.ReadAll(stdout)
		fmt.Printf("%s\n", string(b))

		if err := gvt.Wait(); err != nil {
			fmt.Printf("%s\n", err.Error())
			// log.Fatal(err)
		}

	}

}
