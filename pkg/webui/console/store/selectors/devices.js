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

import { GET_DEVICES_LIST_BASE, SEARCH_DEVICES_LIST_BASE } from '../actions/devices'
import { createFetchingSelector } from './fetching'
import { createErrorSelector } from './error'

const selectDevsStore = state => state.devices
const selectDevsFetching = createFetchingSelector([ GET_DEVICES_LIST_BASE, SEARCH_DEVICES_LIST_BASE ])
const selectDevsError = createErrorSelector([ GET_DEVICES_LIST_BASE, SEARCH_DEVICES_LIST_BASE ])

// devices
export const selectDevices = state => selectDevsStore(state).devices
export const selectDevicesTotalCount = state => selectDevsStore(state).totalCount
export const selectDevicesFetching = state => selectDevsFetching(state)
export const selectDevicesError = state => selectDevsError(state)
