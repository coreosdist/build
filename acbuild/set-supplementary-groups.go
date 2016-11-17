// Copyright 2015 The appc Authors
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

package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	cmdSetSupplementaryGroups = &cobra.Command{
		Use:     "set-supp-groups [GROUPS]",
		Short:   "Set supplementary GID's",
		Long:    "Set the supplementary groups the app will run as inside the container",
		Example: "acbuild set-supp-groups 200 300 400",
		Run:     runWrapper(runSetSuppGroups),
	}
)

func init() {
	cmdAcbuild.AddCommand(cmdSetSupplementaryGroups)
}

func runSetSuppGroups(cmd *cobra.Command, args []string) (exit int) {
	if len(args) == 0 {
		cmd.Usage()
		return 1
	}

	if debug {
		stderr("Setting supplementary groups to ", args)
	}

	var intArgs = []int{}
	for _, stringArg := range args {
		intArg, err := strconv.Atoi(stringArg)
		if err != nil {
			stderr("error parsing group argument %v", err)
			return getErrorCode(err)
		}
		intArgs = append(intArgs, intArg)
	}

	err := newACBuild().SetSuppGroups(intArgs)

	if err != nil {
		stderr("set-supp-groups: %v", err)
		return getErrorCode(err)
	}

	return exit
}
