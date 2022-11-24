// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package feedtoken

import (
<<<<<<< HEAD
	"github.com/strong-network/azure-devops-go-api/azuredevops/v6"
=======
	"github.com/strong-network/azure-devops-go-api/azuredevops/v6"
>>>>>>> a488a0a323616cb06afcd7d54b36eb698d590557
)

// A cut-down version of SessionToken that just has what FeedSessionTokenController needs to serve the UI and which actually generates a TypeScript type for the UI to use
type FeedSessionToken struct {
	Token   *string           `json:"token,omitempty"`
	ValidTo *azuredevops.Time `json:"validTo,omitempty"`
}
