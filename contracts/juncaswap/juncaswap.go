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

package juncaswap

import (
	"github.com/juncachain/juncachain/accounts/abi/bind"
	"github.com/juncachain/juncachain/common"
	"github.com/juncachain/juncachain/contracts/juncaswap/contract"
)

type JuncaswapFactory struct {
	*contract.JuncaswapFactorySession
	contractBackend bind.ContractBackend
}

func NewJuncaswapFactory(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*JuncaswapFactory, error) {
	factory, err := contract.NewJuncaswapFactory(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &JuncaswapFactory{
		&contract.JuncaswapFactorySession{
			Contract:     factory,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployJuncaswapFactory(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, feeToSetter common.Address) (common.Address, *JuncaswapFactory, error) {
	contractAddr, _, _, err := contract.DeployJuncaswapFactory(transactOpts, contractBackend, feeToSetter)
	if err != nil {
		return contractAddr, nil, err
	}

	factory, err := NewJuncaswapFactory(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, factory, nil
}

type JuncaswapRouter struct {
	*contract.JuncaswapRouterSession
	contractBackend bind.ContractBackend
}

func NewJuncaswapRouter(transactOpts *bind.TransactOpts, contractAddr common.Address,
	contractBackend bind.ContractBackend) (*JuncaswapRouter, error) {
	router, err := contract.NewJuncaswapRouter(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &JuncaswapRouter{
		&contract.JuncaswapRouterSession{
			Contract:     router,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployJuncaswapRouter(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend,
	factory, weth common.Address) (common.Address, *JuncaswapRouter, error) {
	contractAddr, _, _, err := contract.DeployJuncaswapRouter(transactOpts, contractBackend, factory, weth)
	if err != nil {
		return contractAddr, nil, err
	}

	router, err := NewJuncaswapRouter(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, router, nil
}
