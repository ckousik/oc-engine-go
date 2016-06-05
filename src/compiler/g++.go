package compiler;

import (
	"os/exec";
	"path";
)

type GPP struct {
	Codepath, Codefile string
	Execpath, Execfile string
}

func (c *GPP) Compile () error {
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
