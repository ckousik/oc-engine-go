package compiler;

import (
	"fmt";
	"os";
	"os/exec";
	"path";
	"testcase";
	"strings";
	"errors";
)

type GPP struct {
	T *testcase.TestGroup
}

func (c *GPP) Compile () error {
	t := c.T;
	
	err := c.setExecfile();

	if err != nil {
		fmt.Println(err);
		return err;
	}
	
	in := path.Join(t.Codepath, t.Codefile);
	
	out := path.Join(t.Execpath, t.Execfile);
	
	fmt.Println("Compiling: input:", in);
	cmd := exec.Command ("g++",in,"-o",out);
	done := make(chan error,1);
	cmd.Start();
	cmd.Stderr = os.Stdout;
	go func() {
		done <- cmd.Wait();
	}();
	return <- done;
}

func (c *GPP) setExecfile () error {
	t := c.T;

	if !strings.HasSuffix(t.Codefile,".cpp") {
		return errors.New("Invalid file extension for C++ compiler");
	}
	
	t.Execfile = strings.TrimSuffix(t.Codefile,".cpp");
	return nil;
}
