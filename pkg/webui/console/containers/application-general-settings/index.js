// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

import React, { useCallback, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import { defineMessages } from 'react-intl'

import toast from '@ttn-lw/components/toast'

import ApplicationGeneralSettingsForm from '@console/components/application-general-settings-form'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import PropTypes from '@ttn-lw/lib/prop-types'
import diff from '@ttn-lw/lib/diff'
import { selectCollaboratorsTotalCount } from '@ttn-lw/lib/store/selectors/collaborators'

import {
  checkFromState,
  mayDeleteApplication,
  mayViewOrEditApplicationApiKeys,
  mayViewOrEditApplicationCollaborators,
  mayPurgeEntities,
  mayViewApplicationLink,
} from '@console/lib/feature-checks'

import { updateApplicationLink } from '@console/store/actions/link'
import {
  deleteAppPkgDefaultAssoc,
  setAppPkgDefaultAssoc,
} from '@console/store/actions/application-packages'
import { updateApplication, deleteApplication } from '@console/store/actions/applications'

import { selectApplicationPackageDefaultAssociation } from '@console/store/selectors/application-packages'
import { selectWebhooksTotalCount } from '@console/store/selectors/webhooks'
import { selectPubsubsTotalCount } from '@console/store/selectors/pubsubs'
import { selectApiKeysTotalCount } from '@console/store/selectors/api-keys'
import {
  selectApplicationLink,
  selectSelectedApplication,
} from '@console/store/selectors/applications'

const promisifiedSetAppPkgDefaultAssoc = attachPromise(setAppPkgDefaultAssoc)
const promisifiedDeleteAppPkgDefaultAssoc = attachPromise(deleteAppPkgDefaultAssoc)
const promisifiedUpdateApplicationLink = attachPromise(updateApplicationLink)
const promisifiedUpdateApplication = attachPromise(updateApplication)
const alcsyncPackageName = 'alcsync-v1'

const m = defineMessages({
  updateSuccess: 'Application updated',
  deleteSuccess: 'Application deleted',
})

const ApplicationGeneralSettingsContainer = ({ appId }) => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const [error, setError] = useState()
  const application = useSelector(selectSelectedApplication)
  const link = useSelector(selectApplicationLink)
  const apiKeysCount = useSelector(selectApiKeysTotalCount)
  const collaboratorsCount = useSelector(selectCollaboratorsTotalCount)
  const webhooksCount = useSelector(selectWebhooksTotalCount)
  const pubsubsCount = useSelector(selectPubsubsTotalCount)
  const mayViewApiKeys = useSelector(state =>
    checkFromState(mayViewOrEditApplicationApiKeys, state),
  )
  const mayViewCollaborators = useSelector(state =>
    checkFromState(mayViewOrEditApplicationCollaborators, state),
  )
  const mayPurgeApp = useSelector(state => checkFromState(mayPurgeEntities, state))
  const mayViewLink = useSelector(state => checkFromState(mayViewApplicationLink, state))
  const hasIntegrations = webhooksCount > 0 || pubsubsCount > 0
  const hasApiKeys = apiKeysCount > 0
  // Note: there is always at least one collaborator.
  const hasAddedCollaborators = collaboratorsCount > 1
  const isPristine = !hasApiKeys && !hasAddedCollaborators && !hasIntegrations
  const shouldConfirmDelete =
    !isPristine || !mayViewCollaborators || !mayViewApiKeys || Boolean(error)
  const packageAssoc = useSelector(state => selectApplicationPackageDefaultAssociation(state, 202))
  const alcsync =
    packageAssoc?.package_name === alcsyncPackageName ? { alcsync: true } : { alcsync: false }

  // Add technical and administrative contact to the initial values.
  const { administrative_contact, technical_contact, ...applicationValues } = application
  const technicalContact =
    application.technical_contact !== undefined && application.technical_contact !== null
      ? {
          _technical_contact_id: application.technical_contact.user_ids
            ? application.technical_contact.user_ids.user_id
            : application.technical_contact.organization_ids.organization_id,
          _technical_contact_type: application.technical_contact.user_ids ? 'user' : 'organization',
        }
      : {
          _technical_contact_id: '',
          _technical_contact_type: '',
        }
  const administrativeContact =
    application.administrative_contact !== undefined && application.administrative_contact !== null
      ? {
          _administrative_contact_id: application.administrative_contact.user_ids
            ? application.administrative_contact.user_ids.user_id
            : application.administrative_contact.organization_ids.organization_id,
          _administrative_contact_type: application.administrative_contact.user_ids
            ? 'user'
            : 'organization',
        }
      : {
          _administrative_contact_id: '',
          _administrative_contact_type: '',
        }
  const initialValues = {
    ...technicalContact,
    ...administrativeContact,
    ...applicationValues,
    ...link,
    ...alcsync,
  }

  const handleAlcsyncUpdate = useCallback(
    async (appId, alcsync) => {
      if (alcsync) {
        return await dispatch(
          promisifiedSetAppPkgDefaultAssoc(appId, 202, {
            package_name: alcsyncPackageName,
          }),
        )
      }

      return await dispatch(
        promisifiedDeleteAppPkgDefaultAssoc(appId, 202, {
          package_name: alcsyncPackageName,
        }),
      )
    },
    [dispatch],
  )

  const handleSubmit = useCallback(
    async (values, { resetForm, setSubmitting }) => {
      setError(undefined)

      const {
        _administrative_contact_id,
        _administrative_contact_type,
        _technical_contact_id,
        _technical_contact_type,
      } = values

      const administrative_contact =
        _administrative_contact_id !== ''
          ? {
              [`${_administrative_contact_type}_ids`]: {
                [`${_administrative_contact_type}_id`]: _administrative_contact_id,
              },
            }
          : ''

      const technical_contact =
        _technical_contact_id !== ''
          ? {
              [`${_technical_contact_type}_ids`]: {
                [`${_technical_contact_type}_id`]: _technical_contact_id,
              },
            }
          : ''

      const changed = diff(
        application,
        { administrative_contact, technical_contact, ...values },
        {
          exclude: [
            '_administrative_contact_id',
            '_administrative_contact_type',
            '_technical_contact_id',
            '_technical_contact_type',
          ],
        },
      )

      // If there is a change in attributes, copy all attributes so they don't get
      // overwritten.
      const update =
        'attributes' in changed
          ? {
              ...changed,
              attributes: values.attributes,
            }
          : changed

      if (technical_contact === '') {
        update.technical_contact = null
      }
      if (administrative_contact === '') {
        update.administrative_contact = null
      }

      const {
        ids: { application_id },
      } = application

      try {
        const { skip_payload_crypto, alcsync, ...applicationUpdate } = update
        const linkUpdate = { skip_payload_crypto }
        await dispatch(promisifiedUpdateApplication(application_id, applicationUpdate))
        await dispatch(promisifiedUpdateApplicationLink(application_id, linkUpdate))
        if ('alcsync' in update) {
          await handleAlcsyncUpdate(application_id, alcsync)
        }
        resetForm({ values })
        toast({
          title: application_id,
          message: m.updateSuccess,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setSubmitting(false)
        setError(error)
      }
    },
    [application, handleAlcsyncUpdate, dispatch],
  )

  const onDeleteSuccess = useCallback(() => navigate(`/applications`), [navigate])
  const handleDelete = useCallback(
    async shouldPurge => {
      setError(undefined)

      try {
        await dispatch(attachPromise(deleteApplication(appId, shouldPurge)))
        toast({
          title: appId,
          message: m.deleteSuccess,
          type: toast.types.SUCCESS,
        })
        onDeleteSuccess()
      } catch (error) {
        setError(error)
      }
    },
    [appId, onDeleteSuccess, dispatch],
  )

  return (
    <ApplicationGeneralSettingsForm
      initialValues={initialValues}
      handleSubmit={handleSubmit}
      handleDelete={handleDelete}
      error={error}
      mayViewApplicationLink={mayViewLink}
      mayDeleteApplication={mayDeleteApplication}
      appId={appId}
      applicationName={application.name}
      shouldConfirmDelete={shouldConfirmDelete}
      mayPurge={mayPurgeApp}
    />
  )
}

ApplicationGeneralSettingsContainer.propTypes = {
  appId: PropTypes.string.isRequired,
}

export default ApplicationGeneralSettingsContainer
