package compiler;

import (	
	"os/exec";
	"path";
	"test";
	"errors";
	"strings";
)

type GCC struct {
	*test.TestGroup;
};

func (c *GCC) Compile () error {
	in := path.Join(c.Codepath, c.Codefile);
	err := c.setExecfile();
	if err != nil {
		return err;
	}
	out := path.Join(c.Execpath, c.Execfile);

	cmd := exec.Command("gcc",in,"-o",out);

	done := make(chan error, 1);

	go func(){
		done <- cmd.Wait();
	}();

	return <- done;
}

func (c *GCC) setExecfile () error {
	if strings.HasSuffix(c.Codefile, ".c") {

		c.Execfile = strings.TrimSuffix(c.Codefile, ".c");

	}else{

		return errors.New("Invalid file extension");
	}

	return nil;
}
