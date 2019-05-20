// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

import {
  getRights,
  createGetRightsActionType,
  getRightsFailure,
  createGetRightsFailureActionType,
  getRightsSuccess,
  createGetRightsSuccessActionType,
  createGetRightsBaseActionType,
} from './rights'

import {
  getPagination,
  createGetPaginationActionType,
  getPaginationSuccess,
  createGetPaginationSuccessActionType,
  getPaginationFailure,
  createGetPaginationFailureActionType,
  createGetBaseActionType,
} from './pagination'

import {
  getApiKey,
  createGetApiKeyActionType,
  getApiKeySuccess,
  createGetApiKeySuccessActionType,
  getApiKeyFailure,
  createGetApiKeyFailureActionType,
  getApiKeys,
  createGetApiKeysActionType,
  getApiKeysSuccess,
  createGetApiKeysSuccessActionType,
  getApiKeysFailure,
  createGetApiKeysFailureActionType,
  createGetApiKeyBaseActionType,
  createGetApiKeysSharedName,
} from './api-keys'

// utils
export const SHARED_NAME = 'APPLICATIONS'
export const SHARED_NAME_SINGLE = 'APPLICATION'
export const SHARED_NAME_API_KEYS = createGetApiKeysSharedName(SHARED_NAME_SINGLE)
export const GET_APP_BASE = 'GET_APPLICATION'
export const GET_APPS_LIST_BASE = createGetBaseActionType(SHARED_NAME)
export const GET_APP_API_KEY_BASE = createGetApiKeyBaseActionType(SHARED_NAME_SINGLE)
export const GET_APP_RIGHTS_LIST_BASE = createGetRightsBaseActionType(SHARED_NAME_SINGLE)

// application
export const GET_APP = `${GET_APP_BASE}_REQUEST`
export const GET_APP_SUCCESS = `${GET_APP_BASE}_SUCCESS`
export const GET_APP_FAILURE = `${GET_APP_BASE}_FAILURE`

export const getApplicationById = (id, meta) => (
  { type: GET_APP, id, meta }
)
export const getApplicationSuccess = application => (
  { type: GET_APP_SUCCESS, application }
)
export const getApplicationFailure = error => (
  { type: GET_APP_FAILURE, error }
)

// applications
export const GET_APPS_LIST = createGetPaginationActionType(SHARED_NAME)
export const GET_APPS_LIST_SUCCESS = createGetPaginationSuccessActionType(SHARED_NAME)
export const GET_APPS_LIST_FAILURE = createGetPaginationFailureActionType(SHARED_NAME)

export const getApplications = getPagination(SHARED_NAME)
export const getApplicationsSuccess = getPaginationSuccess(SHARED_NAME)
export const getApplicationsFailure = getPaginationFailure(SHARED_NAME)

// rights
export const GET_APP_RIGHTS_LIST = createGetRightsActionType(SHARED_NAME_SINGLE)
export const GET_APP_RIGHTS_LIST_SUCCESS = createGetRightsSuccessActionType(SHARED_NAME_SINGLE)
export const GET_APP_RIGHTS_LIST_FAILURE = createGetRightsFailureActionType(SHARED_NAME_SINGLE)

export const getApplicationRights = getRights(SHARED_NAME_SINGLE)
export const getApplicationRightsSuccess = getRightsSuccess(SHARED_NAME_SINGLE)
export const getApplicationRightsFailure = getRightsFailure(SHARED_NAME_SINGLE)

// api keys
export const GET_APP_API_KEYS = createGetApiKeysActionType(SHARED_NAME_SINGLE)
export const GET_APP_API_KEYS_SUCCESS = createGetApiKeysSuccessActionType(SHARED_NAME_SINGLE)
export const GET_APP_API_KEYS_FAILURE = createGetApiKeysFailureActionType(SHARED_NAME_SINGLE)

export const getApplicationApiKeys = getApiKeys(SHARED_NAME_SINGLE)
export const getApplicationApiKeysSuccess = getApiKeysSuccess(SHARED_NAME_SINGLE)
export const getApplicationApiKeysFailure = getApiKeysFailure(SHARED_NAME_SINGLE)

// api key
export const GET_APP_API_KEY = createGetApiKeyActionType(SHARED_NAME_SINGLE)
export const GET_APP_API_KEY_SUCCESS = createGetApiKeySuccessActionType(SHARED_NAME_SINGLE)
export const GET_APP_API_KEY_FAILURE = createGetApiKeyFailureActionType(SHARED_NAME_SINGLE)

export const getApplicationApiKey = getApiKey(SHARED_NAME_SINGLE)
export const getApplicationApiKeySuccess = getApiKeySuccess(SHARED_NAME_SINGLE)
export const getApplicationApiKeyFailure = getApiKeyFailure(SHARED_NAME_SINGLE)
