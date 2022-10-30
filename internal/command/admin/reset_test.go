// Copyright 2022 Giuseppe De Palma, Matteo Trentin
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

package admin

import (
	"context"
	"errors"
	"testing"

	"github.com/funlessdev/fl-cli/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestResetRun(t *testing.T) {
	reset := reset{}
	ctx := context.TODO()

	deployer := mocks.NewDevDeployer(t)

	t.Run("should return error when setup client fails", func(t *testing.T) {
		deployer.On("Setup", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(func(ctx context.Context, coreImg string, workerImg string) error {
				return errors.New("error")
			}).Once()

		_, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)
		require.Error(t, err)
	})

	t.Run("should return error when removing Core fails", func(t *testing.T) {
		deployer.On("Setup", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(func(ctx context.Context, coreImg string, workerImg string) error {
				return nil
			})
		deployer.On("RemoveCoreContainer", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)
		require.Error(t, err)
	})

	t.Run("should return error when removing Worker fails", func(t *testing.T) {
		deployer.On("RemoveCoreContainer", ctx).Return(func(ctx context.Context) error {
			return nil
		})
		deployer.On("RemoveWorkerContainer", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)
		require.Error(t, err)
	})

	t.Run("should return error when removing Prometheus fails", func(t *testing.T) {
		deployer.On("RemoveWorkerContainer", ctx).Return(func(ctx context.Context) error {
			return nil
		})
		deployer.On("RemovePromContainer", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)
		require.Error(t, err)
	})

	t.Run("should return error when removing FL network fails", func(t *testing.T) {
		deployer.On("RemovePromContainer", ctx).Return(func(ctx context.Context) error {
			return nil
		})
		deployer.On("RemoveFLNetwork", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)
		require.Error(t, err)
	})

	t.Run("successful prints when everything goes well", func(t *testing.T) {
		deployer.On("RemoveFLNetwork", ctx).Return(func(ctx context.Context) error {
			return nil
		})

		outbuf, testLogger := testLogger()
		err := reset.Run(ctx, deployer, testLogger)

		expectedOutput := `Removing local FunLess deployment...

Removing Core container... ☠️
done
Removing Worker container... 🔪
done
Removing Prometheus container... ⚰️
done
Removing fl network... ✂️
done

All clear! 👍
`
		assert.NoError(t, err)
		assert.Equal(t, expectedOutput, outbuf.String())
	})

}
