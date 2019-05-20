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
  createGetRightsActionType,
  createGetRightsSuccessActionType,
  createGetRightsFailureActionType,
  getRights,
  getRightsSuccess,
  getRightsFailure,
} from '../rights'
import {
  testActionCreatorsTypeByGroups,
  testActionCreatorsPayload,
} from './helpers'

describe('applications action creators', function () {
  const name = 'ENTITY'

  const GET_ENTITY_RIGHTS_REQUEST = createGetRightsActionType(name)
  const GET_ENTITY_RIGHTS_SUCCESS = createGetRightsSuccessActionType(name)
  const GET_ENTITY_RIGHTS_FAILURE = createGetRightsFailureActionType(name)
  const getEntityRightsListRequest = getRights(name)
  const getEntityRightsListSuccess = getRightsSuccess(name)
  const getEntityRightsListFailure = getRightsFailure(name)

  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getEntityRightsListRequest,
            getEntityRightsListSuccess,
            getEntityRightsListFailure,
          ],
          actionTypes: [
            GET_ENTITY_RIGHTS_REQUEST,
            GET_ENTITY_RIGHTS_SUCCESS,
            GET_ENTITY_RIGHTS_FAILURE,
          ],
        },
      ]
    )
  })

  describe('should have correct structure', function () {
    const entityRightsListRequestPayload = { entityId: 'entity-id' }
    const entityRightsListSuccessPayload = { rights: [ 'RIGHT_ENTITY_ALL' ], entityId: 'entity-id' }
    const entityRightsListFailurePayload = { error: { status: 404 }}

    describe('payload', function () {
      testActionCreatorsPayload({
        getRights: {
          actual: getEntityRightsListRequest(entityRightsListRequestPayload.entityId),
          expected: entityRightsListRequestPayload,
        },
        getRightsSuccess: {
          actual: getEntityRightsListSuccess(
            entityRightsListSuccessPayload.rights,
            entityRightsListSuccessPayload.entityId
          ),
          expected: entityRightsListSuccessPayload,
        },
        getRightsFailure: {
          actual: getEntityRightsListFailure(entityRightsListFailurePayload.error),
          expected: entityRightsListFailurePayload,
        },
      })
    })
  })
})
