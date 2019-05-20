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

const testActionCreatorType = function (actionCreator, type) {
  const actionCreatorType = actionCreator().type
  it(`action creator type: ${actionCreatorType} equals predefined type: ${type}`, function () {
    expect(actionCreatorType).toEqual(type)
  })
}

const testActionCreatorTypeFormat = function (actionCreator, format) {
  const actionCreatorType = actionCreator().type
  const matches = /(.*)_(REQUEST|SUCCESS|FAILURE)/.exec(actionCreatorType)

  it(`${actionCreatorType} should have \`${format}\` format`, function () {
    expect(matches[2]).toEqual(format)
  })
}

export const testActionCreatorsType = function (creators, types) {
  if (creators.length !== types.length) {
    throw new Error(
      'Every action creator should have exactly one matching type'
    )
  }

  creators.forEach(function (creator, index) {
    const type = types[index]

    testActionCreatorType(creator, type)
  })
}

export const testActionCreatorsTypeByGroups = function (groups) {
  const formats = [ 'REQUEST', 'SUCCESS', 'FAILURE' ]

  groups.forEach(function ({ actionCreators, actionTypes }) {
    testActionCreatorsType(actionCreators, actionTypes)

    if (actionCreators.length > 3) {
      throw new Error('Use `testActionCreatorsType` to test action creators types')
    }

    actionCreators.forEach(function (creator, index) {
      testActionCreatorTypeFormat(creator, formats[index])
    })
  })
}

export const testActionCreatorPayload = function (name, actual, expected) {
  it(`${name}`, function () {
    expect(actual).toEqual(expected)
  })
}

export const testActionCreatorsPayload = function (actionMap) {
  const names = Object.keys(actionMap)

  names.forEach(function (name) {
    const { actual: { type, meta, ...rest }, expected } = actionMap[name]

    testActionCreatorPayload(name, rest, expected)
  })
}
