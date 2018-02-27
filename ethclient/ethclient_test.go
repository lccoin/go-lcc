// Copyright 2016 The go-LCC Authors
// This file is part of the go-LCC library.
//
// The go-LCC library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-LCC library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-LCC library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/LCC/go-LCC"

// Verify that Client implements the LCC interfaces.
var (
	_ = LCC.ChainReader(&Client{})
	_ = LCC.TransactionReader(&Client{})
	_ = LCC.ChainStateReader(&Client{})
	_ = LCC.ChainSyncReader(&Client{})
	_ = LCC.ContractCaller(&Client{})
	_ = LCC.GasEstimator(&Client{})
	_ = LCC.GasPricer(&Client{})
	_ = LCC.LogFilterer(&Client{})
	_ = LCC.PendingStateReader(&Client{})
	// _ = LCC.PendingStateEventer(&Client{})
	_ = LCC.PendingContractCaller(&Client{})
)
