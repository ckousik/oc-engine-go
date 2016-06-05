package runner;

import (
	"fmt";
	"os";
	"os/exec";
)
type Runner interface{
	Run();
	HandleTLE(*exec.Cmd);
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

