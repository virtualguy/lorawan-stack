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

import React, { Component } from 'react'
import { connect } from 'react-redux'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'

import sharedMessages from '../../../lib/shared-messages'
import PropTypes from '../../../lib/prop-types'
import FetchTable from '../fetch-table'

import {
  getApplications,
} from '../../store/actions/applications'
import {
  selectApplications,
  selectApplicationsTotalCount,
  selectApplicationsFetching,
  selectApplicationsError,
} from '../../store/selectors/applications'

const m = defineMessages({
  all: 'All',
  appId: 'Application ID',
  desc: 'Description',
  empty: 'No items matched your criteria',
})

const tabs = [
  {
    title: m.all,
    name: 'all',
    disabled: true,
  },
]

const headers = [
  {
    name: 'ids.application_id',
    displayName: m.appId,
  },
  {
    name: 'description',
    displayName: m.desc,
  },
]

@bind
class ApplicationsTable extends Component {

  baseDataSelector (state) {
    return {
      applications: selectApplications(state),
      totalCount: selectApplicationsTotalCount(state),
      fetching: selectApplicationsFetching(state),
    }
  }

  render () {
    const { error, ...rest } = this.props

    if (error) {
      throw error
    }

    return (
      <FetchTable
        entity="applications"
        headers={headers}
        addMessage={sharedMessages.addApplication}
        getItemsAction={getApplications}
        tabs={tabs}
        baseDataSelector={this.baseDataSelector}
        {...rest}
      />
    )
  }
}

ApplicationsTable.propTypes = {
  error: PropTypes.error,
}

ApplicationsTable.defaultProps = {
  error: null,
}

export default connect(state => ({
  error: selectApplicationsError(state),
}))(ApplicationsTable)
