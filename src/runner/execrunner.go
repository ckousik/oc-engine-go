package runner;

import(
	"fmt";
	"os/exec";
	"path";
	"time";
	"errors";
	"testcase";
)

type ExecRunner struct {
	Test *testcase.TestCase
}

func (r *ExecRunner) Run () error {
	inreader, outwriter, err := r.Test.GetIOStreams();

	defer CloseFiles(inreader, outwriter);
	
	if err != nil {
		return err;
	}
	
	exec_ := path.Join(r.Test.Execpath, r.Test.Execfile);

	cmd := exec.Command(exec_);
	cmd.Stdin = inreader;
	cmd.Stdout = outwriter;

	done := make(chan error,1);

	//Start the process and wait
	ms := time.Duration(r.Test.Maxtime);
	
	cmd.Start();
	go func(){
		done <- cmd.Wait();
	}();
	
	select {
	case <- time.After(ms * time.Millisecond):
		return r.HandleTLE(cmd);
	case err := <- done:
		return err;
	}
}

func (r *ExecRunner) HandleTLE (cmd *exec.Cmd) error{
	fmt.Println("Time limit exceeded");
	if err := cmd.Process.Kill(); err != nil {
		return errors.New("Unable to kill process");
	}

	return errors.New("Time Limit Exceeded");
}
