/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package provenancefakes

import (
	"sync"

	"github.com/in-toto/in-toto-golang/in_toto"

	"sigs.k8s.io/bom/pkg/provenance"
)

type FakeStatementImplementation struct {
	AddSubjectStub        func(*provenance.Statement, string, in_toto.DigestSet)
	addSubjectMutex       sync.RWMutex
	addSubjectArgsForCall []struct {
		arg1 *provenance.Statement
		arg2 string
		arg3 in_toto.DigestSet
	}
	ClonePredicateStub        func(*provenance.Statement, string) error
	clonePredicateMutex       sync.RWMutex
	clonePredicateArgsForCall []struct {
		arg1 *provenance.Statement
		arg2 string
	}
	clonePredicateReturns struct {
		result1 error
	}
	clonePredicateReturnsOnCall map[int]struct {
		result1 error
	}
	ReadSubjectsFromDirStub        func(*provenance.Statement, string) error
	readSubjectsFromDirMutex       sync.RWMutex
	readSubjectsFromDirArgsForCall []struct {
		arg1 *provenance.Statement
		arg2 string
	}
	readSubjectsFromDirReturns struct {
		result1 error
	}
	readSubjectsFromDirReturnsOnCall map[int]struct {
		result1 error
	}
	SubjectFromFileStub        func(string) (in_toto.Subject, error)
	subjectFromFileMutex       sync.RWMutex
	subjectFromFileArgsForCall []struct {
		arg1 string
	}
	subjectFromFileReturns struct {
		result1 in_toto.Subject
		result2 error
	}
	subjectFromFileReturnsOnCall map[int]struct {
		result1 in_toto.Subject
		result2 error
	}
	ToJSONStub        func(*provenance.Statement) ([]byte, error)
	toJSONMutex       sync.RWMutex
	toJSONArgsForCall []struct {
		arg1 *provenance.Statement
	}
	toJSONReturns struct {
		result1 []byte
		result2 error
	}
	toJSONReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	VerifySubjectsStub        func(string, *[]in_toto.Subject) error
	verifySubjectsMutex       sync.RWMutex
	verifySubjectsArgsForCall []struct {
		arg1 string
		arg2 *[]in_toto.Subject
	}
	verifySubjectsReturns struct {
		result1 error
	}
	verifySubjectsReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func(*provenance.Statement, string) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 *provenance.Statement
		arg2 string
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatementImplementation) AddSubject(arg1 *provenance.Statement, arg2 string, arg3 in_toto.DigestSet) {
	fake.addSubjectMutex.Lock()
	fake.addSubjectArgsForCall = append(fake.addSubjectArgsForCall, struct {
		arg1 *provenance.Statement
		arg2 string
		arg3 in_toto.DigestSet
	}{arg1, arg2, arg3})
	stub := fake.AddSubjectStub
	fake.recordInvocation("AddSubject", []interface{}{arg1, arg2, arg3})
	fake.addSubjectMutex.Unlock()
	if stub != nil {
		fake.AddSubjectStub(arg1, arg2, arg3)
	}
}

func (fake *FakeStatementImplementation) AddSubjectCallCount() int {
	fake.addSubjectMutex.RLock()
	defer fake.addSubjectMutex.RUnlock()
	return len(fake.addSubjectArgsForCall)
}

func (fake *FakeStatementImplementation) AddSubjectCalls(stub func(*provenance.Statement, string, in_toto.DigestSet)) {
	fake.addSubjectMutex.Lock()
	defer fake.addSubjectMutex.Unlock()
	fake.AddSubjectStub = stub
}

func (fake *FakeStatementImplementation) AddSubjectArgsForCall(i int) (*provenance.Statement, string, in_toto.DigestSet) {
	fake.addSubjectMutex.RLock()
	defer fake.addSubjectMutex.RUnlock()
	argsForCall := fake.addSubjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStatementImplementation) ClonePredicate(arg1 *provenance.Statement, arg2 string) error {
	fake.clonePredicateMutex.Lock()
	ret, specificReturn := fake.clonePredicateReturnsOnCall[len(fake.clonePredicateArgsForCall)]
	fake.clonePredicateArgsForCall = append(fake.clonePredicateArgsForCall, struct {
		arg1 *provenance.Statement
		arg2 string
	}{arg1, arg2})
	stub := fake.ClonePredicateStub
	fakeReturns := fake.clonePredicateReturns
	fake.recordInvocation("ClonePredicate", []interface{}{arg1, arg2})
	fake.clonePredicateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStatementImplementation) ClonePredicateCallCount() int {
	fake.clonePredicateMutex.RLock()
	defer fake.clonePredicateMutex.RUnlock()
	return len(fake.clonePredicateArgsForCall)
}

func (fake *FakeStatementImplementation) ClonePredicateCalls(stub func(*provenance.Statement, string) error) {
	fake.clonePredicateMutex.Lock()
	defer fake.clonePredicateMutex.Unlock()
	fake.ClonePredicateStub = stub
}

