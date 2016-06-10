package worker;

import (
	"runner";
	"errors";
	"os";
	"path";
)
/*
Input Payload
--------------

Lang string : language used

Codefile string : path where the code is stored

TestId string : Id of test to be executed

RunId string : Id used to identify unique runs
*/

var(
	InvalidTestIdError error = errors.New("No test found with given testId");
	CodeNotFoundError error = errors.New("Code not found at the given path");
	InvalidLangError error = errors.New("Language not supported");
	CompilationError error = errors.New("Compile Error");
)

var (
	Langmap map[string]runner.Runner = map[string]runner.Runner {
		"c": runner.C{},
		"cpp": runner.CPP{},
		"py2": runner.Py2{},
		"py3": runner.Py3{},
	}
)

type InputPayload struct {

	Lang string

	Codefile string

	TestId string

	RunId string
}

func (inpayload InputPayload) Verify () error {

	if Langmap[inpayload.Lang] == nil {
		return InvalidLangError;
	}

	if _, err := os.Stat(path.Join(os.Getenv("OC_TEST"),inpayload.TestId)); os.IsNotExist(err) {
		return InvalidTestIdError;
	}

	if _, err := os.Stat(inpayload.Codefile); os.IsNotExist(err) {
		return CodeNotFoundError;
	}

	return nil;

}

/*
Result Payload
--------------

RunId string : Identifies unique runs

Results : Slice of statuscodes
*/

type ResultPayload struct {

	RunId string

	Results []runner.StatusCode

	Err error
	
}
