package main;

import (
	"runner";
	"compiler";
	"path";
	"testcase";
)

// A test run
func main() {
	c := &compiler.GPP{
		Codepath : path.Join("..","hr","pangram"),
	        Codefile : "pangram.cpp",
		Execpath : path.Join("..","hr","pangram"),
		Execfile : "pg",
	}
	err := c.Compile();
	if err != nil {
		panic(err);
	}

	r := &runner.ExecRunner{
		Test: &testcase.TestCase{
			Name: "pangram",
			Execpath : c.Execpath,
			Execfile : c.Execfile,
			Inputpath : c.Codepath,
			Inputfile : "test.txt",
			Outputpath : ".",
			Outputfile : "pangram.txt",
			Maxtime : 2,
		},
	};

	r.Run();
}
