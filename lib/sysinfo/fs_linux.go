/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sysinfo

import (
	"os"
	"syscall"
	"time"
)

func CreateTime(name string) (*time.Time, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	st := fi.Sys().(*syscall.Stat_t)
	tm := time.Unix(int64(st.Ctim.Sec), int64(st.Ctim.Nsec))
	return &tm, nil
}
