import React from 'react';
import FormControl from '@material-ui/core/FormControl';
import Button from '@material-ui/core/Button';
import Input from '@material-ui/core/Input';

const MessageForm = () => {
  return (
    <FormControl>
      <Input type="text" placeholder="Message..." />
      <Button type="submit" >
        Post
      </Button>
    </FormControl>
  );
};

export default MessageForm;