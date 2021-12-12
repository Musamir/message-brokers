package implementation

import "fmt"

var (
	//SomethingWentWrong ...
	SomethingWentWrong = fmt.Errorf("Something went wrong ")
	//FailedToDial ...
	FailedToDial = fmt.Errorf("Failed to dial with the broker ")
)
