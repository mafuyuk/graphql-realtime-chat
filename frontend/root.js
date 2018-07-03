import React, { Fragment } from "react";
import ReactDOM from "react-dom";
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import Login from './js/components/Login';
import MessageForm from './js/components/MessageForm';
import NotFound from './js/components/NotFound';

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
      <Switch>
        <Route exact path="/" component={MessageForm} />
        <Route path="/login" component={Login} />
        <Route component={NotFound} />
      </Switch>
    </Fragment>
  );
};

const rootElement = document.getElementById("root");
ReactDOM.render((
  <Router>
    <Rooting />
  </Router>
), rootElement);