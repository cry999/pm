import { NextPage } from 'next';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';

import {
  Button,
  Card,
  CardContent,
  CardHeader,
  createStyles,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  Grid,
  Link as MaterialLink,
  makeStyles,
  TextField,
  Typography,
} from '@material-ui/core';

import AuthPage from '../../src/components/templates/AuthPage';
import { ProjectApi, useFacade } from '../../src/lib/api';
import { KeyType, ResultsItemType } from '../../src/lib/type';

const useStyles = makeStyles((theme) =>
  createStyles({
    card: {
      width: '100%',
      height: 150,
      marginBottom: theme.spacing(1),
    },
  })
);

type P = {};

type Project = ResultsItemType<
  KeyType<ProjectApi, 'listProjectsRelatedWithUser'>
>;

const ProjectList: NextPage<P> = ({}) => {
  const classes = useStyles();

  const api = useFacade(ProjectApi);

  const [projects, setProjects] = useState<Project[]>([]);
  // state for creating new project
  const [open, setOpen] = useState(false);
  const [name, setName] = useState('');
  const [elevatorPitch, setElevatorPitch] = useState('');

  const load = async () => {
    try {
      const newProjects = await api.listProjectsRelatedWithUser();
      setProjects(newProjects.results);
    } catch (e) {
      alert(e);
    }
  };

  const clean = () => {
    setName('');
    setElevatorPitch('');
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleCancel = () => {
    clean();
    setOpen(false);
  };

  const handleCreate = async () => {
    try {
      await api.createProject({ projectForm: { name, elevatorPitch } });
      setOpen(false);
      clean();
      await load();
    } catch (error) {
      alert(error);
    }
  };

  useEffect(() => {
    load();
  }, []);

  return (
    <AuthPage title="Your Projects">
      <Grid item xs={12}>
        <Dialog open={open} onClose={handleClose}>
          <DialogContent>
            <DialogContentText>Create a new project.</DialogContentText>
            <TextField
              id="name"
              label="Name"
              type="text"
              fullWidth
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <TextField
              id="description"
              label="Elevator Pitch"
              type="text"
              fullWidth
              value={elevatorPitch}
              onChange={(e) => setElevatorPitch(e.target.value)}
            />
          </DialogContent>
          <DialogActions>
            <Button variant="outlined" onClick={handleCancel}>
              Cancel
            </Button>
            <Button variant="outlined" color="primary" onClick={handleCreate}>
              Create
            </Button>
          </DialogActions>
        </Dialog>
        <Button
          variant="outlined"
          color="primary"
          onClick={() => setOpen(true)}
        >
          New Project
        </Button>
      </Grid>
      <Grid item xs={12}>
        {projects.map((p) => (
          <Card key={p.id} className={classes.card} variant="outlined">
            <CardHeader
              title={
                <Link passHref href={`/projects/${p.id}`}>
                  <MaterialLink>{p.name}</MaterialLink>
                </Link>
              }
            />
            <CardContent>
              <Typography variant="body2">{p.elevatorPitch}</Typography>
            </CardContent>
          </Card>
        ))}
      </Grid>
    </AuthPage>
  );
};

export default ProjectList;
