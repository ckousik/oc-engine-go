package compiler;

type Compiler interface {
	
	Compile() (error,string);
	
}
