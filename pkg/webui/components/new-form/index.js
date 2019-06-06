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
import { Formik } from 'formik'
import bind from 'autobind-decorator'

import Notification from '../notification'
import PropTypes from '../../lib/prop-types'
import FormContext from './context'
import FormField from './field'
import FormSubmit from './submit'

const formRenderer = ({ children, ...rest }) => class InnerForm extends React.Component {
  render () {
    const { className, horizontal, error } = rest
    const { handleSubmit, ...restFormikProps } = this.props

    return (
      <form className={className} onSubmit={handleSubmit}>
        {error && <Notification error={error} />}
        <FormContext.Provider value={{
          ...restFormikProps,
          horizontal,
        }}
        >
          {children}
        </FormContext.Provider>
      </form>
    )
  }
}

@bind
class Form extends React.PureComponent {
  render () {
    const {
      onSubmit,
      onReset,
      initialValues,
      isInitialValid,
      validateOnBlur,
      validateOnChange,
      validationSchema,
      ...rest
    } = this.props
    return (
      <Formik
        component={formRenderer(rest)}
        onSubmit={onSubmit}
        onReset={onReset}
        initialValues={initialValues}
        isInitialValid={isInitialValid}
        validateOnBlur={validateOnBlur}
        validateOnChange={validateOnChange}
        validationSchema={validationSchema}
      />
    )
  }
}

Form.propTypes = {
  // formik props
  onSubmit: PropTypes.func.isRequired,
  onReset: PropTypes.func,
  initialValues: PropTypes.object.isRequired,
  validateOnBlur: PropTypes.bool,
  validateOnChange: PropTypes.bool,
  validationSchema: PropTypes.object,
  isInitialValid: PropTypes.bool,
  // custom props
  horizontal: PropTypes.bool,
  className: PropTypes.string,
  error: PropTypes.error,
}

Form.defaultProps = {
  className: null,
  submitEnabledWhenInvalid: false,
  validateOnBlur: true,
  validateOnChange: false,
  validationSchema: {},
  isInitialValid: false,
  onReset: () => null,
  error: '',
  horizontal: true,
}

Form.Field = FormField
Form.Submit = FormSubmit

export default Form
