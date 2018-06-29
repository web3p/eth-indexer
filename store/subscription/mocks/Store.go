// Code generated by mockery v1.0.0
package mocks

import common "github.com/ethereum/go-ethereum/common"
import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Find provides a mock function with given fields: blockNumber
func (_m *Store) Find(blockNumber int64) ([]*model.Subscription, error) {
	ret := _m.Called(blockNumber)

	var r0 []*model.Subscription
	if rf, ok := ret.Get(0).(func(int64) []*model.Subscription); ok {
		r0 = rf(blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByAddresses provides a mock function with given fields: addrs
func (_m *Store) FindByAddresses(addrs [][]byte) ([]*model.Subscription, error) {
	ret := _m.Called(addrs)

	var r0 []*model.Subscription
	if rf, ok := ret.Get(0).(func([][]byte) []*model.Subscription); ok {
		r0 = rf(addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([][]byte) error); ok {
		r1 = rf(addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTotalBalance provides a mock function with given fields: blockNumber, token, group
func (_m *Store) FindTotalBalance(blockNumber int64, token common.Address, group int64) (*model.TotalBalance, error) {
	ret := _m.Called(blockNumber, token, group)

	var r0 *model.TotalBalance
	if rf, ok := ret.Get(0).(func(int64, common.Address, int64) *model.TotalBalance); ok {
		r0 = rf(blockNumber, token, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TotalBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, common.Address, int64) error); ok {
		r1 = rf(blockNumber, token, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *Store) Insert(data *model.Subscription) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Subscription) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertTotalBalance provides a mock function with given fields: data
func (_m *Store) InsertTotalBalance(data *model.TotalBalance) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.TotalBalance) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields: from, to
func (_m *Store) Reset(from int64, to int64) error {
	ret := _m.Called(from, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlockNumber provides a mock function with given fields: data
func (_m *Store) UpdateBlockNumber(data *model.Subscription) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Subscription) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
