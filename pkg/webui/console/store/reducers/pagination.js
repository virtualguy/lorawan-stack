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

import { getEntityId } from '../../../lib/selectors/id'
import { createGetPaginationSuccessActionType } from '../actions/pagination'

const defaultState = {
  ids: [],
  totalCount: 0,
}

const createNamedPaginationReducer = function (reducerName = '') {
  const GET_PAGINATION_SUCCESS = createGetPaginationSuccessActionType(reducerName)

  return function (state = defaultState, action) {
    switch (action.type) {
    case GET_PAGINATION_SUCCESS:
      return {
        ...state,
        totalCount: action.totalCount,
        ids: action.entities.map(entity => getEntityId(entity)),
      }
    default:
      return state
    }
  }
}

export default createNamedPaginationReducer
