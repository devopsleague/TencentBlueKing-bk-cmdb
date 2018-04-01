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
 
package actions_test

import (
	"bytes"

	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateModel(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("POST", server.URL+"/topo/v1/model/mainline", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}

func TestDeleteModel(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("DELETE", server.URL+"/topo/v1/model/mainline/owners/1/objectids/1", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}

func TestSelectModel(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("GET", server.URL+"/topo/v1/model/1", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}

func TestSelectModelByClsID(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("GET", server.URL+"/topo/v1/model/1/1/1", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}

func TestSelectObjectInst(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("GET", server.URL+"/topo/v1/inst/1/1", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}

func TestSelectInstChild(t *testing.T) {
	server := CCAPITester(t)

	originData := map[string]interface{}{}
	reqbody, _ := json.Marshal(originData)
	r, _ := http.NewRequest("GET", server.URL+"/topo/v1/inst/child/1/1/1/1", bytes.NewBuffer(reqbody))
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}
