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

import { combineReducers } from 'redux'
import {
  SHARED_NAME as APPLICATIONS_SHARED_NAME,
  SHARED_NAME_SINGLE as APPLICATION_SHARED_NAME,
  SHARED_NAME_API_KEYS as APPLICATION_API_KEYS_SHARED_NAME,
} from '../actions/applications'
import {
  SHARED_NAME as GATEWAYS_SHARED_NAME,
  SHARED_NAME_SINGLE as GATEWAY_SHARED_NAME,
  SHARED_NAME_API_KEYS as GATEWAY_API_KEYS_SHARED_NAME,
} from '../actions/gateways'
import { SHARED_NAME as DEVICE_SHARED_NAME } from '../actions/device'
import user from './user'
import client from './client'
import init from './init'
import applications from './applications'
import devices from './devices'
import device from './device'
import gateways from './gateways'
import gateway from './gateway'
import configuration from './configuration'
import apiKeys from './api-keys'
import createNamedRightsReducer from './rights'
import createNamedCollaboratorsReducer from './collaborators'
import createNamedEventsReducer from './events'
import fetching from './ui/fetching'
import error from './ui/error'
import createNamedPaginationReducer, {
  createNamedPaginationReducerById,
} from './pagination'

export default combineReducers({
  user,
  client,
  init,
  applications,
  devices,
  device,
  gateways,
  gateway,
  configuration,
  apiKeys,
  rights: combineReducers({
    applications: createNamedRightsReducer(APPLICATION_SHARED_NAME),
    gateways: createNamedRightsReducer(GATEWAY_SHARED_NAME),
  }),
  collaborators: combineReducers({
    applications: createNamedCollaboratorsReducer(APPLICATION_SHARED_NAME),
  }),
  events: combineReducers({
    applications: createNamedEventsReducer(APPLICATION_SHARED_NAME),
    devices: createNamedEventsReducer(DEVICE_SHARED_NAME),
    gateways: createNamedEventsReducer(GATEWAY_SHARED_NAME),
  }),
  ui: combineReducers({
    fetching,
    error,
  }),
  pagination: combineReducers({
    applications: createNamedPaginationReducer(APPLICATIONS_SHARED_NAME),
    gateways: createNamedPaginationReducer(GATEWAYS_SHARED_NAME),
    apiKeysByApplication: createNamedPaginationReducerById(APPLICATION_API_KEYS_SHARED_NAME),
    apiKeysByGateway: createNamedPaginationReducerById(GATEWAY_API_KEYS_SHARED_NAME),
  }),
})
