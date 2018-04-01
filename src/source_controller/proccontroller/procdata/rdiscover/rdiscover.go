/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package rdiscover

import (
	"encoding/json"
	"os"
	"time"

	"configcenter/src/common/RegisterDiscover"
	"configcenter/src/common/blog"
	"configcenter/src/common/types"
	"configcenter/src/common/version"

	"context"
)

// RegDiscover register and discover
type RegDiscover struct {
	ip      string
	port    uint
	isSSL   bool
	rd      *RegisterDiscover.RegDiscover
	rootCtx context.Context
	cancel  context.CancelFunc
}

// NewRegDiscover create a RegDiscover object
func NewRegDiscover(zkserv string, ip string, port uint, isSSL bool) *RegDiscover {
	return &RegDiscover{
		ip:    ip,
		port:  port,
		isSSL: isSSL,
		rd:    RegisterDiscover.NewRegDiscoverEx(zkserv, 10*time.Second),
	}
}

// Start the register and discover
func (r *RegDiscover) Start() error {
	//create root context
	r.rootCtx, r.cancel = context.WithCancel(context.Background())

	//start regdiscover
	if err := r.rd.Start(); err != nil {
		blog.Error("fail to start register and discover serv. err:%s", err.Error())
		return err
	}

	// register proccontroller server
	if err := r.registerProcController(); err != nil {
		blog.Error("fail to register proccontroller(%s), err:%s", r.ip, err.Error())
		return err
	}

	// here: discover other services

	for {
		select {
		case <-r.rootCtx.Done():
			blog.Warn("register and discover serv done")
			return nil
		}
	}
}

// Stop the register and discover
func (r *RegDiscover) Stop() error {
	r.cancel()

	r.rd.Stop()

	return nil
}

func (r *RegDiscover) registerProcController() error {
	procController := new(types.ProcControllerServInfo)

	procController.IP = r.ip
	procController.Port = r.port
	procController.Scheme = "http"
	if r.isSSL {
		procController.Scheme = "https"
	}

	procController.Version = version.GetVersion()
	procController.Pid = os.Getpid()

	data, err := json.Marshal(procController)
	if err != nil {
		blog.Error("fail to marshal ProcController server info to json. err:%s", err.Error())
		return err
	}

	path := types.CC_SERV_BASEPATH + "/" + types.CC_MODULE_PROCCONTROLLER + "/" + r.ip

	return r.rd.RegisterAndWatchService(path, data)
}
