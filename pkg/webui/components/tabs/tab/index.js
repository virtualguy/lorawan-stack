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
import PropTypes from 'prop-types'
import classnames from 'classnames'
import bind from 'autobind-decorator'
import { NavLink } from 'react-router-dom'

import style from './tab.styl'

class Tab extends React.PureComponent {
  static propTypes = {
    /** A flag specifying whether the tab is active. */
    active: PropTypes.bool,
    children: PropTypes.node,
    className: PropTypes.string,
    /** A flag specifying whether the tab is disabled. */
    disabled: PropTypes.bool,
    exact: PropTypes.bool,
    link: PropTypes.string,
    /** The name of the tab. */
    name: PropTypes.string.isRequired,
    narrow: PropTypes.bool,
    /**
     * A click handler to be called when the selected tab changes. Passes the
     * name of the new active tab as an argument.
     */
    onClick: PropTypes.func,
  }

  static defaultProps = {
    children: undefined,
    className: undefined,
    link: undefined,
    onClick: () => null,
    active: false,
    disabled: false,
    narrow: false,
    exact: true,
  }

  @bind
  handleClick() {
    const { onClick, name, disabled } = this.props

    if (!disabled) {
      onClick(name)
    }
  }

  render() {
    const {
      className,
      onClick,
      name,
      active = false,
      disabled = false,
      narrow,
      children,
      link,
      exact = true,
      ...rest
    } = this.props

    const tabItemClassNames = classnames(className, style.tabItem, {
      [style.tabItemNarrow]: narrow,
      [style.tabItemActive]: !disabled && active,
      [style.tabItemDefault]: !disabled && !active,
      [style.tabItemDisabled]: disabled,
    })

    // There is no support for disabled on anchors in html and hence in
    // `react-router`. So, do not render the link component if the tab is
    // disabled, but render regular tab item instead.
    const canRenderLink = link && !disabled

    const Component = canRenderLink ? NavLink : 'span'
    const props = {
      role: 'button',
      className: tabItemClassNames,
      children,
    }

    if (canRenderLink) {
      props.end = exact
      props.to = link
      props.className = ({ isActive }) =>
        classnames(tabItemClassNames, style.tabItem, {
          [style.tabItemActive]: !disabled && isActive,
        })
    } else {
      props.onClick = this.handleClick
    }

    return (
      <li {...rest} className={style.tab}>
        <Component {...props} children={children} />
      </li>
    )
  }
}

export default Tab
