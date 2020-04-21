// +build unit

package chains

import (
	"encoding/json"
	"net/http"
)

var patchChainByUUIDTests = []HTTPRouteTests{
	{
		name:       "TestPatchChainByUUIDByID200",
		chainAgent: UseMockChainRegistry,
		httpMethod: http.MethodPatch,
		path:       "/chains/1",
		body: func() []byte {
			body, _ := json.Marshal(&PatchRequest{
				URLs: []string{"http://test.com"},
			})
			return body
		},
		expectedStatusCode:  http.StatusOK,
		expectedContentType: expectedSuccessStatusContentType,
		expectedBody:        func() string { return expectedSuccessStatusBody },
	},
	{
		name:       "TestPatchChainByUUID400WithWrongURL",
		chainAgent: UseMockChainRegistry,
		httpMethod: http.MethodPatch,
		path:       "/chains/1",
		body: func() []byte {
			body, _ := json.Marshal(&PatchRequest{
				URLs: []string{"test.com"},
			})
			return body
		},
		expectedStatusCode:  http.StatusBadRequest,
		expectedContentType: expectedErrorStatusContentType,
		expectedBody: func() string {
			return expectedNotUniqueURLsError
		},
	},
	{
		name:                "TestPatchChainByUUID400WrongBody",
		chainAgent:          UseMockChainRegistry,
		httpMethod:          http.MethodPatch,
		path:                "/chains/1",
		body:                func() []byte { return []byte(`{"unknownField":"error"}`) },
		expectedStatusCode:  http.StatusBadRequest,
		expectedContentType: expectedErrorStatusContentType,
		expectedBody: func() string {
			return expectedUnknownBodyError
		},
	},
	{
		name:       "TestPatchChainByUUID404",
		chainAgent: UseErrorChainRegistry,
		httpMethod: http.MethodPatch,
		path:       "/chains/0",
		body: func() []byte {
			body, _ := json.Marshal(&PatchRequest{
				URLs: []string{"http://test.com"},
			})
			return body
		},
		expectedStatusCode:  http.StatusNotFound,
		expectedContentType: expectedErrorStatusContentType,
		expectedBody:        func() string { return expectedNotFoundErrorBody },
	},
	{
		name:       "TestPatchChainByUUID500",
		chainAgent: UseErrorChainRegistry,
		httpMethod: http.MethodPatch,
		path:       "/chains/1",
		body: func() []byte {
			body, _ := json.Marshal(&PatchRequest{
				URLs: []string{"http://test.com"},
			})
			return body
		},
		expectedStatusCode:  http.StatusInternalServerError,
		expectedContentType: expectedErrorStatusContentType,
		expectedBody:        func() string { return expectedInternalServerErrorBody },
	},
}