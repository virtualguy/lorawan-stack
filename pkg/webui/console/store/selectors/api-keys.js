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
  SHARED_NAME_SINGLE as APP_SHARED_NAME_SINGLE,
} from '../actions/applications'
import {
  SHARED_NAME_SINGLE as GTW_SHARED_NAME_SINGLE,
} from '../actions/gateways'
import {
  createGetApiKeysBaseActionType,
} from '../actions/api-keys'
import { selectSelectedApplicationId } from './applications'
import { selectSelectedGatewayId } from './gateways'
import {
  createPaginationIdsSelectorByEntityAndId,
  createPaginationTotalCountSelectorByEntityAndId,
} from './pagination'
import { createFetchingSelector } from './fetching'
import { createErrorSelector } from './error'

const selectStore = state => state.apiKeys

// api key
export const selectApiKeyById = (state, keyId) => selectStore(state)[keyId]

// application api key
export const selectSelectedApplicationApiKey = function (state, keyId) {
  const key = selectApiKeyById(state, keyId) || {}
  const selectedAppId = selectSelectedApplicationId(state)

  if (key.entityId === selectedAppId) {
    return key
  }
}

// application api keys
const APP_ENTITY = 'apiKeysByApplication'
const GET_APP_API_KEYS_BASE = createGetApiKeysBaseActionType(APP_SHARED_NAME_SINGLE)

const selectAppApiKeysIdsById = createPaginationIdsSelectorByEntityAndId(APP_ENTITY)
const selectAppApiKeysTotalCountById = createPaginationTotalCountSelectorByEntityAndId(APP_ENTITY)
const selectAppApiKeysFetching = createFetchingSelector(GET_APP_API_KEYS_BASE)
const selectAppApiKeysError = createErrorSelector(GET_APP_API_KEYS_BASE)

export const selectApplicationApiKeysById = (state, appId) => selectAppApiKeysIdsById(state, appId).map(id => selectApiKeyById(state, id))
export const selectApplicationTotalCountById = (state, appId) => selectAppApiKeysTotalCountById(state, appId)
export const selectApplicationApiKeysFetching = state => selectAppApiKeysFetching(state)
export const selectApplicationApiKeysError = state => selectAppApiKeysError(state)

// gateway api key
export const selectSelectedGatewayApiKey = function (state, keyId) {
  const key = selectApiKeyById(state, keyId) || {}
  const selectedGtwId = selectSelectedGatewayId(state)

  if (key.entityId === selectedGtwId) {
    return key
  }
}

// gateway api keys
const GTW_ENTITY = 'apiKeysByGateway'
const GET_GTW_API_KEYS_BASE = createGetApiKeysBaseActionType(GTW_SHARED_NAME_SINGLE)

const selectGtwApiKeysIdsById = createPaginationIdsSelectorByEntityAndId(GTW_ENTITY)
const selectGtwApiKeysTotalCountById = createPaginationTotalCountSelectorByEntityAndId(GTW_ENTITY)
const selectGtwApiKeysFetching = createFetchingSelector(GET_GTW_API_KEYS_BASE)
const selectGtwApiKeysError = createErrorSelector(GET_GTW_API_KEYS_BASE)

export const selectGatewayApiKeysById = (state, gtwId) => selectGtwApiKeysIdsById(state, gtwId).map(id => selectApiKeyById(state, id))
export const selectGatewayTotalCountById = (state, gtwId) => selectGtwApiKeysTotalCountById(state, gtwId)
export const selectGatewayApiKeysFetching = state => selectGtwApiKeysFetching(state)
export const selectGatewayApiKeysError = state => selectGtwApiKeysError(state)
