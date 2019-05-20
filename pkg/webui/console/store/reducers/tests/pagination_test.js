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

import createNamedPaginationReducer,
{ createNamedPaginationReducerById } from '../pagination'
import {
  getPagination,
  getPaginationSuccess,
  getPaginationFailure,
} from '../../actions/pagination'

describe('pagination reducers', function () {
  const NAME = 'ENTITY'
  const getEntityPagination = getPagination(NAME)
  const getEntityPaginationSuccess = getPaginationSuccess(NAME)
  const getEntityPaginationFailure = getPaginationFailure(NAME)

  describe('flat', function () {
    const reducer = createNamedPaginationReducer(NAME)
    const defaultState = { ids: [], totalCount: 0 }

    it('should return the initial state', function () {
      expect(reducer(undefined, {})).toEqual(defaultState)
    })

    it('should ignore the `request` action', function () {
      expect(reducer(defaultState, getEntityPagination())).toEqual(defaultState)
    })

    it('should ignore the `failure` action', function () {
      expect(reducer(defaultState, getEntityPaginationFailure())).toEqual(defaultState)
    })

    describe('receives the `success` action', function () {
      const entities = [{ id: '1' }, { id: '2' }]
      const totalCount = entities.length
      const action = getEntityPaginationSuccess(entities, totalCount)

      let newState = null

      beforeAll(function () {
        newState = reducer(defaultState, action)
      })

      it('should update the state', function () {
        expect(newState).not.toEqual(defaultState)
      })

      it('should only store ids', function () {
        const { ids } = newState

        expect(ids).toEqual(entities.map(e => e.id))
      })

      it('should store `totalCount`', function () {
        const { totalCount: newTotalCount } = newState

        expect(newTotalCount).toEqual(totalCount)
      })
    })
  })

  describe('by id', function () {
    const reducer = createNamedPaginationReducerById(NAME)
    const defaultState = {}

    it('should return the initial state', function () {
      expect(reducer(undefined, {})).toEqual(defaultState)
    })

    it('should ignore the `request` action', function () {
      expect(reducer(defaultState, getEntityPagination())).toEqual(defaultState)
    })

    it('should ignore the `failure` action', function () {
      expect(reducer(defaultState, getEntityPaginationFailure())).toEqual(defaultState)
    })

    describe('receives the `success` action', function () {
      const entityId = 'entity-id'
      const entities = [{ id: '1' }, { id: '2' }]
      const totalCount = entities.length
      const action = getEntityPaginationSuccess(entities, totalCount, entityId)

      let newState = null

      beforeAll(function () {
        newState = reducer(defaultState, action)
      })

      it('should ignore without `entityId` in the payload', function () {
        const newState = reducer(defaultState, getEntityPaginationSuccess(entities, totalCount))

        expect(newState).toEqual(defaultState)
      })

      it('should update the state', function () {
        expect(newState).not.toEqual(defaultState)
      })

      it('should store results per entity id', function () {
        const { [entityId]: results } = newState

        expect(results).not.toBeUndefined()
      })

      it('should only store ids', function () {
        const { [entityId]: results } = newState

        expect(results.ids).toEqual(entities.map(e => e.id))
      })

      it('should store `totalCount`', function () {
        const { [entityId]: results } = newState

        expect(results.totalCount).toEqual(totalCount)
      })

      describe('receives the `success` action for another entity', function () {
        const otherEntityId = 'other-entity-id'
        const otherEntities = [{ id: '3' }, { id: '4' }, { id: '5' }]
        const otherTotalCount = otherEntities.length
        const action = getEntityPaginationSuccess(otherEntities, otherTotalCount, otherEntityId)

        let otherNewState = null

        beforeAll(function () {
          otherNewState = reducer(newState, action)
        })

        it('should update the state', function () {
          expect(otherNewState).not.toEqual(newState)
        })

        it('should preserve previous entries', function () {
          const { [entityId]: results } = otherNewState

          expect(results).not.toBeUndefined()
          expect(results.ids).toEqual(entities.map(e => e.id))
          expect(results.totalCount).toEqual(totalCount)
        })

        it('should store results per entity id', function () {
          const { [otherEntityId]: results } = otherNewState

          expect(results).not.toBeUndefined()
        })

        it('should only store ids', function () {
          const { [otherEntityId]: results } = otherNewState

          expect(results.ids).toEqual(otherEntities.map(e => e.id))
        })

        it('should store `totalCount`', function () {
          const { [otherEntityId]: results } = otherNewState

          expect(results.totalCount).toEqual(otherTotalCount)
        })
      })
    })
  })
})
