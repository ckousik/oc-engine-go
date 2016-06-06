package runner;

import(
	"fmt";
	"os/exec";
	"path";
	"time";
	"errors";
	"test";
)

type ExecRunner struct {
	*test.TestCase
}

func (r *ExecRunner) Run () error {
	inreader, outwriter, err := r.GetIOStreams();

	defer CloseFiles(inreader, outwriter);
	
	if err != nil {
		return err;
	}
	
	exec_ := path.Join(r.Execpath, r.Execfile);

	cmd := exec.Command(exec_);
	cmd.Stdin = inreader;
	cmd.Stdout = outwriter;

	done := make(chan error,1);

	//Start the process and wait
	timeout := time.After(time.Duration(r.Maxtime) * time.Millisecond);
	cmd.Start();
	go func(){
		done <- cmd.Wait();
	}();
	
	select {
	case <- timeout:
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
