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
import bind from 'autobind-decorator'
import { Container, Col, Row } from 'react-grid-system'
import { replace } from 'connected-react-router'

import { withBreadcrumb } from '../../../components/breadcrumbs/context'
import Breadcrumb from '../../../components/breadcrumbs/breadcrumb'
import sharedMessages from '../../../lib/shared-messages'
import Spinner from '../../../components/spinner'
import Message from '../../../lib/components/message'
import IntlHelmet from '../../../lib/components/intl-helmet'
import { ApiKeyEditForm } from '../../../components/api-key-form'

import {
  getApplicationApiKey,
  getApplicationRights,
  GET_APP_API_KEY_BASE,
  GET_APP_RIGHTS_LIST_BASE,
} from '../../store/actions/applications'
import {
  selectApplicationRightsById,
} from '../../store/selectors/applications'
import {
  selectApiKeyById,
} from '../../store/selectors/api-keys'
import {
  createFetchingSelector,
} from '../../store/selectors/fetching'
import {
  createErrorSelector,
} from '../../store/selectors/error'

import api from '../../api'

const selectPageFetching = createFetchingSelector([
  GET_APP_API_KEY_BASE,
  GET_APP_RIGHTS_LIST_BASE,
])
const selectPageError = createErrorSelector([
  GET_APP_API_KEY_BASE,
  GET_APP_RIGHTS_LIST_BASE,
])

@connect(function (state, props) {
  const { appId, apiKeyId } = props.match.params

  return {
    apiKeyId,
    appId,
    apiKey: selectApiKeyById(state, apiKeyId),
    rights: selectApplicationRightsById(state, appId),
    fetching: selectPageFetching(state),
    error: selectPageError(state),
  }
},
dispatch => ({
  loadPageData (appId, apiKeyId) {
    dispatch(getApplicationApiKey(appId, apiKeyId))
    dispatch(getApplicationRights(appId))
  },
}))
@withBreadcrumb('apps.single.api-keys.edit', function (props) {
  const { appId, apiKeyId } = props

  return (
    <Breadcrumb
      path={`/console/applications/${appId}/api-keys/${apiKeyId}/edit`}
      icon="general_settings"
      content={sharedMessages.edit}
    />
  )
})
@bind
export default class ApplicationApiKeyEdit extends React.Component {

  constructor (props) {
    super(props)

    this.deleteApplicationKey = id => api.application.apiKeys.delete(props.appId, id)
    this.editApplicationKey = key => api.application.apiKeys.update(
      props.appId,
      props.apiKey.id,
      key
    )
  }

  componentDidMount () {
    const { loadPageData, appId, apiKeyId } = this.props

    loadPageData(appId, apiKeyId)
  }

  onDeleteSuccess () {
    const { appId, dispatch } = this.props

    dispatch(replace(`/console/applications/${appId}/api-keys`))
  }

  render () {
    const { apiKey, rights, fetching, error } = this.props

    if (error) {
      throw error
    }

    if (fetching || !apiKey) {
      return <Spinner center />
    }

    return (
      <Container>
        <Row>
          <Col lg={8} md={12}>
            <IntlHelmet title={sharedMessages.keyEdit} />
            <Message component="h2" content={sharedMessages.keyEdit} />
          </Col>
        </Row>
        <Row>
          <Col lg={8} md={12}>
            <ApiKeyEditForm
              rights={rights}
              apiKey={apiKey}
              onEdit={this.editApplicationKey}
              onDelete={this.deleteApplicationKey}
              onDeleteSuccess={this.onDeleteSuccess}
            />
          </Col>
        </Row>
      </Container>
    )
  }
}
