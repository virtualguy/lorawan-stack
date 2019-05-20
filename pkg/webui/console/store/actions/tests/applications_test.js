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
  GET_APP,
  GET_APP_SUCCESS,
  GET_APP_FAILURE,
  getApplicationById,
  getApplicationSuccess,
  getApplicationFailure,
} from '../applications'
import {
  testActionCreatorsTypeByGroups,
  testActionCreatorsPayload,
} from './helpers'

describe('applications action creators', function () {
  describe('should have type', function () {
    testActionCreatorsTypeByGroups(
      [
        {
          actionCreators: [
            getApplicationById,
            getApplicationSuccess,
            getApplicationFailure,
          ],
          actionTypes: [
            GET_APP,
            GET_APP_SUCCESS,
            GET_APP_FAILURE,
          ],
        },
      ]
    )
  })

  describe('should have correct structure', function () {
    const getApplicationRequestPayload = { id: 'app-id' }
    const getApplicationSuccessPayload = { application: { name: 'app-name' }}
    const getApplicationFailurePayload = { error: { status: 404 }}

    describe('payload', function () {
      testActionCreatorsPayload({
        getApplicationById: {
          actual: getApplicationById(getApplicationRequestPayload.id),
          expected: getApplicationRequestPayload,
        },
        getApplicationSuccess: {
          actual: getApplicationSuccess(getApplicationSuccessPayload.application),
          expected: getApplicationSuccessPayload,
        },
        getApplicationFailure: {
          actual: getApplicationFailure(getApplicationFailurePayload.error),
          expected: getApplicationFailurePayload,
        },
      })
    })
  })
})
