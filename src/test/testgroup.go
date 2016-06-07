package test;

import (
	"errors";
	"strings";
	"os";
	"path/filepath";
	"path";
)

/*

*/

type TestGroup struct {
	
	TestId, RunId, Lang string

	Codefile string

	Maxtime int64

}

func (t *TestGroup) GenerateTestCases () ([]*TestCase, error) {

	inputpath := path.Join(os.Getenv("OC_INPUTS"),t.TestId);
	outputpath := path.Join(os.Getenv("OC_OUTPUTS"), t.RunId);
	testpath := path.Join(os.Getenv("OC_TEST"), t.TestId);

	if _, err := os.Stat(inputpath); os.IsNotExist(err) {
		return nil, errors.New("Inputs not found");
	}

	if _, err := os.Stat(testpath); os.IsNotExist(err) {
		return nil, errors.New("Test outputs not found");
	}

	infiles, testfiles := []string{},[]string{};

	var addinput bool; //Toggle for adding input files or test files
	var addfiles filepath.WalkFunc = func (root string, info os.FileInfo, ferr error) error {
		name := info.Name();

		if ferr != nil || info.IsDir() {
			return nil; //Ignore if error
		}

		if strings.HasSuffix(name, ".txt"){
			if addinput {
				infiles = append(infiles, name);	
			}else{
				testfiles = append(testfiles, name);
			}
		}
		return nil;
	}

	addinput = true;
	filepath.Walk(inputpath, addfiles);
	addinput = false;
	filepath.Walk(testpath, addfiles);

	if len(infiles) != len(testfiles) {
		return nil, errors.New("Should contain equal number of inputs and tests.");
	}

	testcases := make([]*TestCase, len(infiles));
	for i, f := range infiles {
		testcases[i] = &TestCase{
			Inputfile: path.Join(inputpath, f),
			Outputfile: path.Join(outputpath, f),
			Testfile: path.Join(testpath, testfiles[i]),
			Maxtime: t.Maxtime,
		};
	}
	
	return testcases, nil;
}
