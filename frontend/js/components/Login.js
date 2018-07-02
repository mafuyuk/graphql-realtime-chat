import React, {Fragment} from 'react';

const Login = () => {
  return (
    <Fragment>
      <input  type="text" className="form-control" placeholder="Username..." required />
      <button className="btn btn-secondary" type="submit">Log in</button>
    </Fragment>
  )
}

export default Login;
