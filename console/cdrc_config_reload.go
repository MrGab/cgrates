/*
Rating system designed to be used in VoIP Carriers World
Copyright (C) 2012-2015 ITsysCOM

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package console

import (
	"github.com/cgrates/cgrates/apier/v1"
)

func init() {
	c := &CmdCdrcConfigReload{
		name:      "cdrc_config_reload",
		rpcMethod: "ApierV1.ReloadCdrcConfig",
	}
	commands[c.Name()] = c
	c.CommandExecuter = &CommandExecuter{c}
}

// Commander implementation
type CmdCdrcConfigReload struct {
	name      string
	rpcMethod string
	rpcParams *v1.AttrReloadConfig
	*CommandExecuter
}

func (self *CmdCdrcConfigReload) Name() string {
	return self.name
}

func (self *CmdCdrcConfigReload) RpcMethod() string {
	return self.rpcMethod
}

func (self *CmdCdrcConfigReload) RpcParams(reset bool) interface{} {
	if reset || self.rpcParams == nil {
		self.rpcParams = new(v1.AttrReloadConfig)
	}
	return self.rpcParams
}

func (self *CmdCdrcConfigReload) PostprocessRpcParams() error {
	return nil
}

func (self *CmdCdrcConfigReload) RpcResult() interface{} {
	var s string
	return &s
}