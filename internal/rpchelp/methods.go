// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/bitum-project/bitumd/bitumjson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]bitumjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"addticket", nil},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*bitumjson.CreateMultiSigResult)(nil)}},
	{"createnewaccount", nil},
	{"dumpprivkey", returnsString},
	{"exportwatchingwallet", returnsString},
	{"generatevote", []interface{}{(*bitumjson.GenerateVoteResult)(nil)}},
	{"getaccountaddress", returnsString},
	{"getaccount", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", []interface{}{(*bitumjson.GetBalanceResult)(nil)}},
	{"getbestblockhash", returnsString},
	{"getbestblock", []interface{}{(*bitumjson.GetBestBlockResult)(nil)}},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*bitumjson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*bitumjson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"getstakeinfo", []interface{}{(*bitumjson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"gettickets", []interface{}{(*bitumjson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*bitumjson.GetTransactionResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"getvotechoices", []interface{}{(*bitumjson.GetVoteChoicesResult)(nil)}},
	{"getwalletfee", returnsNumber},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"listlockunspent", []interface{}{(*[]bitumjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]bitumjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]bitumjson.ListReceivedByAddressResult)(nil)}},
	{"listscripts", []interface{}{(*bitumjson.ListScriptsResult)(nil)}},
	{"listsinceblock", []interface{}{(*bitumjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*bitumjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"purchaseticket", returnsString},
	{"redeemmultisigout", []interface{}{(*bitumjson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*bitumjson.RedeemMultiSigOutResult)(nil)}},
	{"renameaccount", nil},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"setticketfee", returnsBool},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*bitumjson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*bitumjson.SignRawTransactionsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*bitumjson.StakePoolUserInfoResult)(nil)}},
	{"startautobuyer", nil},
	{"stopautobuyer", nil},
	{"sweepaccount", []interface{}{(*bitumjson.SweepAccountResult)(nil)}},
	{"ticketsforaddress", returnsBool},
	{"validateaddress", []interface{}{(*bitumjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]bitumjson.VersionResult)(nil)}},
	{"walletinfo", []interface{}{(*bitumjson.WalletInfoResult)(nil)}},
	{"walletislocked", returnsBool},
	{"walletlock", nil},
	{"walletpassphrasechange", nil},
	{"walletpassphrase", nil},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
