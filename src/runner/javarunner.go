package runner;

import (
	"os/exec";
	"path";
	"time";
	"testcase";
)

type JavaRunner struct {
	Test *testcase.TestCase
}

func (r *JavaRunner) Run() error {
	inreader, outwriter, err := r.Test.GetIOStreams();
	defer CloseFiles(inreader, outwriter);
	
	if err != nil {
		return err;
	}
	
	exec_ := path.Join(r.Test.Execpath, r.Test.Execfile);
        cmd := exec.Command("java",exec_);
        ms := time.Duration(r.Test.Maxtime);

        cmd.Stdin = inreader;
        cmd.Stdout = outwriter;
	done := make(chan error, 1);

	//Start execution and wait
	cmd.Start();
	go func() {
		done <- cmd.Wait();
	}();

	select{
	case <- time.After(ms * time.Millisecond):
		return r.HandleTLE(cmd);

	case err := <- done:
		return err;
	}
}

func (r *JavaRunner) HandleTLE(cmd *exec.Cmd) error{
	return cmd.Process.Kill();
}
