// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package admin

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/funlessdev/fl-cli/pkg/deploy"
	"github.com/funlessdev/fl-cli/pkg/log"
)

type reset struct{}

func (r *reset) Run(ctx context.Context, logger log.FLogger) error {
	logger.Info("Removing local funless deployment...\n")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		return err
	}

	_ = logger.StartSpinner("Removing Core container... ☠️")
	if err := logger.StopSpinner(deploy.RemoveFLContainer(ctx, cli, "fl-core")); err != nil {
		return err
	}

	_ = logger.StartSpinner("Removing Worker container... 🔪")
	if err := logger.StopSpinner(deploy.RemoveFLContainer(ctx, cli, "fl-worker")); err != nil {
		return err
	}

	_ = logger.StartSpinner("Removing the function containers... 🔫")
	if err := logger.StopSpinner(deploy.RemoveFunctionContainers(ctx, cli)); err != nil {
		return err
	}

	_ = logger.StartSpinner("Removing fl_net network... ✂️")
	if err := logger.StopSpinner(deploy.RemoveFLNetwork(ctx, cli, "fl_net")); err != nil {
		return err
	}

	_ = logger.StartSpinner("Removing fl_runtime_net network... ✂️")
	if err := logger.StopSpinner(deploy.RemoveFLNetwork(ctx, cli, "fl_runtime_net")); err != nil {
		return err
	}

	logger.Info("\nAll clear! 👍")

	return err
}
