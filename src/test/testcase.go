package test;

/*
This represents a test case to be executed by a Runner

Fields:

Name : string : name of the test

Inputpath : string : Path where the inputs for the test case are stored

Inputfile : string : Name of the file with the inputs for the test case

Outputpath : string : Path where the outputs are stored

Outputfile : string : Name of the file to which the output will be written

Testpath : string : Path where the corect outputs are stored

Testfile : string : File where the correct output for this test case is stored

Maxtime : int64 : Maximum run time in Milliseconds

*/

import (
	"os";
	"path";
	"fmt";
)


type TestCase struct {

	Name string
	
	Inputpath, Inputfile string

	Execpath, Execfile string

	Outputpath, Outputfile string

	Testpath, Testfile string

	Maxtime int64
}

func (t *TestCase) GetIOStreams () (*os.File, *os.File, error) {

	var outwriter *os.File;
	
	in := path.Join(t.Inputpath, t.Inputfile);
	out := path.Join(t.Outputpath, t.Outputfile);
	inreader, err :=  os.Open(in);

	if err != nil {
		return nil, nil, err;
	}

	if _ , err := os.Stat(out); os.IsNotExist(err) {
		fmt.Println("Creating file for writing: ", out);
		outwriter, err = os.Create(out);
	}else {
		fmt.Println("Opening file for writing:", out);
		outwriter, err = os.Open(out);
	}
	
	if err != nil {
		return nil, nil, err;
	}

	return inreader, outwriter, nil;
}
