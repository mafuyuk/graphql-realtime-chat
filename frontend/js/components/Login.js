import React, {Fragment} from 'react';
import {Input, Button} from 'react-bootstrap';

const Login = () => {
  return (
    <Fragment>
      <Input className="form-control" type="text" label="A" value="a" bsSize="lg" bsStyle="success" />
      <Button className="btn btn-secondary">{'Log in'}</Button>
    </Fragment>
  )
};

export default Login;
