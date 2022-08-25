// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func pullFLImage(ctx context.Context, c *client.Client, image string) error {
	if err := pullImage(ctx, c, image); err != nil {
		return err
	}
	return nil
}

func pullImage(ctx context.Context, c *client.Client, image string) error {
	out, err := c.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	d := json.NewDecoder(out)

	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}

	var event *Event
	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if event.Error != "" {
			return fmt.Errorf("pulling image: %s", event.Error)
		}
	}
	return nil
}

func flNetExists(ctx context.Context, client *client.Client) (bool, types.NetworkResource, error) {
	nets, err := client.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "name", Value: "fl_net"}),
	})
	if err != nil {
		return false, types.NetworkResource{}, err
	}

	if len(nets) == 0 {
		return false, types.NetworkResource{}, nil
	}

	return true, nets[0], nil
}

func flNetCreate(ctx context.Context, client *client.Client) (string, error) {
	res, err := client.NetworkCreate(ctx, "fl_net", types.NetworkCreate{})
	if err != nil {
		return "", err
	}
	if res.Warning != "" {
		fmt.Printf("Warning creating fl_net network: %s\n", res.Warning)
	}
	return res.ID, nil
}

func startContainer(ctx context.Context, c *client.Client, containerConfig *container.Config, hostConfig *container.HostConfig) error {
	resp, err := c.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")

	if err != nil {
		return err
	}

	if err := c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}
