import React, { useState } from 'react';

import {
  Button,
  createStyles,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  IconButton,
  makeStyles,
  TextField,
  Typography,
} from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';

const useStyles = makeStyles(() =>
  createStyles({
    addIcon: {
      width: '100%',
      height: '100%',
      borderRadius: 5,
    },
  })
);

export type NewTaskForm = { name: string; description: string };

type P = {
  onCreate?: (form: NewTaskForm) => any;
};

const FormDialog: React.FC<P> = ({ onCreate = () => {} }) => {
  const classNames = useStyles();
  const [open, setOpen] = useState(false);
  const [name, setName] = useState('');
  const [desc, setDesc] = useState('');

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleCancel = () => {
    handleClose();
  };

  const handleCreate = () => {
    onCreate({ name, description: desc });
    handleClose();
  };

  return (
    <div>
      <IconButton
        className={classNames.addIcon}
        color="primary"
        onClick={handleClickOpen}
      >
        <AddIcon fontSize="large" />
        <Typography variant="body2">Create a new Task</Typography>
      </IconButton>
      <Dialog open={open} onClose={handleClose}>
        <DialogContent>
          <DialogContentText>Create a new Task.</DialogContentText>
          <TextField
            margin="dense"
            id="name"
            label="Name"
            type="text"
            fullWidth
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <TextField
            margin="dense"
            id="description"
            label="Description"
            type="text"
            multiline
            fullWidth
            value={desc}
            onChange={(e) => setDesc(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCancel}>Cancel</Button>
          <Button onClick={handleCreate} color="primary">
            Create
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default FormDialog;
