import React, { Component, Fragment } from 'react';
import Home from '../components/Home';
import Login from '../components/Login';

class LoginContainer extends Component {
  constructor() {
    super();
    this.state = {
    }
  }

  render() {
    return (
      <Fragment>
        <Home />
        <Login />
      </Fragment>
    )
  }
}

export default LoginContainer;