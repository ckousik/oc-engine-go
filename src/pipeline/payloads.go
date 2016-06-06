package pipeline;

import (
	"os";
	"errors";
	"path/filepath";
	"path";
	"test";
	"strings";
)

type Payload struct {

	Lang string

	Codepath, Codefile string

	Name, Id string

	TestId string
}

func (p *Payload) VerifyPayload () error {

	if _, err := os.Stat(path.Join(p.Codepath,p.Codefile)); os.IsNotExist(err) {
		return errors.New("Code not found at " + path.Join(p.Codepath, p.Codefile));
	}

	if _, err := os.Stat(path.Join(os.Getenv("OC_INPUT_PATH"), p.TestId)); os.IsNotExist(err) {
		return errors.New("No input files found for test Id: " +  p.TestId);
	}

	if _, err := os.Stat(path.Join(os.Getenv("OC_TEST_PATH"), p.TestId)); os.IsNotExist(err) {
		return errors.New("No test files found for test Id: " +  p.TestId);
	}
	return nil;
}

func (p *Payload) GenerateTestGroup () (*test.TestGroup, error) {

	if err := p.VerifyPayload(); err != nil {
		return nil, err;
	}
	
/*
* Get input and test files from directory
*/
	var addInput bool;

	Inputfiles := []string{};
	Inputpath := path.Join(os.Getenv("OC_INPUT_PATH"), p.TestId);

	Testfiles := []string{};
	Testpath := path.Join(os.Getenv("OC_TEST_PATH"), p.TestId);

	add := func(root string, finfo os.FileInfo, ferr error) error {

		if finfo.IsDir() {
			return nil;
		}
		
		name := finfo.Name();
		if strings.HasSuffix(name, ".txt") ||
			strings.HasSuffix(name, ".in") ||
			strings.HasSuffix(name,".test") {
			if addInput {
				Inputfiles = append(Inputfiles, name);				
			}else {
				Testfiles = append(Testfiles, name);
			}

		}
		return nil;
	}

	addInput = true;
	filepath.Walk(Inputpath, add);
	addInput = false;
	filepath.Walk(Testpath, add);

	//Create output directory
	Outputpath := path.Join(os.Getenv("OC_OUTPUT_PATH"), p.Id, "out");
	if err := os.MkdirAll(Outputpath, 0777); err!= nil {
		return nil, err;
	} 

	//Create directory to store executables
	Execpath := path.Join(os.Getenv("OC_USER_EXEC_PATH"), p.Id, "exec");
	if err := os.MkdirAll(Execpath, 0777); err!= nil {
		return nil, err;
	}
	
	return &test.TestGroup {

		Name: p.Name,
		Id: p.Id,
		
		Codepath: p.Codepath,
		Codefile: p.Codefile,
		
		Inputpath: Inputpath,
		Inputfiles: Inputfiles,

		Testpath: Testpath,
		Testfiles: Testfiles,

		Outputpath: Outputpath,

		Execpath: Execpath,
	}, nil;
}
