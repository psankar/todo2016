import React, {Component} from 'react';

import Login from './components/login';
import NewTask from './components/new_task';
import TaskList from './components/task_list';

import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo"/>
          <Login/>
        </div>
        <p className="App-intro">
          <TaskList/>
          <NewTask/>
        </p>
      </div>
    );
  }
}

export default App;
