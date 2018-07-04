import React, { Component, Fragment } from 'react';
import Home from '../components/Home';
import NotFound from '../components/NotFound';

class NotFoundContainer extends Component {
  constructor() {
    super();
    this.state = {
    }
  }

  render() {
    return (
      <Fragment>
        <Home />
        <NotFound />
      </Fragment>
    )
  }
}

export default NotFoundContainer;