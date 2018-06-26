import React, { Fragment} from "react";
import ReactDOM from "react-dom";

import './css/index.css';

const Root = () => {
  return (
    <Fragment>
      <p>Hello</p>
    </Fragment>
);

};
const rootElement = document.getElementById("root");
ReactDOM.render(<Root />, rootElement);
