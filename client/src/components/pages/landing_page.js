import React, {Component} from 'react';

class LandingPage extends Component {
    render() {
        return (
            <div>
                LandingPage {this.props.children}
            </div>
        );
    }
}

export default LandingPage;