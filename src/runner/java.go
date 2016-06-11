package runner;
/*
import(
	"task";
	"os";
	"os/exec";
)

type Java struct{};

func (j Java) Start(t *task.TestGroup) chan StatusCode {
	status := make(chan StatusCode);
	//Make directory to store outputs
	os.Mkdir(path.Join(os.Getenv("OC_OUTPUTS"),t.RunId ),0777);
	go func(){
		//Compile
		compile_input := t.Codefile;
		compile_output := path.Join(os.Getenv("OC_EXEC"), t.RunId);
		compile_command := exec.Command("javac",compile_input,"-d",compile_output);
	}
}
*/
