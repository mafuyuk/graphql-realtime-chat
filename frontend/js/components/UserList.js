import React, { Fragment } from 'react';

const UserList = (props) => {
  const users = props.users.map((user) => <p>{user}</p>);
  return (
    <Fragment>
      <h3>Users</h3>
      { users }
    </Fragment>
  );
};
export default UserList;