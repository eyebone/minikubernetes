// Copyright 2022 Google LLC
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

// Code generated by aliasgen. DO NOT EDIT.

// Package dataform aliases all exported identifiers in package
// "cloud.google.com/go/dataform/apiv1beta1/dataformpb".
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package dataform

import (
	src "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
const (
	CompilationResultAction_Relation_INCREMENTAL_TABLE                   = src.CompilationResultAction_Relation_INCREMENTAL_TABLE
	CompilationResultAction_Relation_MATERIALIZED_VIEW                   = src.CompilationResultAction_Relation_MATERIALIZED_VIEW
	CompilationResultAction_Relation_RELATION_TYPE_UNSPECIFIED           = src.CompilationResultAction_Relation_RELATION_TYPE_UNSPECIFIED
	CompilationResultAction_Relation_TABLE                               = src.CompilationResultAction_Relation_TABLE
	CompilationResultAction_Relation_VIEW                                = src.CompilationResultAction_Relation_VIEW
	FetchFileGitStatusesResponse_UncommittedFileChange_ADDED             = src.FetchFileGitStatusesResponse_UncommittedFileChange_ADDED
	FetchFileGitStatusesResponse_UncommittedFileChange_DELETED           = src.FetchFileGitStatusesResponse_UncommittedFileChange_DELETED
	FetchFileGitStatusesResponse_UncommittedFileChange_HAS_CONFLICTS     = src.FetchFileGitStatusesResponse_UncommittedFileChange_HAS_CONFLICTS
	FetchFileGitStatusesResponse_UncommittedFileChange_MODIFIED          = src.FetchFileGitStatusesResponse_UncommittedFileChange_MODIFIED
	FetchFileGitStatusesResponse_UncommittedFileChange_STATE_UNSPECIFIED = src.FetchFileGitStatusesResponse_UncommittedFileChange_STATE_UNSPECIFIED
	Repository_GitRemoteSettings_INVALID                                 = src.Repository_GitRemoteSettings_INVALID
	Repository_GitRemoteSettings_NOT_FOUND                               = src.Repository_GitRemoteSettings_NOT_FOUND
	Repository_GitRemoteSettings_TOKEN_STATUS_UNSPECIFIED                = src.Repository_GitRemoteSettings_TOKEN_STATUS_UNSPECIFIED
	Repository_GitRemoteSettings_VALID                                   = src.Repository_GitRemoteSettings_VALID
	WorkflowInvocationAction_CANCELLED                                   = src.WorkflowInvocationAction_CANCELLED
	WorkflowInvocationAction_DISABLED                                    = src.WorkflowInvocationAction_DISABLED
	WorkflowInvocationAction_FAILED                                      = src.WorkflowInvocationAction_FAILED
	WorkflowInvocationAction_PENDING                                     = src.WorkflowInvocationAction_PENDING
	WorkflowInvocationAction_RUNNING                                     = src.WorkflowInvocationAction_RUNNING
	WorkflowInvocationAction_SKIPPED                                     = src.WorkflowInvocationAction_SKIPPED
	WorkflowInvocationAction_SUCCEEDED                                   = src.WorkflowInvocationAction_SUCCEEDED
	WorkflowInvocation_CANCELING                                         = src.WorkflowInvocation_CANCELING
	WorkflowInvocation_CANCELLED                                         = src.WorkflowInvocation_CANCELLED
	WorkflowInvocation_FAILED                                            = src.WorkflowInvocation_FAILED
	WorkflowInvocation_RUNNING                                           = src.WorkflowInvocation_RUNNING
	WorkflowInvocation_STATE_UNSPECIFIED                                 = src.WorkflowInvocation_STATE_UNSPECIFIED
	WorkflowInvocation_SUCCEEDED                                         = src.WorkflowInvocation_SUCCEEDED
)

// Deprecated: Please use vars in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
var (
	CompilationResultAction_Relation_RelationType_name             = src.CompilationResultAction_Relation_RelationType_name
	CompilationResultAction_Relation_RelationType_value            = src.CompilationResultAction_Relation_RelationType_value
	FetchFileGitStatusesResponse_UncommittedFileChange_State_name  = src.FetchFileGitStatusesResponse_UncommittedFileChange_State_name
	FetchFileGitStatusesResponse_UncommittedFileChange_State_value = src.FetchFileGitStatusesResponse_UncommittedFileChange_State_value
	File_google_cloud_dataform_v1beta1_dataform_proto              = src.File_google_cloud_dataform_v1beta1_dataform_proto
	Repository_GitRemoteSettings_TokenStatus_name                  = src.Repository_GitRemoteSettings_TokenStatus_name
	Repository_GitRemoteSettings_TokenStatus_value                 = src.Repository_GitRemoteSettings_TokenStatus_value
	WorkflowInvocationAction_State_name                            = src.WorkflowInvocationAction_State_name
	WorkflowInvocationAction_State_value                           = src.WorkflowInvocationAction_State_value
	WorkflowInvocation_State_name                                  = src.WorkflowInvocation_State_name
	WorkflowInvocation_State_value                                 = src.WorkflowInvocation_State_value
)

// `CancelWorkflowInvocation` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CancelWorkflowInvocationRequest = src.CancelWorkflowInvocationRequest

// Represents the author of a Git commit.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CommitAuthor = src.CommitAuthor

// `CommitWorkspaceChanges` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CommitWorkspaceChangesRequest = src.CommitWorkspaceChangesRequest

// Represents the result of compiling a Dataform project.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResult = src.CompilationResult

// Represents a single Dataform action in a compilation result.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction = src.CompilationResultAction

// Represents an assertion upon a SQL query which is required return zero
// rows.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Assertion = src.CompilationResultAction_Assertion
type CompilationResultAction_Assertion_ = src.CompilationResultAction_Assertion_

// Represents a relation which is not managed by Dataform but which may be
// referenced by Dataform actions.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Declaration = src.CompilationResultAction_Declaration
type CompilationResultAction_Declaration_ = src.CompilationResultAction_Declaration_

// Represents a list of arbitrary database operations.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Operations = src.CompilationResultAction_Operations
type CompilationResultAction_Operations_ = src.CompilationResultAction_Operations_

// Represents a database relation.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Relation = src.CompilationResultAction_Relation
type CompilationResultAction_Relation_ = src.CompilationResultAction_Relation_

// Contains settings for relations of type `INCREMENTAL_TABLE`.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Relation_IncrementalTableConfig = src.CompilationResultAction_Relation_IncrementalTableConfig

// Indicates the type of this relation.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResultAction_Relation_RelationType = src.CompilationResultAction_Relation_RelationType

// An error encountered when attempting to compile a Dataform project.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CompilationResult_CompilationError = src.CompilationResult_CompilationError
type CompilationResult_GitCommitish = src.CompilationResult_GitCommitish
type CompilationResult_Workspace = src.CompilationResult_Workspace

// `CreateCompilationResult` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CreateCompilationResultRequest = src.CreateCompilationResultRequest

// `CreateRepository` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CreateRepositoryRequest = src.CreateRepositoryRequest

// `CreateWorkflowInvocation` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CreateWorkflowInvocationRequest = src.CreateWorkflowInvocationRequest

// `CreateWorkspace` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type CreateWorkspaceRequest = src.CreateWorkspaceRequest

// DataformClient is the client API for Dataform service. For semantics around
// ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type DataformClient = src.DataformClient

// DataformServer is the server API for Dataform service.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type DataformServer = src.DataformServer

// `DeleteRepository` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type DeleteRepositoryRequest = src.DeleteRepositoryRequest

// `DeleteWorkflowInvocation` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type DeleteWorkflowInvocationRequest = src.DeleteWorkflowInvocationRequest

// `DeleteWorkspace` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type DeleteWorkspaceRequest = src.DeleteWorkspaceRequest

// `FetchFileDiff` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileDiffRequest = src.FetchFileDiffRequest

// `FetchFileDiff` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileDiffResponse = src.FetchFileDiffResponse

// `FetchFileGitStatuses` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileGitStatusesRequest = src.FetchFileGitStatusesRequest

// `FetchFileGitStatuses` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileGitStatusesResponse = src.FetchFileGitStatusesResponse

// Represents the Git state of a file with uncommitted changes.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileGitStatusesResponse_UncommittedFileChange = src.FetchFileGitStatusesResponse_UncommittedFileChange

// Indicates the status of an uncommitted file change.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchFileGitStatusesResponse_UncommittedFileChange_State = src.FetchFileGitStatusesResponse_UncommittedFileChange_State

// `FetchGitAheadBehind` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchGitAheadBehindRequest = src.FetchGitAheadBehindRequest

// `FetchGitAheadBehind` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchGitAheadBehindResponse = src.FetchGitAheadBehindResponse

// `FetchRemoteBranches` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchRemoteBranchesRequest = src.FetchRemoteBranchesRequest

// `FetchRemoteBranches` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type FetchRemoteBranchesResponse = src.FetchRemoteBranchesResponse

// `GetCompilationResult` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type GetCompilationResultRequest = src.GetCompilationResultRequest

// `GetRepository` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type GetRepositoryRequest = src.GetRepositoryRequest

// `GetWorkflowInvocation` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type GetWorkflowInvocationRequest = src.GetWorkflowInvocationRequest

// `GetWorkspace` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type GetWorkspaceRequest = src.GetWorkspaceRequest

// `InstallNpmPackages` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type InstallNpmPackagesRequest = src.InstallNpmPackagesRequest

// `InstallNpmPackages` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type InstallNpmPackagesResponse = src.InstallNpmPackagesResponse

// `ListCompilationResults` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListCompilationResultsRequest = src.ListCompilationResultsRequest

// `ListCompilationResults` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListCompilationResultsResponse = src.ListCompilationResultsResponse

// `ListRepositories` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListRepositoriesRequest = src.ListRepositoriesRequest

// `ListRepositories` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListRepositoriesResponse = src.ListRepositoriesResponse

// `ListWorkflowInvocations` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListWorkflowInvocationsRequest = src.ListWorkflowInvocationsRequest

// `ListWorkflowInvocations` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListWorkflowInvocationsResponse = src.ListWorkflowInvocationsResponse

// `ListWorkspaces` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListWorkspacesRequest = src.ListWorkspacesRequest

// `ListWorkspaces` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ListWorkspacesResponse = src.ListWorkspacesResponse

// `MakeDirectory` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MakeDirectoryRequest = src.MakeDirectoryRequest

// `MakeDirectory` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MakeDirectoryResponse = src.MakeDirectoryResponse

// `MoveDirectory` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MoveDirectoryRequest = src.MoveDirectoryRequest

// `MoveDirectory` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MoveDirectoryResponse = src.MoveDirectoryResponse

// `MoveFile` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MoveFileRequest = src.MoveFileRequest

// `MoveFile` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type MoveFileResponse = src.MoveFileResponse

// `PullGitCommits` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type PullGitCommitsRequest = src.PullGitCommitsRequest

// `PushGitCommits` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type PushGitCommitsRequest = src.PushGitCommitsRequest

// `QueryCompilationResultActions` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryCompilationResultActionsRequest = src.QueryCompilationResultActionsRequest

// `QueryCompilationResultActions` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryCompilationResultActionsResponse = src.QueryCompilationResultActionsResponse

// `QueryDirectoryContents` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryDirectoryContentsRequest = src.QueryDirectoryContentsRequest

// `QueryDirectoryContents` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryDirectoryContentsResponse = src.QueryDirectoryContentsResponse

// `QueryWorkflowInvocationActions` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryWorkflowInvocationActionsRequest = src.QueryWorkflowInvocationActionsRequest

// `QueryWorkflowInvocationActions` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type QueryWorkflowInvocationActionsResponse = src.QueryWorkflowInvocationActionsResponse

// `ReadFile` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ReadFileRequest = src.ReadFileRequest

// `ReadFile` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ReadFileResponse = src.ReadFileResponse

// Describes a relation and its columns.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type RelationDescriptor = src.RelationDescriptor

// Describes a column.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type RelationDescriptor_ColumnDescriptor = src.RelationDescriptor_ColumnDescriptor

// `RemoveDirectory` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type RemoveDirectoryRequest = src.RemoveDirectoryRequest

// `RemoveFile` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type RemoveFileRequest = src.RemoveFileRequest

// Represents a Dataform Git repository.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type Repository = src.Repository

// Controls Git remote configuration for a repository.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type Repository_GitRemoteSettings = src.Repository_GitRemoteSettings

// Indicates the status of a Git authentication token.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type Repository_GitRemoteSettings_TokenStatus = src.Repository_GitRemoteSettings_TokenStatus

// `ResetWorkspaceChanges` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type ResetWorkspaceChangesRequest = src.ResetWorkspaceChangesRequest

// Represents an action identifier. If the action writes output, the output
// will be written to the referenced database object.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type Target = src.Target

// UnimplementedDataformServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type UnimplementedDataformServer = src.UnimplementedDataformServer

// `UpdateRepository` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type UpdateRepositoryRequest = src.UpdateRepositoryRequest

// Represents a single invocation of a compilation result.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WorkflowInvocation = src.WorkflowInvocation

// Represents a single action in a workflow invocation.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WorkflowInvocationAction = src.WorkflowInvocationAction

// Represents a workflow action that will run against BigQuery.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WorkflowInvocationAction_BigQueryAction = src.WorkflowInvocationAction_BigQueryAction

// Represents the current state of an workflow invocation action.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WorkflowInvocationAction_State = src.WorkflowInvocationAction_State

// Represents the current state of a workflow invocation.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WorkflowInvocation_State = src.WorkflowInvocation_State

// Represents a Dataform Git workspace.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type Workspace = src.Workspace

// `WriteFile` request message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WriteFileRequest = src.WriteFileRequest

// `WriteFile` response message.
//
// Deprecated: Please use types in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
type WriteFileResponse = src.WriteFileResponse

// Deprecated: Please use funcs in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
func NewDataformClient(cc grpc.ClientConnInterface) DataformClient { return src.NewDataformClient(cc) }

// Deprecated: Please use funcs in: cloud.google.com/go/dataform/apiv1beta1/dataformpb
func RegisterDataformServer(s *grpc.Server, srv DataformServer) { src.RegisterDataformServer(s, srv) }
