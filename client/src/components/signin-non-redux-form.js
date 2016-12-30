import React, {Component} from 'react';

class SignIn extends Component {
    constructor(props) {
        super(props);

        this.state = {
            username: '',
            password: '',
            error: null
        };

        this.handleSubmit = this
            .handleSubmit
            .bind(this);
    };

    componentWillMount() {}

    renderError() {
        if (this.state.error) {
            return (
                <div>
                    {this.state.error}
                </div>
            );
        }

        return <div></div>;
    }

    handleSubmit(e) {
        e.preventDefault();
        console.log('handleSubmittttttt', this.state.username, this.state.password);
        if (this.state.username.length < 1) {
            this.setState({error: "Invalid username"});
            console.log(this.state.error);
            return
        }

        if (this.state.password.length <= 1) {
            this.setState({error: "Invalid password"});
            return
        }

        // Reseting error
        this.setState({error: null});
        console.log("Emit success event");
    }

    render() {
        // onChange callback gets an event as arg. Need to use event.target.value
        return (
            <form onSubmit={this.handleSubmit}>
                <input
                    placeholder="Username"
                    required
                    value={this.state.username}
                    onChange={(username) => this.setState({username: username.target.value})}/>
                <input
                    placeholder="Password"
                    type="password"
                    required
                    value={this.state.password}
                    onChange={(password) => this.setState({password: password.target.value})}/>
                <button type="submit">Sign In</button>
                {this.renderError()}
            </form>
        )
    }
}

export default SignIn;