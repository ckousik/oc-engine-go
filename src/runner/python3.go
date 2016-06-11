package runner;

import (
	"os";
	"os/exec";
	"path";
	"time";
	"task";
)

type Py3 struct {}

func (p Py3) Start (t *task.TestGroup) chan StatusCode {
	status := make(chan StatusCode, 1);
	go func(){
		
		os.MkdirAll(path.Join(os.Getenv("OC_OUTPUTS"),t.RunId), 0777);
		defer os.RemoveAll(path.Join(os.Getenv("OC_OUTPUTS"),t.RunId));

		testcases, err := t.GenerateTestCases();

		if err != nil {
			status <- CompileError; //Temporary placeholder
			return;
		}
		status <- ExecutionStarted;
		for _, tc := range testcases {

			inreader, outwriter, err := tc.GetIOStreams();
			if err != nil {
				status <- FileError;
				continue;
			}

			exec_command := exec.Command("python3", t.Codefile);
			done := make(chan error, 1);
			exec_command.Stdin = inreader;
			exec_command.Stdout = outwriter;

			exec_command.Start();
			go func(){
				done <- exec_command.Wait();
			}();

			select {

			case err:= <- done :
				if err != nil {
					status <- TestRunError;
					continue;
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
				
			case <- time.After(time.Duration(tc.Maxtime) * time.Millisecond):
				status <- TestTLE;
			}
		}
		status <- ExecutionCompleted;
	}();
	return status;
}
