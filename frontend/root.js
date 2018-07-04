import React, { Fragment } from "react";
import ReactDOM from "react-dom";
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom'

import AppContainer from './js/containers/app';
import LoginContainer from './js/containers/login';
import NotFoundContainer from './js/containers/NotFound';

import './css/index.css';


const Rooting = () => {
  return (
    <Fragment>
      <Switch>
        <Route exact path="/" component={AppContainer} />
        <Route path="/login" component={LoginContainer} />
        <Route component={NotFoundContainer} />
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