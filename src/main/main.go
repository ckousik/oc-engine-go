package main;

import (
	"fmt";
	"runner";
	"compiler";
//	"path";
	"testcase";
)

// A test run
func main() {
	t := &testcase.TestGroup {
		Name: "Pangram",
		Id: "test-id",
		Lang: "C++",

		Codepath: "../hr/pangram",
		Codefile: "pangram.cpp",
		Execpath: "../hr/pangram",
		Inputpath: "../hr/pangram",
		Inputfiles: []string{"test.txt"},

		Testpath: "",
		Testfiles: []string{"none"},

		Outputpath: ".",
		Maxtime: 3,
	}
	c := &compiler.GPP{T: t,}
	err := c.Compile();
	if err != nil {
		fmt.Println(err);
		panic(err);
	}

	_, tests := t.GenerateTestCases();
	for _, test := range tests{
		r := &runner.ExecRunner{
			Test: &test,
		};

		err := r.Run();

		if err != nil {
			panic(err);
		}
	}
}
