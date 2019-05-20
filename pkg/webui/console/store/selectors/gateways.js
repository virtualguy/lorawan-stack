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
  GET_GTWS_LIST_BASE,
  GET_GTW_BASE,
  GET_GTW_RIGHTS_LIST_BASE,
} from '../actions/gateways'
import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from './pagination'
import {
  createRightsSelectorById,
} from './rights'
import { createFetchingSelector } from './fetching'
import { createErrorSelector } from './error'

const ENTITY = 'gateways'

// gatway
export const selectGatewayStore = state => state.gateways
export const selectGatewayById = (state, id) => selectGatewayStore(state)[id]
export const selectGatewayFetching = createFetchingSelector(GET_GTW_BASE)
export const selectGatewayError = createErrorSelector(GET_GTW_BASE)

// gateways
const selectGtwsIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectGtwsTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)
const selectGtwsFetching = createFetchingSelector(GET_GTWS_LIST_BASE)
const selectGtwsError = createErrorSelector(GET_GTWS_LIST_BASE)

export const selectGateways = state => selectGtwsIds(state).map(id => selectGatewayById(state, id))
export const selectGatewaysTotalCount = state => selectGtwsTotalCount(state)
export const selectGatewaysFetching = state => selectGtwsFetching(state)
export const selectGatewaysError = state => selectGtwsError(state)

// rights
export const selectGatewayRightsById = createRightsSelectorById(ENTITY)
export const selectGatewayRightsFetching = createFetchingSelector(GET_GTW_RIGHTS_LIST_BASE)
export const selectGatewayRightsError = createErrorSelector(GET_GTW_RIGHTS_LIST_BASE)
