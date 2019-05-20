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
import * as applications from '../actions/applications'
import * as application from '../actions/application'

const getApplicationsLogic = createLogic({
  type: applications.GET_APPS_LIST,
  async process ({ action }, dispatch, done) {
    const { page, pageSize: limit } = action.params
    try {
      const data = await api.applications.list({ page, limit })
      dispatch(applications.getApplicationsSuccess(data.applications, data.totalCount))
    } catch (error) {
      dispatch(applications.getApplicationsFailure(error))
    }

    done()
  },
})

const getApplicationLogic = createLogic({
  type: applications.GET_APP,
  async process ({ action }, dispatch, done) {
    const { id, meta: { selectors = []}} = action

    try {
      const app = await api.application.get(id, selectors)
      dispatch(application.startApplicationEventsStream(id))
      dispatch(applications.getApplicationSuccess(app))
    } catch (e) {
      dispatch(applications.getApplicationFailure(e))
    }

    done()
  },
})

const getApplicationRightsLogic = createLogic({
  type: applications.GET_APP_RIGHTS_LIST,
  async process ({ action }, dispatch, done) {
    const { entityId } = action
    try {
      const result = await api.rights.applications(entityId)

      dispatch(applications.getApplicationRightsSuccess(result.rights.sort(), entityId))
    } catch (error) {
      dispatch(applications.getApplicationRightsFailure(error))
    }

    done()
  },
})

const getApplicationApiKeysLogic = createLogic({
  type: applications.GET_APP_API_KEYS,
  async process ({ action }, dispatch, done) {
    const { entityId, params } = action
    try {
      const res = await api.application.apiKeys.list(entityId, params)

      dispatch(
        applications.getApplicationApiKeysSuccess(
          res.api_keys.map(key => ({ ...key, entityId })),
          res.totalCount,
          entityId,
        )
      )
    } catch (error) {
      dispatch(applications.getApplicationApiKeysFailure(error))
    }

    done()
  },
})

const getApplicationApiKeyLogic = createLogic({
  type: applications.GET_APP_API_KEY,
  async process ({ action }, dispatch, done) {
    const { entityId, keyId } = action
    try {
      const res = await api.application.apiKeys.get(entityId, keyId)

      dispatch(applications.getApplicationApiKeySuccess({ ...res, entityId }))
    } catch (error) {
      dispatch(applications.getApplicationApiKeyFailure(error))
    }

    done()
  },
})

export default [
  getApplicationsLogic,
  getApplicationLogic,
  getApplicationRightsLogic,
  getApplicationApiKeysLogic,
  getApplicationApiKeyLogic,
]
