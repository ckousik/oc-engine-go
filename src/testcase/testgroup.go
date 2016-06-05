package testcase;

import (
	"errors";
	"strings";
)

/*

*/

type TestGroup struct {
	
	Name, Id, Lang string

	Codefile, Codepath string
	Execfile, Execpath string

	Inputpath string
	Inputfiles []string

	Outputpath string
	Outputfiles []string

	Testpath string
	Testfiles []string

	Maxtime int64

}

func (t *TestGroup) GenerateTestCases () (error, []TestCase) {

	if len(t.Inputfiles) != len(t.Testfiles) {
		return errors.New("Number of input files and test files should be equal"), nil;
	}

	t.Outputfiles = make([]string, len(t.Inputfiles));
	tc := []TestCase{};
	
	for i, inp := range t.Inputfiles {

		var suffix string;

		if strings.HasSuffix(inp, ".txt"){
			suffix = ".txt";
		}else if strings.HasSuffix(inp, ".in"){
			suffix = ".in";
		}else {
			return errors.New("Input file has invalid extension"), nil;
		}
		
		t.Outputfiles[i] = strings.TrimSuffix(inp, suffix) + ".out";
		
		tc = append(tc, TestCase{

			Inputpath: t.Inputpath,
			Inputfile: inp,

			Outputpath: t.Outputpath,
			Outputfile: t.Outputfiles[i],

			Testpath: t.Testpath,
			Testfile: t.Testfiles[i],

			Execpath: t.Execpath,
			Execfile: t.Execfile,

			Maxtime: t.Maxtime,
		});
	}

	return nil, tc;
}
