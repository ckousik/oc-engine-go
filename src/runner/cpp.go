package runner;

import (
	"os";
	"os/exec";
	"test";
	"path";
	"time";
)

type CPP struct {};

func (c CPP) Start (t *test.TestGroup) chan StatusCode {
	status := make(chan StatusCode);

	//Make directory to store outputs
	os.Mkdir(path.Join(os.Getenv("OC_OUTPUTS"),t.RunId),0777);
		
	go func () {
		// Compile
		compile_input := t.Codefile;
		compile_output := path.Join(os.Getenv("OC_EXEC"),t.RunId);
		compile_command := exec.Command("g++",compile_input,"-o",compile_output);
		err := compile_command.Run();

		//Terminate on compilation error
		if err != nil {
			status <- CompileError;
			return;
		}
		
		testCases, err := t.GenerateTestCases();
		for _, tc := range testCases {

			inreader, outwriter, err := tc.GetIOStreams();
			//Error opening file
			if err != nil {
				status <- FileError;
				continue;
			}

			exec_command := exec.Command(compile_output);
			exec_command.Stdin = inreader;
			exec_command.Stdout = outwriter;

			//Run test cases
			done := make(chan error, 1);
			timeout := time.After(time.Duration(tc.Maxtime) * time.Millisecond);
			exec_command.Start();
			status <- ExecutionStarted;
			go func(){
				defer inreader.Close();
				defer outwriter.Close();
				done <- exec_command.Wait();
			}();

			select {
			case err:= <- done :
				if err != nil {
					status <- TestRunError;
				}

				//Compare files here
				res, err := CompareFiles(tc.Outputfile, tc.Testfile);
				if err != nil {
					status <- TestRunError;
					continue;
				}
				if res {
					status <- TestSuccess;
				}else {
					status <- TestFail;
				}
				
				
			case <- timeout:
				status <- TestTLE;
			}
			
			//Close streams
			inreader.Close();
			outwriter.Close();
		}
		status <- ExecutionCompleted;

		//Cleanup
		Cleanup(path.Join(os.Getenv("OC_OUTPUTS"),t.RunId));
		Cleanup(path.Join(os.Getenv("OC_EXEC"),t.RunId));
	}();
	return status;
}
