package task;

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
	"fmt";
)


type TestCase struct {

	Inputfile string

	Outputfile string

	Testfile string

	Maxtime int64
}

func (t *TestCase) GetIOStreams () (*os.File, *os.File, error) {

	var outwriter *os.File;
	inreader, err :=  os.Open(t.Inputfile);

	if err != nil {
		return nil, nil, err;
	}

	if _ , err := os.Stat(t.Outputfile); os.IsNotExist(err) {
		fmt.Println("Creating file for writing: ", t.Outputfile);
		outwriter, err = os.Create(t.Outputfile);
	}else {
		fmt.Println("Opening file for writing:", t.Outputfile);
		outwriter, err = os.Open(t.Outputfile);
	}
	
	if err != nil {
		return nil, nil, err;
	}

	return inreader, outwriter, nil;
}
