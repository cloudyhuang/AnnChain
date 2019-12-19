// Copyright © 2017 ZhongAn Technology
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/spf13/viper"

	"github.com/dappledger/AnnChain/chain/app/evm"
	"github.com/dappledger/AnnChain/gemmill/types"
)

type AppMaker func(*viper.Viper) (types.Application, error)

var AppMap = map[string]AppMaker{
	"evm": func(c *viper.Viper) (types.Application, error) {
		return evm.NewEVMApp(c)
	},
}