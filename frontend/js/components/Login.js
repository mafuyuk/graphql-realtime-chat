import React from 'react';
import FormControl from '@material-ui/core/FormControl';
import Button from '@material-ui/core/Button';
import Input from '@material-ui/core/Input';

const Login = () => {
  return (
    <FormControl>
      <Input type="text" />
      <Button variant="raised" color="primary">
        Log in
      </Button>
    </FormControl>
  )
};

export default Login;
