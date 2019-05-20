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

import { getGatewayId } from '../../../lib/selectors/id'
import {
  GET_GTW,
  GET_GTW_SUCCESS,
  GET_GTWS_LIST_SUCCESS,
} from '../actions/gateways'

const gateway = function (state = {}, gateway) {
  return {
    ...state,
    ...gateway,
  }
}

const defaultState = {
  entities: {},
  selectedGateway: null,
}

const gateways = function (state = defaultState, action) {
  switch (action.type) {
  case GET_GTW:
    return {
      ...state,
      selectedGateway: action.id,
    }
  case GET_GTWS_LIST_SUCCESS:
    const entities = action.entities.reduce(function (acc, app) {
      const id = getGatewayId(app)

      acc[id] = gateway(acc[id], app)
      return acc
    }, { ...state.entities })

    return {
      ...state,
      entities,
    }
  case GET_GTW_SUCCESS:
    const id = getGatewayId(action.gateway)

    return {
      ...state,
      entities: {
        ...state.entities,
        [id]: gateway(state[id], action.gateway),
      },
    }
  default:
    return state
  }
}

export default gateways
