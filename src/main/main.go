package main;

import (
	"fmt";
	"test";
	"runner";
	"os";
	"path";
	"github.com/taskcluster/slugid-go/slugid";
);

func TestCPP () {
	codefile := path.Join(os.Getenv("GOPATH"),"tests","count.cpp");
	runid := slugid.Nice();

	tg := &test.TestGroup {
		Codefile: codefile,
		RunId: runid,
		TestId: "count",
		Lang: "cpp",
		Maxtime: 10,
	};

	r := runner.CPP{};
	status := r.Start(tg);
	var st runner.StatusCode;
	st = <- status;
	for st != runner.ExecutionCompleted  {
		fmt.Println(st);
		st = <- status;
	}
}

func TestC () {
	codefile := path.Join(os.Getenv("GOPATH"),"tests","count.c");
	runid := slugid.Nice();

	tg := &test.TestGroup {
		Codefile: codefile,
		RunId: runid,
		TestId: "count",
		Lang: "c",
		Maxtime: 10,
	};

	r := runner.C{};
	status := r.Start(tg);
	var st runner.StatusCode;
	st = <- status;
	for st != runner.ExecutionCompleted  {
		fmt.Println(st);
		st = <- status;
	}
}


func main() {
	TestC();
}
