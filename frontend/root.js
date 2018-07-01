import React, { Fragment} from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route } from 'react-router-dom'

import Login from './js/components/Login';

import './css/index.css';

const Root = () => {
  return (
    <Fragment>
      <p>Hello</p>
    </Fragment>
);

};
const rootElement = document.getElementById("root");
ReactDOM.render((
  <BrowserRouter>
    <Route path="/" component={Root} />
    <Route path="/login" component={Login} />
  </BrowserRouter>
), rootElement);