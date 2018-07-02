import React from 'react';

const MessageForm = () => {
  return (
    <form class="col-12">
      <div class="input-group">
        <input type="text" class="form-control" placeholder="Message..." >
        <div class="input-group-append">
          <button class="btn btn-outline-secondary" type="submit">Post</button>
        </div>
      </div>
    </form>
  );
};

export default MessageForm;