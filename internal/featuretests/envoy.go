// Copyright © 2019 VMware
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

package featuretests

// envoy helpers

import (
	"time"

	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/projectcontour/contour/internal/protobuf"
)

func routeCluster(cluster string) *envoy_api_v2_route.Route_Route {
	return &envoy_api_v2_route.Route_Route{
		Route: &envoy_api_v2_route.RouteAction{
			ClusterSpecifier: &envoy_api_v2_route.RouteAction_Cluster{
				Cluster: cluster,
			},
		},
	}
}

func withResponseTimeout(route *envoy_api_v2_route.Route_Route, timeout time.Duration) *envoy_api_v2_route.Route_Route {
	route.Route.Timeout = protobuf.Duration(timeout)
	return route
}

func withIdleTimeout(route *envoy_api_v2_route.Route_Route, timeout time.Duration) *envoy_api_v2_route.Route_Route {
	route.Route.IdleTimeout = protobuf.Duration(timeout)
	return route
}

func withMirrorPolicy(route *envoy_api_v2_route.Route_Route, mirror string) *envoy_api_v2_route.Route_Route {
	route.Route.RequestMirrorPolicy = &envoy_api_v2_route.RouteAction_RequestMirrorPolicy{
		Cluster: mirror,
	}
	return route
}
