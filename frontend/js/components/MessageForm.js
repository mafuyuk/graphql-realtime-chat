import React from 'react';

const MessageForm = () => {
  return (
    <form className="col-12">
      <div className="input-group">
        <input type="text" className="form-control" placeholder="Message..." />
        <div className="input-group-append">
          <button className="btn btn-outline-secondary" type="submit">Post</button>
        </div>
      </div>
    </form>
  );
};

export default MessageForm;