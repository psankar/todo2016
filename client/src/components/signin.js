import React from 'react';
import {Field, reduxForm, SubmissionError} from 'redux-form';

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

function submit(values) {
    return sleep(1000) // simulate server latency
        .then(() => {
        console.log("What is happening ?");
        if (!['john', 'paul', 'george', 'ringo'].includes(values.username)) {
            throw new SubmissionError({username: 'User does not exist', _error: 'Login failed!'});
        } else if (values.password !== 'redux-form') {
            throw new SubmissionError({password: 'Wrong password', _error: 'Login failed!'});
        } else {
            window.alert(`You submitted:\n\n${JSON.stringify(values, null, 2)}`);
        }
    })
}

const validate = values => {
    const errors = {};
    if (!values.username) {
        errors.username = 'Required';
    } else if (values.username.length <= 3) {
        errors.username = 'Must be more than 3 characters';
    }

    if (!values.password) {
        errors.password = 'Required';
    } else if (values.password.length <= 2) {
        errors.password = 'Must be more than 2 characters';
    }
    return errors;
}

const warn = values => {
    const warnings = {};
    if (values.username === 'guest') {
        warnings.username = 'Hmm, you seem to be using guest account';
    }
    return warnings;
}

const renderField = ({
    input: {
        value,
        onChange
    },
    label,
    type,
    meta: {
        touched,
        error
    }
}) => (
    <div>
        <label>{label}</label>
        <div>
            <input {...value} placeholder={label} type={type} onChange={onChange}/> {touched && error && <span>{error}</span>}
        </div>
    </div>
)

const SignIn = (props) => {
    const {error, handleSubmit, pristine, reset, submitting} = props;
    return (
        <form onSubmit={handleSubmit(submit)}>
            <Field name="username" type="text" component={renderField} label="Username"/>
            <Field
                name="password"
                type="password"
                component={renderField}
                label="Password"/> {error && <strong>{error}</strong>}
            <div>
                <button type="submit" disabled={submitting}>Log In</button>
                <button type="button" disabled={pristine || submitting} onClick={reset}>Clear Values</button>
            </div>
        </form>
    );
}

export default reduxForm({form: 'SignInForm', validate, warn})(SignIn);