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
  getPagination,
  createGetPaginationActionType,
  getPaginationSuccess,
  createGetPaginationSuccessActionType,
  getPaginationFailure,
  createGetPaginationFailureActionType,
  createGetBaseActionType,
} from './pagination'

// utils
export const createGetApiKeysSharedName = name => (
  `${name}_API_KEYS`
)
export const createGetApiKeysBaseActionType = name => (
  createGetBaseActionType(createGetApiKeysSharedName(name))
)
export const createGetApiKeyBaseActionType = name => (
  `GET_${name}_API_KEY`
)

// api keys
export const createGetApiKeysActionType = name => (
  `${createGetPaginationActionType(createGetApiKeysSharedName(name))}`
)
export const createGetApiKeysSuccessActionType = name => (
  `${createGetPaginationSuccessActionType(createGetApiKeysSharedName(name))}`
)
export const createGetApiKeysFailureActionType = name => (
  `${createGetPaginationFailureActionType(createGetApiKeysSharedName(name))}`
)

export const getApiKeys = name => getPagination(createGetApiKeysSharedName(name))
export const getApiKeysSuccess = name => getPaginationSuccess(createGetApiKeysSharedName(name))
export const getApiKeysFailure = name => getPaginationFailure(createGetApiKeysSharedName(name))

// api key
export const createGetApiKeyActionType = name => (
  `${createGetApiKeyBaseActionType(name)}_REQUEST`
)
export const createGetApiKeySuccessActionType = name => (
  `${createGetApiKeyBaseActionType(name)}_SUCCESS`
)
export const createGetApiKeyFailureActionType = name => (
  `${createGetApiKeyBaseActionType(name)}_FAILURE`
)

export const getApiKey = name => (entityId, keyId) => (
  { type: createGetApiKeyActionType(name), entityId, keyId }
)
export const getApiKeySuccess = name => key => (
  { type: createGetApiKeySuccessActionType(name), key }
)
export const getApiKeyFailure = name => error => (
  { type: createGetApiKeyFailureActionType(name), error }
)
