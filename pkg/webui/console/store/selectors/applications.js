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

import { GET_APPS_LIST_BASE, SEARCH_APPS_LIST_BASE } from '../actions/applications'
import { createFetchingSelector } from './fetching'
import { createErrorSelector } from './error'

const selectAppsStore = state => state.applications
const selectAppsFetching = createFetchingSelector([ GET_APPS_LIST_BASE, SEARCH_APPS_LIST_BASE ])
const selectAppsError = createErrorSelector([ GET_APPS_LIST_BASE, SEARCH_APPS_LIST_BASE ])

// applications
export const selectApplications = state => selectAppsStore(state).applications
export const selectApplicationsTotalCount = state => selectAppsStore(state).totalCount
export const selectApplicationsFetching = state => selectAppsFetching(state)
export const selectApplicationsError = state => selectAppsError(state)
