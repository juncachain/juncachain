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

package jrc21issuer

import (
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts/jrc21issuer/contract"
	"math/big"
)

type JRC21Issuer struct {
	*contract.JRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewJRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*JRC21Issuer, error) {
	jrc21issuer, err := contract.NewJRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &JRC21Issuer{
		&contract.JRC21IssuerSession{
			Contract:     jrc21issuer,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployJRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, gasPerTx *big.Int) (common.Address, *JRC21Issuer, error) {
	jrc21IssuerAddr, _, _, err := contract.DeployJRC21Issuer(transactOpts, contractBackend, gasPerTx)
	if err != nil {
		return jrc21IssuerAddr, nil, err
	}

	jrc21Issuer, err := NewJRC21Issuer(transactOpts, jrc21IssuerAddr, contractBackend)
	if err != nil {
		return jrc21IssuerAddr, nil, err
	}

	return jrc21IssuerAddr, jrc21Issuer, nil
}
