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

package ocmd

import (
	"net/http"

	"github.com/cs3org/reva/pkg/appctx"
	"github.com/cs3org/reva/pkg/rhttp/router"
)

type invitesHandler struct {
	gatewayAddr string
}

func (h *invitesHandler) init(c *Config) {
	h.gatewayAddr = c.GatewaySvc
}

func (h *invitesHandler) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := appctx.GetLogger(r.Context())
		var head string
		head, r.URL.Path = router.ShiftPath(r.URL.Path)
		log.Debug().Str("head", head).Str("tail", r.URL.Path).Msg("http routing")

		switch head {
		case "accept":
			h.acceptInvite(w, r)
		case "forward":
			h.forwardInvite(w, r)
		case "":
			h.generateInviteToken(w, r)

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func (h *invitesHandler) acceptInvite(w http.ResponseWriter, r *http.Request) {
}

func (h *invitesHandler) forwardInvite(w http.ResponseWriter, r *http.Request) {
}

func (h *invitesHandler) generateInviteToken(w http.ResponseWriter, r *http.Request) {
}
