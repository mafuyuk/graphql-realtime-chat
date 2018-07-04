import React, {Fragment} from 'react';
import Button from '@material-ui/core/Button';
import Input from '@material-ui/core/Input';

const Login = () => {
  return (
    <Fragment>
      <Input type="text" />
      <Button variant="raised" color="primary">
        Log in
      </Button>
    </Fragment>
  )
};

export default Login;
