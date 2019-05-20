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

export const createGetRightsBaseActionType = name => (
  `GET_${name}_RIGHTS_LIST`
)

export const createGetRightsActionType = name => (
  `${createGetRightsBaseActionType(name)}_REQUEST`
)

export const createGetRightsSuccessActionType = name => (
  `${createGetRightsBaseActionType(name)}_SUCCESS`
)

export const createGetRightsFailureActionType = name => (
  `${createGetRightsBaseActionType(name)}_FAILURE`
)

export const getRights = name => entityId => (
  { type: createGetRightsActionType(name), entityId }
)

export const getRightsSuccess = name => (rights, entityId) => (
  { type: createGetRightsSuccessActionType(name), rights, entityId }
)

export const getRightsFailure = name => error => (
  { type: createGetRightsFailureActionType(name), error }
)
