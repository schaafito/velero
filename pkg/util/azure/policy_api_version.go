/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/azcore/runtime/policy_api_version.go
type apiVersionPolicy struct {
	location runtime.APIVersionLocation
	name     string
	version  string
}

func (a *apiVersionPolicy) Do(req *policy.Request) (*http.Response, error) {
	if a.version != "" {
		switch a.location {
		case runtime.APIVersionLocationHeader:
			req.Raw().Header.Set(a.name, a.version)
		case runtime.APIVersionLocationQueryParam:
			q := req.Raw().URL.Query()
			q.Set(a.name, a.version)
			req.Raw().URL.RawQuery = q.Encode()
		default:
			return nil, fmt.Errorf("unknown APIVersionLocation %d", a.location)
		}
	}
	return req.Next()
}
