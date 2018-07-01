import React, {Fragment} from 'react';

const Login = () => {
  return (
    <Fragment>
      <input  type="text" class="form-control" placeholder="Username..." required >
      <button class="btn btn-secondary" type="submit">Log in</button>
    </Fragment>
  )
}

export default Login;
