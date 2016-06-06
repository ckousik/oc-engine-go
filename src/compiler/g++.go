package compiler;

import (
	"fmt";
	"os/exec";
	"path";
	"test";
	"strings";
	"errors";
)

type GPP struct {
	*test.TestGroup
}

func (c *GPP) Compile () error {
	err := c.setExecfile();

	if err != nil {
		fmt.Println(err);
		return err;
	}
	
	in := path.Join(c.Codepath, c.Codefile);
	out := path.Join(c.Execpath, c.Execfile);
	
	cmd := exec.Command ("g++",in,"-o",out);

	done := make(chan error,1);

	cmd.Start();
	go func() {
		done <- cmd.Wait();
	}();
	return <- done;
}

func (c *GPP) setExecfile () error {
	
	if !strings.HasSuffix(c.Codefile,".cpp") {
		return errors.New("Invalid file extension for C++ compiler");
	}
	
	c.Execfile = strings.TrimSuffix(c.Codefile,".cpp");
	return nil;
}
