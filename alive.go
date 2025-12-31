// Copyright 2025 肖其顿 (XIAO QI DUN)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command alive 远程虽断，现场犹存 — 赋予无人值守瞬间以永恒
package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	if !windows.GetCurrentProcessToken().IsElevated() {
		runAsAdmin()
		return
	}
	sid, err := getSessionID()
	if err != nil {
		os.Exit(1)
	}
	cmd := exec.Command("tscon", strconv.FormatUint(uint64(sid), 10), "/dest:console")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if cmd.Run() != nil {
		os.Exit(1)
	}
}

func runAsAdmin() {
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	var args []string
	for _, v := range os.Args[1:] {
		args = append(args, syscall.EscapeArg(v))
	}
	_ = windows.ShellExecute(
		0,
		windows.StringToUTF16Ptr("runas"),
		windows.StringToUTF16Ptr(exe),
		windows.StringToUTF16Ptr(strings.Join(args, " ")),
		windows.StringToUTF16Ptr(cwd),
		windows.SW_SHOWNORMAL,
	)
	os.Exit(0)
}

func getSessionID() (uint32, error) {
	var sessionID uint32
	if err := windows.ProcessIdToSessionId(uint32(os.Getpid()), &sessionID); err != nil {
		return 0, err
	}
	return sessionID, nil
}
