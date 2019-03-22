/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package prometheus

import (
	"net/http"

	"github.com/Comcast/trickster/internal/proxy"
)

// SeriesHandler ...
func (c Client) SeriesHandler(w http.ResponseWriter, r *http.Request) {
	u := c.BuildUpstreamURL(r)
	proxy.ProxyRequest(proxy.NewRequest(c.Name, otPrometheus, "SeriesHandler", http.MethodGet, u, r.Header, r), w)
}