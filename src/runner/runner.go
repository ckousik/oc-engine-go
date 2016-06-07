package runner;

import (
	"test";
	"os";
	"io/ioutil";
	"crypto/sha1";
	"bytes";
)

type StatusCode int

const (
	CompileError = 1
	RuntimeError = 2
	FileError = 3
	TestSuccess = 4
	TestFail = 5
	TestTLE = 6
	TestRunError = 7
	ExecutionCompleted = 8
)


type Runner interface {
	Start (* test.TestGroup) chan StatusCode
}

func CompareFiles (file1, file2 string) (bool, error) {
	f1, err := os.Open(file1);
	if err != nil {
		return false, err;
	}

	f2, err := os.Open(file2);
	if err != nil {
		return false, err;
	}

	hash := sha1.New();

	fc1, _ := ioutil.ReadAll(f1);
	fc2, _ := ioutil.ReadAll(f2);

	return bytes.Equal(hash.Sum(fc1), hash.Sum(fc2)), nil;
}
