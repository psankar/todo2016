import React, {Component} from 'react';
import {createStore, applyMiddleware} from 'redux';
import {Provider} from 'react-redux';
import {Router, Route, browserHistory} from 'react-router';

import Login from './components/login';
import LandingPage from './components/pages/landing_page';
import HomePage from './components/pages/home_page';
import rootReducer from './reducers/index';
import requireAuth from './components/hoc/require_authentication';

import logo from './logo.svg';
import './App.css';

const createStoreWithMiddleware = applyMiddleware()(createStore);

class App extends Component {
  render() {
    return (
      <Provider store={createStoreWithMiddleware(rootReducer)}>
        <div className="App">
          <div className="App-header">
            <img src={logo} className="App-logo" alt="logo"/>
            <Login/>
          </div>
          <div className="App-intro">
            <Router history={browserHistory}>
              <Route path="/" component={LandingPage}>
                <Route path="home" component={requireAuth(HomePage)}/>
              </Route>
            </Router>
          </div>
        </div>
      </Provider>
    );
  }
}

export default App;