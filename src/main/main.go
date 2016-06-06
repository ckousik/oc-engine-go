package main;

import (
	"fmt";
	"runner";
	"compiler";
//	"path";
	"test";
	"sync";
)

// A test run
func main() {
	testGroup := &test.TestGroup {
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
	
	c := &compiler.GPP{ testGroup };
	err := c.Compile();
	if err != nil {
		fmt.Println(err);
		panic(err);
	}

	tests, _ := testGroup.GenerateTestCases();
	var wg sync.WaitGroup;
	
	for _, t := range tests{
		wg.Add(1);
		go func(testCase *test.TestCase){
			defer wg.Done();
			r := &runner.ExecRunner{
				testCase,
			};
			r.Run();
		}(&t)
	}

	wg.Wait();
}
