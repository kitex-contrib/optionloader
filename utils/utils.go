// Copyright 2024 CloudWeGo Authors
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

package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"time"
)

// PathExists check whether the file or directory exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}
func Printpath() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current directory:", dir)
}

func ParseDuration(s string) (time.Duration, error) {
	// 定义一个正则表达式来匹配时间量和单位
	// 假设时间量是一个整数，单位可以是 s(秒), m(分钟), h(小时) 等
	re := regexp.MustCompile(`^(\d+)([smh])$`)
	matches := re.FindStringSubmatch(s)
	if matches == nil || len(matches) != 3 {
		return 0, fmt.Errorf("invalid duration format: %s", s)
	}

	// 将时间量从字符串转换为整数
	amount, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid duration amount: %s", matches[1])
	}

	// 根据单位计算纳秒数
	var duration time.Duration
	switch matches[2] {
	case "s":
		duration = time.Duration(amount) * time.Second
	case "m":
		duration = time.Duration(amount) * time.Minute
	case "h":
		duration = time.Duration(amount) * time.Hour
	default:
		return 0, fmt.Errorf("unsupported duration unit: %s", matches[2])
	}

	return duration, nil
}
