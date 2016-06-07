package runner;

import (
	"fmt";
	"os";
	"os/exec";
	"test";
)
type Runner interface{
	Run() error;
	HandleTLE(*exec.Cmd) error;
	SetTestCase (*test.TestCase);
}


func CloseFiles (inreader *os.File,outwriter *os.File) {
	var err error;

	err = inreader.Close();
	if err != nil {
		fmt.Println("Error closing input file");
	}

	err = outwriter.Close();
	if err != nil {
		fmt.Println("Error closing output file");
	}

}

