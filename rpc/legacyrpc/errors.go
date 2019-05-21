// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"fmt"

	"github.com/bitum-project/bitumd/bitumjson"
	"github.com/bitum-project/bitumwallet/errors"
)

func convertError(err error) *bitumjson.RPCError {
	if err, ok := err.(*bitumjson.RPCError); ok {
		return err
	}

	code := bitumjson.ErrRPCWallet
	if err, ok := err.(*errors.Error); ok {
		switch err.Kind {
		case errors.Bug:
			code = bitumjson.ErrRPCInternal.Code
		case errors.Encoding:
			code = bitumjson.ErrRPCInvalidParameter
		case errors.Locked:
			code = bitumjson.ErrRPCWalletUnlockNeeded
		case errors.Passphrase:
			code = bitumjson.ErrRPCWalletPassphraseIncorrect
		case errors.NoPeers:
			code = bitumjson.ErrRPCClientNotConnected
		case errors.InsufficientBalance:
			code = bitumjson.ErrRPCWalletInsufficientFunds
		}
	}
	return &bitumjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcError(code bitumjson.RPCErrorCode, err error) *bitumjson.RPCError {
	return &bitumjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcErrorf(code bitumjson.RPCErrorCode, format string, args ...interface{}) *bitumjson.RPCError {
	return &bitumjson.RPCError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// Errors variables that are defined once here to avoid duplication.
var (
	errUnloadedWallet = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCWallet,
		Message: "request requires a wallet but wallet has not loaded yet",
	}

	errRPCClientNotConnected = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCClientNotConnected,
		Message: "disconnected from consensus RPC",
	}

	errNoNetwork = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCClientNotConnected,
		Message: "disconnected from network",
	}

	errAccountNotFound = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCWalletInvalidAccountName,
		Message: "account not found",
	}

	errAddressNotInWallet = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	errNotImportedAccount = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	errNeedPositiveAmount = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCInvalidParameter,
		Message: "amount must be positive",
	}

	errWalletUnlockNeeded = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCWalletUnlockNeeded,
		Message: "enter the wallet passphrase with walletpassphrase first",
	}

	errReservedAccountName = &bitumjson.RPCError{
		Code:    bitumjson.ErrRPCInvalidParameter,
		Message: "account name is reserved by RPC server",
	}
)
