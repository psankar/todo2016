import React, {Component} from 'react';

import SignIn from './signin';
import SignOut from './signout';

class AuthUI extends Component {

    authButton() {
        if (this.props.authenticated) {
            return <SignOut/>;
        }
        return <SignIn/>;
    }

    render() {
        return (
            <div>
                {this.authButton()}
            </div>
        );
    }
}

export default AuthUI;