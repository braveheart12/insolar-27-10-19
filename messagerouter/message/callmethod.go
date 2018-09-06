package message

import "github.com/insolar/insolar/core"

type CallMethodMessage struct {
	BaseMessage
	Request   core.RecordRef
	Method    string
	Arguments core.Arguments
}
