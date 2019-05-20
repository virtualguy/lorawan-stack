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

import { createLogic } from 'redux-logic'

import api from '../../api'
import * as gateways from '../actions/gateways'
import * as gateway from '../actions/gateway'

const getGatewaysLogic = createLogic({
  type: gateways.GET_GTWS_LIST,
  async process ({ action }, dispatch, done) {
    const { page, pageSize: limit } = action.params

    try {
      const data = await api.gateways.list({ page, limit }, [ 'name', 'description', 'frequency_plan_id' ])
      dispatch(gateways.getGatewaysSuccess(data.gateways, data.totalCount))
    } catch (error) {
      dispatch(gateways.getGatewaysFailure(error))
    }

    done()
  },
})

const getGatewayLogic = createLogic({
  type: gateways.GET_GTW,
  async process ({ action }, dispatch, done) {
    const { id, meta = { selectors: []}} = action

    try {
      const gtw = await api.gateway.get(id, meta.selectors)
      dispatch(gateway.startGatewayEventsStream(id))
      dispatch(gateways.getGatewaySuccess(gtw))
    } catch (error) {
      dispatch(gateways.getGatewayFailure(error))
    }

    done()
  },
})

const getGatewaysRightsLogic = createLogic({
  type: gateways.GET_GTW_RIGHTS_LIST,
  async process ({ action }, dispatch, done) {
    const { entityId } = action
    try {
      const result = await api.rights.gateways(entityId)

      dispatch(gateways.getGatewayRightsSuccess(result.rights.sort(), entityId))
    } catch (error) {
      dispatch(gateways.getGatewayRightsFailure(error))
    }

    done()
  },
})

const getGatewayApiKeysLogic = createLogic({
  type: gateways.GET_GTW_API_KEYS,
  async process ({ action }, dispatch, done) {
    const { entityId, params } = action
    try {
      const res = await api.gateway.apiKeys.list(entityId, params)
      dispatch(
        gateways.getGatewayApiKeysSuccess(
          res.api_keys.map(key => ({ ...key, entityId })),
          res.totalCount,
          entityId
        )
      )
    } catch (error) {
      dispatch(gateways.getGatewayApiKeysFailure(error))
    }

    done()
  },
})

const getGatewayApiKeyLogic = createLogic({
  type: gateways.GET_GTW_API_KEY,
  async process ({ action }, dispatch, done) {
    const { entityId, keyId } = action
    try {
      const res = await api.gateway.apiKeys.get(entityId, keyId)

      dispatch(gateways.getGatewayApiKeySuccess({ ...res, entityId }))
    } catch (error) {
      dispatch(gateways.getGatewayApiKeyFailure(error))
    }

    done()
  },
})

export default [
  getGatewaysLogic,
  getGatewayLogic,
  getGatewaysRightsLogic,
  getGatewayApiKeysLogic,
  getGatewayApiKeyLogic,
]
