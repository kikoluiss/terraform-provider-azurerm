package capacities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResumeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Resume ...
func (c CapacitiesClient) Resume(ctx context.Context, id CapacitiesId) (result ResumeOperationResponse, err error) {
	req, err := c.preparerForResume(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capacities.CapacitiesClient", "Resume", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForResume(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capacities.CapacitiesClient", "Resume", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ResumeThenPoll performs Resume then polls until it's completed
func (c CapacitiesClient) ResumeThenPoll(ctx context.Context, id CapacitiesId) error {
	result, err := c.Resume(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Resume: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Resume: %+v", err)
	}

	return nil
}

// preparerForResume prepares the Resume request.
func (c CapacitiesClient) preparerForResume(ctx context.Context, id CapacitiesId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resume", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForResume sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (c CapacitiesClient) senderForResume(ctx context.Context, req *http.Request) (future ResumeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
