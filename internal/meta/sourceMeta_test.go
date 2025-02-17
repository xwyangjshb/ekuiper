// Copyright 2021 EMQ Technologies Co., Ltd.
//
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

package meta

import (
	"github.com/lf-edge/ekuiper/internal/conf"
	"path"
	"testing"
)

func TestGetMqttSourceMeta(t *testing.T) {
	confDir, err := conf.GetConfLoc()
	if nil != err {
		return
	}

	if err = ReadSourceMetaFile(path.Join(confDir, "mqtt_source.json"), true); nil != err {
		return
	}

	showMeta, err := GetSourceMeta("mqtt", "zh_CN")
	if nil != err {
		t.Error(err)
	}

	fields := showMeta.ConfKeys["default"]

	if len(fields) == 0 {
		t.Errorf("default fields %v", fields)
	}

}
