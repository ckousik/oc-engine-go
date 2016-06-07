package runner;

import (
	"test";
	"testing";
	"os";
	"path";
	"github.com/taskcluster/slugid-go/slugid";
);

func TestC (t *testing.T) {
	codefile := path.Join(os.Getenv("GOPATH"),"tests","count.c");
	runid := slugid.Nice();

	tg := &test.TestGroup {
		Codefile: codefile,
		RunId: runid,
		TestId: "count",
		Lang: "c",
		Maxtime: 10,
	};

	r := C{};
	status := r.Start(tg);
	var st StatusCode;
	st = <- status;
	for st != ExecutionCompleted  {
		t.Log(st);
		st = <- status;
	}
}
