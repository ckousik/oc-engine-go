package compiler;

import "test";

type Compiler interface {
	
	Compile () error
	SetTestGroup (*test.TestGroup)
}
