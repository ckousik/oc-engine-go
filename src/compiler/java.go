package compiler;

import (
	"os/exec";
	"path";
	"strings";
	"errors";
)

type Java struct {
	Codepath, Codefile, Execpath, Execfile string;
}

func (c *Java) Compile () error {

	if !strings.HasSuffix(c.Codefile,".java") {
		return errors.New("File should be .java"); 
	}
	c.Execfile = strings.TrimSuffix(c.Codefile, ".java");
	
	in := path.Join(c.Codepath, c.Codefile);
	out := path.Clean(c.Execpath);
	cmd := exec.Command("javac",in,"-d",out);
	
	done := make(chan error, 1);

	cmd.Start();
	go func(){
		done <- cmd.Wait();
	}();

	return <- done;
	
}
