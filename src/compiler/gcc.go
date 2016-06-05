package compiler;

import (	
	"os/exec";
	"path";
)

type GCC struct {
	Codepath, Codefile, Execpath, Execfile string
};

func (c *GCC) Compile () error {
	in := path.Join(c.Codepath, c.Codefile);
	out := path.Join(c.Execpath, c.Execfile);

	cmd := exec.Command("gcc",in,"-o",out);

	done := make(chan error, 1);

	go func(){
		done <- cmd.Wait();
	}();

	return <- done;
}

