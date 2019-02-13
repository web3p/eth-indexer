// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import big "math/big"

import common "github.com/ethereum/go-ethereum/common"
import context "context"
import ethereum "github.com/ethereum/go-ethereum"
import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"
import types "github.com/ethereum/go-ethereum/core/types"

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// BalanceOf provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) BalanceOf(_a0 context.Context, _a1 common.Hash, _a2 map[common.Address]map[common.Address]*big.Int) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, map[common.Address]map[common.Address]*big.Int) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchBalanceAt provides a mock function with given fields: ctx, accounts, blockHash
func (_m *EthClient) BatchBalanceAt(ctx context.Context, accounts []common.Address, blockHash common.Hash) ([]*big.Int, error) {
	ret := _m.Called(ctx, accounts, blockHash)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address, common.Hash) []*big.Int); ok {
		r0 = rf(ctx, accounts, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []common.Address, common.Hash) error); ok {
		r1 = rf(ctx, accounts, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchCallContract provides a mock function with given fields: ctx, msgs, blockHash
func (_m *EthClient) BatchCallContract(ctx context.Context, msgs []*ethereum.CallMsg, blockHash common.Hash) ([][]byte, error) {
	ret := _m.Called(ctx, msgs, blockHash)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, []*ethereum.CallMsg, common.Hash) [][]byte); ok {
		r0 = rf(ctx, msgs, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*ethereum.CallMsg, common.Hash) error); ok {
		r1 = rf(ctx, msgs, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *EthClient) Close() {
	_m.Called()
}

// GetBlockReceipts provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetBlockReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	ret := _m.Called(ctx, hash)

	var r0 types.Receipts
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) types.Receipts); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Receipts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetERC20 provides a mock function with given fields: ctx, addr
func (_m *EthClient) GetERC20(ctx context.Context, addr common.Address) (*model.ERC20, error) {
	ret := _m.Called(ctx, addr)

	var r0 *model.ERC20
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) *model.ERC20); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ERC20)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalDifficulty provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetTotalDifficulty(ctx context.Context, hash common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, hash)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *big.Int); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransferLogs provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetTransferLogs(ctx context.Context, hash common.Hash) ([]*types.TransferLog, error) {
	ret := _m.Called(ctx, hash)

	var r0 []*types.TransferLog
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []*types.TransferLog); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.TransferLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UncleByBlockHashAndPosition provides a mock function with given fields: ctx, hash, position
func (_m *EthClient) UncleByBlockHashAndPosition(ctx context.Context, hash common.Hash, position uint) (*types.Header, error) {
	ret := _m.Called(ctx, hash, position)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Header); ok {
		r0 = rf(ctx, hash, position)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, hash, position)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnclesByBlockHash provides a mock function with given fields: ctx, blockHash
func (_m *EthClient) UnclesByBlockHash(ctx context.Context, blockHash common.Hash) ([]*types.Header, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 []*types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []*types.Header); ok {
		r0 = rf(ctx, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
