package worker;

import(
	"testing";
	"path";
	"os";
	"github.com/taskcluster/slugid-go/slugid"
)

func TestWorker (t *testing.T) {
	inputchan := make(chan InputPayload);
	resultchan := make(chan ResultPayload);

	resultchan = Deploy(inputchan);
	inputchan <- InputPayload{
		Lang: "cpp",
		Codefile: path.Join(os.Getenv("GOPATH"),"tests","count.cpp"),
		TestId: "count",
		RunId: slugid.Nice(),
	}

	result := <- resultchan;
	t.Logf("Result Codes: ");
	for _, c := range result.Results {
		t.Logf("%d ",int(c));
	}
}
