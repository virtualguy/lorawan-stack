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
import { storiesOf } from '@storybook/react'
import { action } from '@storybook/addon-actions'
import { withInfo } from '@storybook/addon-info'
import * as Yup from 'yup'

import Field from '../field'
import FieldGroup from '../field/group'
import Button from '../button'
import SubmitBar from '../submit-bar'
import SubmitButton from '../submit-button'
import Input from '../input'
import NewForm from '../new-form'
import CheckboxGroup from '../checkbox/group'
import Checkbox from '../checkbox'
import RadioGroup from '../radio-button/group'
import Radio from '../radio-button'
import Form from '.'

const handleSubmit = function (data, { resetForm }) {
  action('Submit')(data)
  setTimeout(() => resetForm(data), 1000)
}

const containerStyles = {
  maxWidth: '300px',
}

const containerHorizontalStyles = {
  maxWidth: '600px',
}


storiesOf('Form', module)
  .addDecorator((story, context) => withInfo({
    inline: true,
    header: false,
    source: true,
    propTables: [ Form ],
  })(story)(context))
  .add('Login', () => (
    <div style={containerStyles}>
      <Form
        onSubmit={handleSubmit}
        initialValues={{
          user_id: '',
          password: '',
        }}
        submitEnabledWhenInvalid
      >
        <Field
          title="Username or Email"
          name="user_id"
          type="text"
        />
        <Field
          title="Password"
          name="password"
          type="password"
        />
        <Button type="submit" message="Login" />
        <Button naked message="Create an account" />
      </Form>
    </div>
  ))
  .add('Field Groups', () => (
    <div style={containerHorizontalStyles}>
      <Form
        onSubmit={handleSubmit}
        initialValues={{
          'radio-story': 'foo',
          'checkbox-story': { foo: true },
        }}
        submitEnabledWhenInvalid
        horizontal
      >
        <FieldGroup
          name="radio-story"
          title="Radio Buttons"
          columns
        >
          <Field
            type="radio"
            title="Foo"
            value="foo"
            name="foo"
          />
          <Field
            type="radio"
            title="Bar"
            value="bar"
            name="foo"
          />
          <Field
            type="radio"
            title="Baz"
            value="baz"
            name="foo"
          />
        </FieldGroup>
        <FieldGroup
          name="checkbox-story"
          title="Checkboxes"
          columns
        >
          <Field
            type="checkbox"
            title="Foo"
            name="foo"
            form
          />
          <Field
            type="checkbox"
            title="Bar"
            name="bar"
            form
          />
          <Field
            type="checkbox"
            title="Baz"
            name="baz"
            form
          />
        </FieldGroup>
        <Button type="submit" message="Save" />
      </Form>
    </div>
  ))
  .add('New Form', () => (
    <div style={containerHorizontalStyles}>
      <NewForm
        validateOnBlur
        validateOnChange
        validate
        horizontal
        onSubmit={handleSubmit}
        validationSchema={Yup.object().shape({
          name: Yup.string()
            .min(5, 'Too Short')
            .max(25, 'Too Long')
            .required('Required'),
          description: Yup.string()
            .min(5, 'Too Short')
            .max(50, 'Too Long'),
          checkboxes: Yup.object().test(
            'checkboxes',
            'Cannot be empty',
            values => Object.values(values).reduce((acc, curr) => acc || curr, false)
          ),
        })}
        initialValues={{
          name: '',
          description: '',
          radio: 'radio1',
          checkboxes: {},
        }}
      >
        <NewForm.Field
          component={Input}
          type="text"
          name="name"
          placeholder="Name"
          title="Name"
          required
        />
        <NewForm.Field
          component={Input}
          type="text"
          name="description"
          placeholder="Description"
          title="Description"
          required
        />
        <NewForm.Field
          component={CheckboxGroup}
          name="checkboxes"
          title="Checkboxes"
          description="Choose at least one"
          required
          children={
            [
              <Checkbox name="cb1" title="Checkbox 1" key="cb1" />,
              <Checkbox name="cb2" title="Checkbox 2" key="cb2" />,
              <Checkbox name="cb3" title="Checkbox 3" key="cb3" />,
            ]
          }
        />
        <NewForm.Field
          component={RadioGroup}
          name="radio"
          title="Radio"
          required
          children={
            [
              <Radio title="Radio 1" value="radio1" key="radio1" />,
              <Radio title="Radio 2" value="radio2" key="radio2" />,
              <Radio title="Radio 3" value="radio3" key="radio3" />,
            ]
          }
        />
        <SubmitBar>
          <NewForm.Submit message="Submit" component={SubmitButton} />
        </SubmitBar>
      </NewForm>
    </div>
  ))
