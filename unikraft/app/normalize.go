// SPDX-License-Identifier: BSD-3-Clause
//
// Authors: Alexander Jung <alex@unikraft.io>
//
// Copyright (c) 2022, Unikraft GmbH.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package app

import (
	"path/filepath"
	"regexp"
	"strings"
)

// normalize a kraft project by moving deprecated attributes to their canonical
// position and injecting implicit defaults
func normalize(project *application, resolvePaths bool) error {
	absWorkingDir, err := filepath.Abs(project.WorkingDir())
	if err != nil {
		return err
	}
	project.workingDir = absWorkingDir

	// Ignore the error here, as it's a false positive
	absKraftfiles, err := absKraftfiles(project.Kraftfiles())
	if err != nil {
		return err
	}
	WithKraftfiles(absKraftfiles)(project)

	return nil
}

func absKraftfiles(kraftfiles []string) ([]string, error) {
	absKraftfiles := make([]string, len(kraftfiles))
	for i, kraftFile := range kraftfiles {
		absKraftfile, err := filepath.Abs(kraftFile)
		if err != nil {
			return nil, err
		}
		absKraftfiles[i] = absKraftfile
	}
	return absKraftfiles, nil
}

func normalizeProjectName(s string) string {
	r := regexp.MustCompile("[a-z0-9_-]")
	s = strings.ToLower(s)
	s = strings.Join(r.FindAllString(s, -1), "")
	return strings.TrimLeft(s, "_-")
}
