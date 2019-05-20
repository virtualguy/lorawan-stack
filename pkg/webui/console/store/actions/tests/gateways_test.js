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
  GET_GTW,
  GET_GTW_SUCCESS,
  GET_GTW_FAILURE,
  getGatewayById,
  getGatewaySuccess,
  getGatewayFailure,
} from '../gateways'
import {
  testActionCreatorsTypeByGroups,
  testActionCreatorsPayload,
} from './helpers'

describe('gateways action creators', function () {
  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getGatewayById,
            getGatewaySuccess,
            getGatewayFailure,
          ],
          actionTypes: [
            GET_GTW,
            GET_GTW_SUCCESS,
            GET_GTW_FAILURE,
          ],
        },
      ]
    )
  })

  describe('should have correct structure', function () {
    const getGatewayRequestPayload = { id: 'gtw-id' }
    const getGatewaySuccessPayload = { gateway: { name: 'gtw-name' }}
    const getGatewayFailurePayload = { error: { status: 404 }}

    describe('payload', function () {
      testActionCreatorsPayload({
        getGatewayById: {
          actual: getGatewayById(getGatewayRequestPayload.id),
          expected: getGatewayRequestPayload,
        },
        getGatewaySuccess: {
          actual: getGatewaySuccess(getGatewaySuccessPayload.gateway),
          expected: getGatewaySuccessPayload,
        },
        getGatewayFailure: {
          actual: getGatewayFailure(getGatewayFailurePayload.error),
          expected: getGatewayFailurePayload,
        },
      })
    })
  })
})
