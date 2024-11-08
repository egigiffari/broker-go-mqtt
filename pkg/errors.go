package pkg

import "fmt"

var (
	ErrRequireHost                 = fmt.Errorf("broker require host")
	ErrInvalidPort                 = fmt.Errorf("borker invalid port")
	ErrRequireClientId             = fmt.Errorf("broker require client id")
	ErrRequireUsername             = fmt.Errorf("broker require username")
	ErrRequirePassword             = fmt.Errorf("broker require password")
	ErrInvalidReconnectInterval    = fmt.Errorf("broker reconnect must be at least 5 second")
	ErrInvalidMaxReconnectInterval = fmt.Errorf("broker reconnect must be at least 1 hour")
)
