// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package accounts

import (
	"context"
	"github.com/google/uuid"
<<<<<<< HEAD
	"github.com/strong-network/azure-devops-go-api/azuredevops/v6"
=======
	"github.com/strong-network/azure-devops-go-api/azuredevops/v6"
>>>>>>> a488a0a323616cb06afcd7d54b36eb698d590557
	"net/http"
	"net/url"
)

var ResourceAreaId, _ = uuid.Parse("0d55247a-1c47-4462-9b1f-5e2125590ee6")

type Client interface {
	// Get a list of accounts for a specific owner or a specific member. One of the following parameters is required: ownerId, memberId.
	GetAccounts(context.Context, GetAccountsArgs) (*[]Account, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// Get a list of accounts for a specific owner or a specific member. One of the following parameters is required: ownerId, memberId.
func (client *ClientImpl) GetAccounts(ctx context.Context, args GetAccountsArgs) (*[]Account, error) {
	queryParams := url.Values{}
	if args.OwnerId != nil {
		queryParams.Add("ownerId", (*args.OwnerId).String())
	}
	if args.MemberId != nil {
		queryParams.Add("memberId", (*args.MemberId).String())
	}
	if args.Properties != nil {
		queryParams.Add("properties", *args.Properties)
	}
	locationId, _ := uuid.Parse("229a6a53-b428-4ffb-a835-e8f36b5b4b1e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Account
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAccounts function
type GetAccountsArgs struct {
	// (optional) ID for the owner of the accounts.
	OwnerId *uuid.UUID
	// (optional) ID for a member of the accounts.
	MemberId *uuid.UUID
	// (optional)
	Properties *string
}
