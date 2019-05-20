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
  createGetApiKeysActionType,
  createGetApiKeysSuccessActionType,
  createGetApiKeysFailureActionType,
  getApiKeys,
  getApiKeysSuccess,
  getApiKeysFailure,
  createGetApiKeyActionType,
  createGetApiKeySuccessActionType,
  createGetApiKeyFailureActionType,
  getApiKey,
  getApiKeySuccess,
  getApiKeyFailure,
} from '../api-keys'

import {
  testActionCreatorsTypeByGroups,
  testActionCreatorsPayload,
} from './helpers'

describe('api keys action creators', function () {
  const name = 'ENTITY'

  const GET_ENTITY_API_KEYS = createGetApiKeysActionType(name)
  const GET_ENTITY_API_KEYS_SUCCESS = createGetApiKeysSuccessActionType(name)
  const GET_ENTITY_API_KEYS_FAILURE = createGetApiKeysFailureActionType(name)
  const getEntityApiKeysRequest = getApiKeys(name)
  const getEntityApiKeysSuccess = getApiKeysSuccess(name)
  const getEntityApiKeysFailure = getApiKeysFailure(name)

  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getEntityApiKeysRequest,
            getEntityApiKeysSuccess,
            getEntityApiKeysFailure,
          ],
          actionTypes: [
            GET_ENTITY_API_KEYS,
            GET_ENTITY_API_KEYS_SUCCESS,
            GET_ENTITY_API_KEYS_FAILURE,
          ],
        },
      ]
    )
  })
})

describe('api key action creators', function () {
  const name = 'ENTITY'

  const GET_ENTITY_API_KEY = createGetApiKeyActionType(name)
  const GET_ENTITY_API_KEY_SUCCESS = createGetApiKeySuccessActionType(name)
  const GET_ENTITY_API_KEY_FAILURE = createGetApiKeyFailureActionType(name)
  const getEntityApiKeyRequest = getApiKey(name)
  const getEntityApiKeySuccess = getApiKeySuccess(name)
  const getEntityApiKeyFailure = getApiKeyFailure(name)

  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getEntityApiKeyRequest,
            getEntityApiKeySuccess,
            getEntityApiKeyFailure,
          ],
          actionTypes: [
            GET_ENTITY_API_KEY,
            GET_ENTITY_API_KEY_SUCCESS,
            GET_ENTITY_API_KEY_FAILURE,
          ],
        },
      ]
    )
  })

  describe('should have correct structure', function () {
    const entityApiKeyRequestPayload = { entityId: 'entity-id', keyId: 'key-id' }
    const entityApiKeySuccessPayload = { key: { id: 'key-id', name: 'key-name' }}
    const entityApiKeyFailurePayload = { error: { status: 404 }}

    describe('payload', function () {
      testActionCreatorsPayload({
        getApiKeys: {
          actual: getEntityApiKeyRequest(
            entityApiKeyRequestPayload.entityId,
            entityApiKeyRequestPayload.keyId
          ),
          expected: entityApiKeyRequestPayload,
        },
        getApiKeysSuccess: {
          actual: getEntityApiKeySuccess(entityApiKeySuccessPayload.key),
          expected: entityApiKeySuccessPayload,
        },
        getApiKeysFailure: {
          actual: getEntityApiKeyFailure(entityApiKeyFailurePayload.error),
          expected: entityApiKeyFailurePayload,
        },
      })
    })
  })
})
