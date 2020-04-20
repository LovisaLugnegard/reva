// Copyright 2018-2020 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package gateway

import (
	"context"

	ocmcore "github.com/cs3org/go-cs3apis/cs3/ocm/core/v1beta1"
	"github.com/cs3org/reva/pkg/rgrpc/status"
	"github.com/cs3org/reva/pkg/rgrpc/todo/pool"
	"github.com/pkg/errors"
)

func (s *svc) CreateOCMCoreShare(ctx context.Context, req *ocmcore.CreateOCMCoreShareRequest) (*ocmcore.CreateOCMCoreShareResponse, error) {
	c, err := pool.GetOCMCoreClient(s.c.OCMCoreEndpoint)
	if err != nil {
		return &ocmcore.CreateOCMCoreShareResponse{
			Status: status.NewInternal(ctx, err, "error getting ocm core client"),
		}, nil
	}

	res, err := c.CreateOCMCoreShare(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "gateway: error calling CreateOCMCoreShare")
	}

	return res, nil
}

func (s *svc) GetOCMCoreShare(ctx context.Context, req *ocmcore.GetOCMCoreShareRequest) (*ocmcore.GetOCMCoreShareResponse, error) {
	c, err := pool.GetOCMCoreClient(s.c.OCMCoreEndpoint)
	if err != nil {
		return &ocmcore.GetOCMCoreShareResponse{
			Status: status.NewInternal(ctx, err, "error getting ocm core client"),
		}, nil
	}

	res, err := c.GetOCMCoreShare(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "gateway: error calling GetOCMCoreShare")
	}

	return res, nil
}

func (s *svc) ListOCMCoreShares(ctx context.Context, req *ocmcore.ListOCMCoreSharesRequest) (*ocmcore.ListOCMCoreSharesResponse, error) {
	c, err := pool.GetOCMCoreClient(s.c.OCMCoreEndpoint)
	if err != nil {
		return &ocmcore.ListOCMCoreSharesResponse{
			Status: status.NewInternal(ctx, err, "error getting ocm core client"),
		}, nil
	}

	res, err := c.CreateOCMCoreShare(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "gateway: error calling ListOCMCoreShares")
	}

	return res, nil
}