func (fake *FakeStatementImplementation) ClonePredicateArgsForCall(i int) (*provenance.Statement, string) {
	fake.clonePredicateMutex.RLock()
	defer fake.clonePredicateMutex.RUnlock()
	argsForCall := fake.clonePredicateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatementImplementation) ClonePredicateReturns(result1 error) {
	fake.clonePredicateMutex.Lock()
	defer fake.clonePredicateMutex.Unlock()
	fake.ClonePredicateStub = nil
	fake.clonePredicateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) ClonePredicateReturnsOnCall(i int, result1 error) {
	fake.clonePredicateMutex.Lock()
	defer fake.clonePredicateMutex.Unlock()
	fake.ClonePredicateStub = nil
	if fake.clonePredicateReturnsOnCall == nil {
		fake.clonePredicateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clonePredicateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDir(arg1 *provenance.Statement, arg2 string) error {
	fake.readSubjectsFromDirMutex.Lock()
	ret, specificReturn := fake.readSubjectsFromDirReturnsOnCall[len(fake.readSubjectsFromDirArgsForCall)]
	fake.readSubjectsFromDirArgsForCall = append(fake.readSubjectsFromDirArgsForCall, struct {
		arg1 *provenance.Statement
		arg2 string
	}{arg1, arg2})
	stub := fake.ReadSubjectsFromDirStub
	fakeReturns := fake.readSubjectsFromDirReturns
	fake.recordInvocation("ReadSubjectsFromDir", []interface{}{arg1, arg2})
	fake.readSubjectsFromDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDirCallCount() int {
	fake.readSubjectsFromDirMutex.RLock()
	defer fake.readSubjectsFromDirMutex.RUnlock()
	return len(fake.readSubjectsFromDirArgsForCall)
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDirCalls(stub func(*provenance.Statement, string) error) {
	fake.readSubjectsFromDirMutex.Lock()
	defer fake.readSubjectsFromDirMutex.Unlock()
	fake.ReadSubjectsFromDirStub = stub
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDirArgsForCall(i int) (*provenance.Statement, string) {
	fake.readSubjectsFromDirMutex.RLock()
	defer fake.readSubjectsFromDirMutex.RUnlock()
	argsForCall := fake.readSubjectsFromDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDirReturns(result1 error) {
	fake.readSubjectsFromDirMutex.Lock()
	defer fake.readSubjectsFromDirMutex.Unlock()
	fake.ReadSubjectsFromDirStub = nil
	fake.readSubjectsFromDirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) ReadSubjectsFromDirReturnsOnCall(i int, result1 error) {
	fake.readSubjectsFromDirMutex.Lock()
	defer fake.readSubjectsFromDirMutex.Unlock()
	fake.ReadSubjectsFromDirStub = nil
	if fake.readSubjectsFromDirReturnsOnCall == nil {
		fake.readSubjectsFromDirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.readSubjectsFromDirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) SubjectFromFile(arg1 string) (in_toto.Subject, error) {
	fake.subjectFromFileMutex.Lock()
	ret, specificReturn := fake.subjectFromFileReturnsOnCall[len(fake.subjectFromFileArgsForCall)]
	fake.subjectFromFileArgsForCall = append(fake.subjectFromFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SubjectFromFileStub
	fakeReturns := fake.subjectFromFileReturns
	fake.recordInvocation("SubjectFromFile", []interface{}{arg1})
	fake.subjectFromFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatementImplementation) SubjectFromFileCallCount() int {
	fake.subjectFromFileMutex.RLock()
	defer fake.subjectFromFileMutex.RUnlock()
	return len(fake.subjectFromFileArgsForCall)
}

func (fake *FakeStatementImplementation) SubjectFromFileCalls(stub func(string) (in_toto.Subject, error)) {
	fake.subjectFromFileMutex.Lock()
	defer fake.subjectFromFileMutex.Unlock()
	fake.SubjectFromFileStub = stub
}

func (fake *FakeStatementImplementation) SubjectFromFileArgsForCall(i int) string {
	fake.subjectFromFileMutex.RLock()
	defer fake.subjectFromFileMutex.RUnlock()
	argsForCall := fake.subjectFromFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStatementImplementation) SubjectFromFileReturns(result1 in_toto.Subject, result2 error) {
	fake.subjectFromFileMutex.Lock()
	defer fake.subjectFromFileMutex.Unlock()
	fake.SubjectFromFileStub = nil
	fake.subjectFromFileReturns = struct {
		result1 in_toto.Subject
		result2 error
	}{result1, result2}
}

func (fake *FakeStatementImplementation) SubjectFromFileReturnsOnCall(i int, result1 in_toto.Subject, result2 error) {
	fake.subjectFromFileMutex.Lock()
	defer fake.subjectFromFileMutex.Unlock()
	fake.SubjectFromFileStub = nil
	if fake.subjectFromFileReturnsOnCall == nil {
		fake.subjectFromFileReturnsOnCall = make(map[int]struct {
			result1 in_toto.Subject
			result2 error
		})
	}
	fake.subjectFromFileReturnsOnCall[i] = struct {
		result1 in_toto.Subject
		result2 error
	}{result1, result2}
}

func (fake *FakeStatementImplementation) ToJSON(arg1 *provenance.Statement) ([]byte, error) {
	fake.toJSONMutex.Lock()
	ret, specificReturn := fake.toJSONReturnsOnCall[len(fake.toJSONArgsForCall)]
	fake.toJSONArgsForCall = append(fake.toJSONArgsForCall, struct {
		arg1 *provenance.Statement
	}{arg1})
	stub := fake.ToJSONStub
	fakeReturns := fake.toJSONReturns
	fake.recordInvocation("ToJSON", []interface{}{arg1})
	fake.toJSONMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatementImplementation) ToJSONCallCount() int {
	fake.toJSONMutex.RLock()
	defer fake.toJSONMutex.RUnlock()
	return len(fake.toJSONArgsForCall)
}

func (fake *FakeStatementImplementation) ToJSONCalls(stub func(*provenance.Statement) ([]byte, error)) {
	fake.toJSONMutex.Lock()
	defer fake.toJSONMutex.Unlock()
	fake.ToJSONStub = stub
}

func (fake *FakeStatementImplementation) ToJSONArgsForCall(i int) *provenance.Statement {
	fake.toJSONMutex.RLock()
	defer fake.toJSONMutex.RUnlock()
	argsForCall := fake.toJSONArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStatementImplementation) ToJSONReturns(result1 []byte, result2 error) {
	fake.toJSONMutex.Lock()
	defer fake.toJSONMutex.Unlock()
	fake.ToJSONStub = nil
	fake.toJSONReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStatementImplementation) ToJSONReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.toJSONMutex.Lock()
	defer fake.toJSONMutex.Unlock()
	fake.ToJSONStub = nil
	if fake.toJSONReturnsOnCall == nil {
		fake.toJSONReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.toJSONReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStatementImplementation) VerifySubjects(arg1 string, arg2 *[]in_toto.Subject) error {
	fake.verifySubjectsMutex.Lock()
	ret, specificReturn := fake.verifySubjectsReturnsOnCall[len(fake.verifySubjectsArgsForCall)]
	fake.verifySubjectsArgsForCall = append(fake.verifySubjectsArgsForCall, struct {
		arg1 string
		arg2 *[]in_toto.Subject
	}{arg1, arg2})
	stub := fake.VerifySubjectsStub
	fakeReturns := fake.verifySubjectsReturns
	fake.recordInvocation("VerifySubjects", []interface{}{arg1, arg2})
	fake.verifySubjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStatementImplementation) VerifySubjectsCallCount() int {
	fake.verifySubjectsMutex.RLock()
	defer fake.verifySubjectsMutex.RUnlock()
	return len(fake.verifySubjectsArgsForCall)
}

func (fake *FakeStatementImplementation) VerifySubjectsCalls(stub func(string, *[]in_toto.Subject) error) {
	fake.verifySubjectsMutex.Lock()
	defer fake.verifySubjectsMutex.Unlock()
	fake.VerifySubjectsStub = stub
}

func (fake *FakeStatementImplementation) VerifySubjectsArgsForCall(i int) (string, *[]in_toto.Subject) {
	fake.verifySubjectsMutex.RLock()
	defer fake.verifySubjectsMutex.RUnlock()
	argsForCall := fake.verifySubjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatementImplementation) VerifySubjectsReturns(result1 error) {
	fake.verifySubjectsMutex.Lock()
	defer fake.verifySubjectsMutex.Unlock()
	fake.VerifySubjectsStub = nil
	fake.verifySubjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) VerifySubjectsReturnsOnCall(i int, result1 error) {
	fake.verifySubjectsMutex.Lock()
	defer fake.verifySubjectsMutex.Unlock()
	fake.VerifySubjectsStub = nil
	if fake.verifySubjectsReturnsOnCall == nil {
		fake.verifySubjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifySubjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) Write(arg1 *provenance.Statement, arg2 string) error {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 *provenance.Statement
		arg2 string
	}{arg1, arg2})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStatementImplementation) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeStatementImplementation) WriteCalls(stub func(*provenance.Statement, string) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeStatementImplementation) WriteArgsForCall(i int) (*provenance.Statement, string) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatementImplementation) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatementImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addSubjectMutex.RLock()
	defer fake.addSubjectMutex.RUnlock()
	fake.clonePredicateMutex.RLock()
	defer fake.clonePredicateMutex.RUnlock()
	fake.readSubjectsFromDirMutex.RLock()
	defer fake.readSubjectsFromDirMutex.RUnlock()
	fake.subjectFromFileMutex.RLock()
	defer fake.subjectFromFileMutex.RUnlock()
	fake.toJSONMutex.RLock()
	defer fake.toJSONMutex.RUnlock()
	fake.verifySubjectsMutex.RLock()
	defer fake.verifySubjectsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatementImplementation) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ provenance.StatementImplementation = new(FakeStatementImplementation)
