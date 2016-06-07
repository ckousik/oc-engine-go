package worker;

import (
	"runner";
	"compiler";
	"errors";
)

type Worker struct {
	c compiler.Compiler
	r runner.Runner
}

func (w *Worker) Start(p *Payload) (chan bool) {
	statusChan := make (chan bool);
	go func() {
		testGroup, err := p.GenerateTestGroup();
		if err != nil {
			statusChan <- false;
			return;
		}
		
		w.c.SetTestGroup(testGroup);
		tests, err := w.c.GetTestGroup().GenerateTestCases();
	}()
	return statusChan;
}

func NewWorker (lang string) (*Worker, error) {
	switch lang {
	case "GCC":
		return &Worker{
			c: compiler.GCC{nil},
			r: runner.ExecRunner{nil},
		}, nil;
	case "GPP":
		return &Worker{
			c: compiler.GPP{nil},
			r: runner.ExecRunner{nil},
		}, nil;
	case "Java":
		return &Worker{
			c: compiler.Java{nil},
			r: runner.JavaRunner{nil},
		}, nil;
	}

	return nil, errors.New("Unknown Worker Type");
}

