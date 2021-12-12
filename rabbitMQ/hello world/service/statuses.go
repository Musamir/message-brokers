package service

import "fmt"

var (
	//SomethingWentWrong ...
	SomethingWentWrong = fmt.Errorf("Something went wrong ")
	//FailedInit ...
	FailedInit = fmt.Errorf("Failed to initialize broker ")
	//EnvironmentNotSet ...
	EnvironmentNotSet = fmt.Errorf("Environment is not set ")
)
