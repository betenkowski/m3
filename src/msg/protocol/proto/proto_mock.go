// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3msg/protocol/proto (interfaces: Encoder,Decoder)

package proto

import (
	"io"

	"github.com/golang/mock/gomock"
)

// Mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *_MockEncoderRecorder
}

// Recorder for MockEncoder (not exported)
type _MockEncoderRecorder struct {
	mock *MockEncoder
}

func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &_MockEncoderRecorder{mock}
	return mock
}

func (_m *MockEncoder) EXPECT() *_MockEncoderRecorder {
	return _m.recorder
}

func (_m *MockEncoder) Bytes() []byte {
	ret := _m.ctrl.Call(_m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

func (_mr *_MockEncoderRecorder) Bytes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bytes")
}

func (_m *MockEncoder) Encode(_param0 Marshaler) error {
	ret := _m.ctrl.Call(_m, "Encode", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockEncoderRecorder) Encode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encode", arg0)
}

// Mock of Decoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *_MockDecoderRecorder
}

// Recorder for MockDecoder (not exported)
type _MockDecoderRecorder struct {
	mock *MockDecoder
}

func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &_MockDecoderRecorder{mock}
	return mock
}

func (_m *MockDecoder) EXPECT() *_MockDecoderRecorder {
	return _m.recorder
}

func (_m *MockDecoder) Decode(_param0 Unmarshaler) error {
	ret := _m.ctrl.Call(_m, "Decode", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDecoderRecorder) Decode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decode", arg0)
}

func (_m *MockDecoder) ResetReader(_param0 io.Reader) {
	_m.ctrl.Call(_m, "ResetReader", _param0)
}

func (_mr *_MockDecoderRecorder) ResetReader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetReader", arg0)
}