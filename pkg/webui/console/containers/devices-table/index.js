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

import React from 'react'
import { connect } from 'react-redux'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'

import sharedMessages from '../../../lib/shared-messages'
import Message from '../../../lib/components/message'
import PropTypes from '../../../lib/prop-types'
import FetchTable from '../fetch-table'
import DateTime from '../../../lib/components/date-time'

import { getApplicationId } from '../../../lib/selectors/id'
import {
  getDevices,
  searchDevices,
} from '../../store/actions/devices'
import {
  selectDevices,
  selectDevicesTotalCount,
  selectDevicesFetching,
  selectDevicesError,
} from '../../store/selectors/devices'

const m = defineMessages({
  deviceId: 'Device ID',
  connectedDevices: 'Connected Devices ({deviceCount})',
  add: 'Add Device',
})

const headers = [
  {
    name: 'ids.device_id',
    displayName: m.deviceId,
  },
  {
    name: 'name',
    displayName: sharedMessages.name,
  },
  {
    name: 'created_at',
    displayName: sharedMessages.created,
    render (datetime) {
      return <DateTime.Relative value={datetime} />
    },
  },
]

@bind
class DevicesTable extends React.Component {
  constructor (props) {
    super(props)

    this.searchDevicesList = filters => searchDevices(props.appId, filters)
    this.getDevicesList = filters => getDevices(props.appId, filters)
  }

  baseDataSelector (state) {
    return {
      devices: selectDevices(state),
      totalCount: selectDevicesTotalCount(state),
      fetching: selectDevicesFetching(state),
    }
  }

  render () {
    const { error, totalCount, devicePathPrefix } = this.props

    if (error) {
      throw error
    }

    return (
      <FetchTable
        entity="devices"
        headers={headers}
        addMessage={m.add}
        tableTitle={<Message content={m.connectedDevices} values={{ deviceCount: totalCount }} />}
        getItemsAction={this.getDevicesList}
        searchItemsAction={this.searchDevicesList}
        itemPathPrefix={devicePathPrefix}
        baseDataSelector={this.baseDataSelector}
        {...this.props}
      />
    )
  }
}

DevicesTable.propTypes = {
  devicePathPrefix: PropTypes.string.isRequired,
  totalCount: PropTypes.number,
  error: PropTypes.error,
}

DevicesTable.defaultProps = {
  totalCount: 0,
  error: null,
}

export default connect(function (state) {
  return {
    error: selectDevicesError(state),
    appId: getApplicationId(state.application.application),
    totalCount: selectDevicesTotalCount(state),
  }
})(DevicesTable)
