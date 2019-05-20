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

// import {
//   createGetApiKeysSuccessActionType,
// } from '../actions/api-keys'
import {
  GET_APP_API_KEYS_SUCCESS,
  GET_APP_API_KEY_SUCCESS,
} from '../actions/applications'
import {
  GET_GTW_API_KEYS_SUCCESS,
} from '../actions/gateways'
import { getApiKeyId } from '../../../lib/selectors/id'


const apiKey = function (state = {}, key) {
  return {
    ...state,
    ...key,
  }
}

const apiKeys = function (state = {}, action) {
  switch (action.type) {
  case GET_APP_API_KEYS_SUCCESS:
  case GET_GTW_API_KEYS_SUCCESS:
    return action.entities.reduce(function (acc, key) {
      const id = getApiKeyId(key)

      acc[id] = apiKey(acc[id], key)
      return acc
    }, { ...state })
  case GET_APP_API_KEY_SUCCESS:
    const id = getApiKeyId(action.key)
    return {
      ...state,
      [id]: apiKey(state[id], action.key),
    }
  default:
    return state
  }
}

export default apiKeys
