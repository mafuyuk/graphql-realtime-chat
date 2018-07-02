import React, { Fragment } from "react";
import ReactDOM from "react-dom";
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom'

import Login from './js/components/Login';
import MessageForm from './js/components/MessageForm';

import './css/index.css';

const Home = () => {
  return (
    <Fragment>
      <h2>Home</h2>
    </Fragment>
  );
};

const Rooting = () => {
  return (
    <Fragment>
      <Home />
      <Route exact path="/" component={MessageForm} />
      <Route path="/login" component={Login} />
    </Fragment>
  );
};

const rootElement = document.getElementById("root");
ReactDOM.render((
  <Router>
    <Rooting />
  </Router>
), rootElement);