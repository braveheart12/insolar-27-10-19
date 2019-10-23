//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package example

import "github.com/insolar/insolar/conveyor/smachine"

type ContractCallType uint8

const (
	_ ContractCallType = iota
	ContractCallMutable
	ContractCallImmutable
	ContractCallSaga
)

type CallResult interface {
}

type ContractRunnerService interface {
	ClassifyCall(code ArtifactBinary, method string) ContractCallType
	CallImmutableMethod(code ArtifactBinary, method string, state ArtifactBinary) CallResult
}

type ContractRunnerServiceAdapter struct {
	svc  ContractRunnerService
	exec smachine.ExecutionAdapter
}

func (a *ContractRunnerServiceAdapter) PrepareSync(ctx smachine.ExecutionContext, fn func(svc ContractRunnerService)) smachine.SyncCallRequester {
	return a.exec.PrepareSync(ctx, func(interface{}) smachine.AsyncResultFunc {
		fn(a.svc)
		return nil
	})
}

func (a *ContractRunnerServiceAdapter) PrepareAsync(ctx smachine.ExecutionContext, fn func(svc ContractRunnerService) smachine.AsyncResultFunc) smachine.AsyncCallRequester {
	return a.exec.PrepareAsync(ctx, func(interface{}) smachine.AsyncResultFunc {
		return fn(a.svc)
	})
}

type contractRunnerService struct{

}

func (c contractRunnerService) ClassifyCall(code ArtifactBinary, method string) ContractCallType {
	if code.GetCacheId()
}

func (c contractRunnerService) CallImmutableMethod(code ArtifactBinary, method string, state ArtifactBinary) CallResult {
	panic("implement me")
}
