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
  createGetPaginationActionType,
  createGetPaginationSuccessActionType,
  createGetPaginationFailureActionType,
  getPagination,
  getPaginationSuccess,
  getPaginationFailure,
} from '../pagination'

import {
  testActionCreatorsTypeByGroups,
  testActionCreatorsPayload,
} from './helpers'

describe('pagination action creators', function () {
  const name = 'ENTITY'

  const GET_ENTITY_LIST_REQUEST = createGetPaginationActionType(name)
  const GET_ENTITY_LIST_SUCCESS = createGetPaginationSuccessActionType(name)
  const GET_ENTITY_LIST_FAILURE = createGetPaginationFailureActionType(name)
  const getEntityListRequest = getPagination(name)
  const getEntityListSuccess = getPaginationSuccess(name)
  const getEntityListFailure = getPaginationFailure(name)

  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getEntityListRequest,
            getEntityListSuccess,
            getEntityListFailure,
          ],
          actionTypes: [
            GET_ENTITY_LIST_REQUEST,
            GET_ENTITY_LIST_SUCCESS,
            GET_ENTITY_LIST_FAILURE,
          ],
        },
      ]
    )
  })

  describe('should have correct structure', function () {
    const entityListRequestPayload = { params: { page: 1, pageSize: 10 }, entityId: 'entity-id' }
    const entityListSuccessPayload = {
      entities: [{ name: 'entity' }], totalCount: 1, entityId: 'entity-id',
    }
    const entityListFailurePayload = { error: { status: 404 }}

    describe('payload', function () {
      testActionCreatorsPayload({
        getPagination: {
          actual: getEntityListRequest(
            entityListRequestPayload.params,
            entityListRequestPayload.entityId,
          ),
          expected: entityListRequestPayload,
        },
        getPaginationSuccess: {
          actual: getEntityListSuccess(
            entityListSuccessPayload.entities,
            entityListSuccessPayload.totalCount,
            entityListSuccessPayload.entityId,
          ),
          expected: entityListSuccessPayload,
        },
        getPaginationFailure: {
          actual: getEntityListFailure(entityListFailurePayload.error),
          expected: entityListFailurePayload,
        },
      })
    })
  })
})
