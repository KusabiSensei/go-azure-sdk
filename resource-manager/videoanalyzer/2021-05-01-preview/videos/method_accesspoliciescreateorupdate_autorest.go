package videos

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPoliciesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccessPolicyEntity
}

// AccessPoliciesCreateOrUpdate ...
func (c VideosClient) AccessPoliciesCreateOrUpdate(ctx context.Context, id AccessPolicyId, input AccessPolicyEntity) (result AccessPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForAccessPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videos.VideosClient", "AccessPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "videos.VideosClient", "AccessPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAccessPoliciesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videos.VideosClient", "AccessPoliciesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAccessPoliciesCreateOrUpdate prepares the AccessPoliciesCreateOrUpdate request.
func (c VideosClient) preparerForAccessPoliciesCreateOrUpdate(ctx context.Context, id AccessPolicyId, input AccessPolicyEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAccessPoliciesCreateOrUpdate handles the response to the AccessPoliciesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c VideosClient) responderForAccessPoliciesCreateOrUpdate(resp *http.Response) (result AccessPoliciesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
