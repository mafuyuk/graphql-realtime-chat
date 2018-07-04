import React, { Fragment } from 'react';

const Message = (props) => {
  return (
    <Fragment>
      <strong>{props.message.user}</strong>: {props.message.text}}
    </Fragment>
  );
};

export default Message;