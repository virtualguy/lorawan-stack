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
  getRightsList,
  createGetRightsListActionType,
  getRightsListFailure,
  createGetRightsListFailureActionType,
  getRightsListSuccess,
  createGetRightsListSuccessActionType,
} from './rights'

import {
  getPagination,
  createGetPaginationActionType,
  getPaginationSuccess,
  createGetPaginationSuccessActionType,
  getPaginationFailure,
  createGetPaginationFailureActionType,
  createGetBaseActionType,
} from './pagination'

// utils
export const SHARED_NAME = 'GATEWAYS'
export const SHARED_NAME_SINGLE = 'GATEWAY'
export const GET_GTWS_LIST_BASE = createGetBaseActionType(SHARED_NAME)
export const GET_GTW_BASE = 'GET_GATEWAY'

// gateway
export const GET_GTW = `${GET_GTW_BASE}_REQUEST`
export const GET_GTW_SUCCESS = `${GET_GTW_BASE}_SUCCESS`
export const GET_GTW_FAILURE = `${GET_GTW_BASE}_FAILURE`

export const getGatewayById = (id, meta) => (
  { type: GET_GTW, id, meta }
)
export const getGatewaySuccess = gateway => (
  { type: GET_GTW_SUCCESS, gateway }
)
export const getGatewayFailure = error => (
  { type: GET_GTW_FAILURE, error }
)

// gateways
export const GET_GTWS_LIST = createGetPaginationActionType(SHARED_NAME)
export const GET_GTWS_LIST_SUCCESS = createGetPaginationSuccessActionType(SHARED_NAME)
export const GET_GTWS_LIST_FAILURE = createGetPaginationFailureActionType(SHARED_NAME)

export const getGateways = getPagination(SHARED_NAME)
export const getGatewaysSuccess = getPaginationSuccess(SHARED_NAME)
export const getGatewaysFailure = getPaginationFailure(SHARED_NAME)

// rights
export const GET_GTWS_RIGHTS_LIST = createGetRightsListActionType(SHARED_NAME)
export const GET_GTWS_RIGHTS_LIST_SUCCESS = createGetRightsListSuccessActionType(SHARED_NAME)
export const GET_GTWS_RIGHTS_LIST_FAILURE = createGetRightsListFailureActionType(SHARED_NAME)

export const getGatewaysRightsList = getRightsList(SHARED_NAME)
export const getGatewaysRightsListSuccess = getRightsListSuccess(SHARED_NAME)
export const getGatewaysRightsListFailure = getRightsListFailure(SHARED_NAME)
