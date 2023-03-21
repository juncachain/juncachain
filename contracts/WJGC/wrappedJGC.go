// Copyright (c) 2018 Juncachain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package WJGC

import (
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts/WJGC/contract"
)

type WrappedJGC struct {
	*contract.WJGCSession
	contractBackend bind.ContractBackend
}

func NewWrappedJGC(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*WrappedJGC, error) {
	wjgc, err := contract.NewWJGC(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &WrappedJGC{
		&contract.WJGCSession{
			Contract:     wjgc,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployWrappedJGC(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *WrappedJGC, error) {
	contractAddr, _, _, err := contract.DeployWJGC(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	wjgc, err := NewWrappedJGC(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, wjgc, nil
}
