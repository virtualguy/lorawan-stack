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
import bind from 'autobind-decorator'
import classnames from 'classnames'
import { getIn } from 'formik'

import from from '../../../lib/from'
import Icon from '../../icon'
import Message from '../../../lib/components/message'
import FormContext from '../context'
import PropTypes from '../../../lib/prop-types'

import style from './field.styl'

export function getPassThroughProps (props, excludeProps) {
  const rest = {}
  for (const property of Object.keys(props)) {
    if (!excludeProps[property]) {
      rest[property] = props[property]
    }
  }
  return rest
}

@bind
class FormField extends React.Component {

  static contextType = FormContext
  static propTypes = {
    className: PropTypes.string,
    name: PropTypes.string.isRequired,
    warning: PropTypes.message,
    required: PropTypes.bool,
    description: PropTypes.message,
    component: PropTypes.oneOfType([
      PropTypes.func,
      PropTypes.string,
    ]).isRequired,
  }

  componentDidMount () {
    const { name } = this.props

    this.context.registerField(name, this)
  }

  componentWillUnmount () {
    const { name } = this.props

    this.context.unregisterField(name)
  }

  handleChange (value) {
    const { name } = this.props

    this.context.setFieldValue(name, value)
  }

  handleBlur (event) {
    const { name } = this.props

    this.context.handleBlur(event)
    if (this.context.validateOnBlur) {
      this.context.setFieldTouched(name, true)
    }
  }

  render () {
    const {
      className,
      name,
      title,
      warning,
      description,
      disabled,
      required,
      readOnly,
      component: Component,
    } = this.props
    const { horizontal } = this.context

    const fieldValue = getIn(this.context.values, name)
    const fieldError = getIn(this.context.errors, name)
    const fieldTouched = getIn(this.context.touched, name)

    const hasError = Boolean(fieldError)
    const hasWarning = Boolean(warning)
    const hasDescription = Boolean(description)

    const showError = fieldTouched && hasError
    const showWarning = !hasError && hasWarning
    const showDescription = !showError && !showWarning && hasDescription

    const fieldMessage = showError ? (
      <div className={style.messages}>
        <Err error={fieldError} />
      </div>
    ) : showWarning ? (
      <div className={style.messages}>
        <Err warning={warning} />
      </div>
    ) : showDescription ? (
      <Message className={style.description} content={description} />
    ) : null

    const fieldComponentProps = {
      value: fieldValue,
      name,
      disabled,
      horizontal,
      onChange: this.handleChange,
      onBlur: this.handleBlur,
    }

    const cls = classnames(className, style.field, from(style, {
      error: showError,
      warning: showWarning,
      horizontal,
      required,
      readOnly,
      disabled,
    }))

    return (
      <div className={cls}>
        <label className={style.label}>
          <Message content={title} className={style.title} />
          <span className={style.reqicon}>&middot;</span>
        </label>
        <Component
          className={style.component}
          {...fieldComponentProps}
          {...getPassThroughProps(this.props, FormField.propTypes)}
        />
        {fieldMessage}
      </div>
    )
  }
}

const Err = function (props) {
  const {
    error,
    warning,
    name,
    className,
  } = props

  const content = error || warning || ''
  const contentValues = content.values || {}

  const icon = error ? 'error' : 'warning'

  const classname = classnames(style.message, className, {
    [style.show]: content && content !== '',
    [style.hide]: !content || content === '',
    [style.err]: error,
    [style.warn]: warning,
  })

  return (
    <div className={classname}>
      <Icon icon={icon} className={style.icon} />
      <Message
        content={content.format || content.error_description || content.message || content}
        values={{
          ...contentValues,
          name: <Message content={name} className={style.name} />,
        }}
      />
    </div>
  )
}

export default FormField
