package runner;

import (
	"os/exec";
	"path";
	"time";
	"test";
)

type JavaRunner struct {
	*test.TestCase
}

func (r *JavaRunner) Run() error {
	inreader, outwriter, err := r.GetIOStreams();
	defer CloseFiles(inreader, outwriter);
	
	if err != nil {
		return err;
	}
	
	exec_ := path.Join(r.Execpath, r.Execfile);
        cmd := exec.Command("java",exec_);
        
        cmd.Stdin = inreader;
        cmd.Stdout = outwriter;
	done := make(chan error, 1);

	//Start execution and wait
	timeout := time.After(time.Duration(r.Maxtime) * time.Millisecond);
	cmd.Start();
	go func() {
		done <- cmd.Wait();
	}();

	select{
	case <- timeout:
		return r.HandleTLE(cmd);

	case err := <- done:
		return err;
	}
}

func (r *JavaRunner) HandleTLE(cmd *exec.Cmd) error{
	return cmd.Process.Kill();
}
