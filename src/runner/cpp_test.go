package runner;

import (
	"test";
	"testing";
	"os";
	"path";
	"github.com/taskcluster/slugid-go/slugid";
);

func TestCPP (t *testing.T) {
	codefile := path.Join(os.Getenv("GOPATH"),"tests","count.cpp");
	runid := slugid.Nice();

	tg := &test.TestGroup {
		Codefile: codefile,
		RunId: runid,
		TestId: "count",
		Lang: "cpp",
		Maxtime: 10,
	};

	r := CPP{};
	status := r.Start(tg);
	var st StatusCode;
	st = <- status;
	for st != ExecutionCompleted  {
		t.Log(st);
		st = <- status;
	}
}
