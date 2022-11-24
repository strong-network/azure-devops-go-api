// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pipelinepermissions

import (
	"github.com/google/uuid"
<<<<<<< HEAD
	"github.com/strong-network/azure-devops-go-api/azuredevops"
	"github.com/strong-network/azure-devops-go-api/azuredevops/pipelineschecks"
	"github.com/strong-network/azure-devops-go-api/azuredevops/webapi"
=======
	"github.com/strong-network/azure-devops-go-api/azuredevops"
	"github.com/strong-network/azure-devops-go-api/azuredevops/pipelineschecks"
	"github.com/strong-network/azure-devops-go-api/azuredevops/webapi"
>>>>>>> a488a0a323616cb06afcd7d54b36eb698d590557
)

type Permission struct {
	Authorized   *bool               `json:"authorized,omitempty"`
	AuthorizedBy *webapi.IdentityRef `json:"authorizedBy,omitempty"`
	AuthorizedOn *azuredevops.Time   `json:"authorizedOn,omitempty"`
}

type PipelinePermission struct {
	Authorized   *bool               `json:"authorized,omitempty"`
	AuthorizedBy *webapi.IdentityRef `json:"authorizedBy,omitempty"`
	AuthorizedOn *azuredevops.Time   `json:"authorizedOn,omitempty"`
	Id           *int                `json:"id,omitempty"`
}

type PipelineProcessResources struct {
	Resources *[]PipelineResourceReference `json:"resources,omitempty"`
}

type PipelineResourceReference struct {
	Authorized   *bool             `json:"authorized,omitempty"`
	AuthorizedBy *uuid.UUID        `json:"authorizedBy,omitempty"`
	AuthorizedOn *azuredevops.Time `json:"authorizedOn,omitempty"`
	DefinitionId *int              `json:"definitionId,omitempty"`
	Id           *string           `json:"id,omitempty"`
	Type         *string           `json:"type,omitempty"`
}

type ResourcePipelinePermissions struct {
	AllPipelines *Permission               `json:"allPipelines,omitempty"`
	Pipelines    *[]PipelinePermission     `json:"pipelines,omitempty"`
	Resource     *pipelineschecks.Resource `json:"resource,omitempty"`
}
