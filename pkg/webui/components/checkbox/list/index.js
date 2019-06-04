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
import classnames from 'classnames'
import bind from 'autobind-decorator'
import { defineMessages } from 'react-intl'

import Checkbox from '..'
import CheckboxGroup from '../group'
import Message from '../../../lib/components/message'
import PropTypes from '../../../lib/prop-types'

import style from './list.styl'

const m = defineMessages({
  checkAll: 'Check All',
  selections: 'Your Selections',
})

@bind
class CheckboxList extends React.Component {

  state = {
    selected: [],
    allSelected: false,
  }

  handleCheckedChange (name, value) {
    const { onChange } = this.props

    if (value) {
      this.setState(prev => ({
        selected: [ name, ...prev.selected ],
      }))
    } else {
      this.setState(prev => ({
        selected: prev.selected.filter(n => n !== name),
      }))
    }

    onChange(name, value)
  }

  handleTransferCheckedChange (name) {
    this.handleCheckedChange(name, false)
  }

  get header () {
    const { transfer, selectAllTitle, children } = this.props
    const { selected, allSelected } = this.state

    const cls = classnames(style.header, {
      [style.headerWithTransfer]: transfer,
    })

    const selectedCount = selected.length
    const totalCount = React.Children.count(children)

    return (
      <div className={cls}>
        <div>
          <Checkbox
            title={selectAllTitle}
            name="select-all"
            value={allSelected}
          />
        </div>
        <span className={style.headerCountLabel}>
          {`(${selectedCount}/${totalCount})`}
        </span>
        { transfer && <Message content={m.selections} />}
      </div>
    )
  }

  get all () {
    const {
      children,
      name,
      value,
    } = this.props

    return (
      <div className={style.bodyContentAll}>
        <CheckboxGroup
          value={value}
          onChange={this.handleCheckedChange}
          name={name}
          horizontal={false}
        >
          {children}
        </CheckboxGroup>
      </div>
    )
  }

  get transfer () {
    const {
      transfer,
      children,
    } = this.props

    if (!transfer) {
      return null
    }

    const { selected } = this.state

    return (
      <div className={style.bodyContentSelected}>
        {React.Children.toArray(children)
          .filter(Child => selected.includes(Child.props.name))
          .map(Child => (
            React.cloneElement(Child, {
              ...Child.props,
              value: true,
              onChange: this.handleTransferCheckedChange,
            })
          ))}
      </div>
    )
  }

  render () {
    const {
      className,
      transfer,
    } = this.props

    const bodyClassNames = classnames(style.body, {
      [style.bodyWithTransfer]: transfer,
    })

    return (
      <div className={classnames(className, style.checkboxList)}>
        {this.header}
        <div className={bodyClassNames}>
          {this.all}
          {this.transfer}
        </div>
      </div>
    )
  }
}

CheckboxList.propTypes = {
  transfer: PropTypes.bool,
  name: PropTypes.string.isRequired,
  selectAllTitle: PropTypes.message,
  onChange: PropTypes.func,
}

CheckboxList.defaultProps = {
  selectAllTitle: m.checkAll,
  onChange: () => null,
}

export default CheckboxList
