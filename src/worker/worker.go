
package worker;

import (

	"test";
	"runner";
)

func Deploy (inputchan chan InputPayload) chan ResultPayload{
	resultchan := make(chan ResultPayload);
	go func(){
		
		for inpayload := <- inputchan ; inpayload.RunId != ""; inpayload = <- inputchan {

			if err := inpayload.Verify(); err != nil {
				resultchan <- ResultPayload{
					RunId: inpayload.RunId,
					Err: err,
				};
				continue;
			}
			
			var tg *test.TestGroup = &test.TestGroup{
				RunId: inpayload.RunId,
				TestId: inpayload.TestId,
				Codefile: inpayload.Codefile,
				Maxtime: 200, //Default
			};

			
			run := Langmap[inpayload.Lang];
			statuschan := run.Start(tg);

			
			rescodes := []runner.StatusCode{};
			addresults := false;
			errored := false;
			status := <- statuschan;
			for ; status != runner.ExecutionCompleted; status = <- statuschan {
				if status == runner.CompileError{
					resultchan <- ResultPayload{
						RunId: inpayload.RunId,
						Err: CompilationError,
					};
					errored = true;
					break;	
				}

				if status == runner.ExecutionStarted{
					addresults = true;
				}

				if addresults {
					rescodes = append(rescodes, status);
				}
			}

			if !errored {
				resultchan <- ResultPayload{
					RunId: inpayload.RunId,
					Err: nil,
					Results: rescodes,
				}
			}
		}
	}();
	return resultchan;
}