// Code generated by mockery v1.0.0. DO NOT EDIT.

package services

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/lbryio/rosetta-sdk-go/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetPeers provides a mock function with given fields: _a0
func (_m *Client) GetPeers(_a0 context.Context) ([]*types.Peer, error) {
	ret := _m.Called(_a0)

	var r0 []*types.Peer
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Peer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RawMempool provides a mock function with given fields: _a0
func (_m *Client) RawMempool(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRawTransaction provides a mock function with given fields: _a0, _a1
func (_m *Client) SendRawTransaction(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestedFeeRate provides a mock function with given fields: _a0, _a1
func (_m *Client) SuggestedFeeRate(_a0 context.Context, _a1 int64) (float64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, int64) float64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
