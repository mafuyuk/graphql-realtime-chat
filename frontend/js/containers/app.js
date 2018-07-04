import React, { Component, Fragment } from 'react';
import Home from '../components/Home';
import UserList from '../components/UserList';
import MessageForm from '../components/MessageForm';
import MessageList from '../components/MessageList';

class AppContainer extends Component {
  constructor() {
    super();
    this.state = {
      users: [],
      messages: []
    }
  }

  render() {
    return (
      <Fragment>
        <Home />
        <UserList users={this.state.users}/>
        <MessageForm />
        <MessageList messages={this.state.messages} />
      </Fragment>
    )
  }
}

export default AppContainer;