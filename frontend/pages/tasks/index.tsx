import { NextPage } from 'next';
import React, { useEffect, useState } from 'react';

import {
  Card,
  createStyles,
  Grid,
  IconButton,
  makeStyles,
  Typography,
} from '@material-ui/core';
import CachedTwoToneIcon from '@material-ui/icons/CachedTwoTone';

import NewTaskDialog, { NewTaskForm } from '../../src/components/organisms/NewTaskDialog';
import TaskList from '../../src/components/organisms/TaskList';
import AuthPage from '../../src/components/templates/AuthPage';
import { TaskApi, useFacade } from '../../src/lib/api';
import { KeyType, ResultsItemType } from '../../src/lib/type';

const useStyles = makeStyles((theme) =>
  createStyles({
    reloadIcon: {
      width: 50,
      height: 50,
      margin: theme.spacing(1),
      float: 'right',
    },
    title: {
      textAlign: 'center',
      float: 'left',
    },
    newCard: {
      borderStyle: 'dashed',
      padding: theme.spacing(2),
    },
    card: {
      margin: theme.spacing(2),
    },
    cardHeader: {},
    cardContent: {},
    cardStatus: {},
  })
);

type P = {};

type Task = ResultsItemType<KeyType<TaskApi, 'listAllAssociatedWithUser'>>;

const TaskDashboard: NextPage<P> = ({}) => {
  const classNames = useStyles();
  const api = useFacade(TaskApi);
  const [tasks, setTasks] = useState<Task[]>([]);

  const load = async () => {
    try {
      const taskList = await api.listAllAssociatedWithUser();
      setTasks(taskList.results);
    } catch (e) {
      alert(e);
      console.error(e);
    }
  };

  useEffect(() => {
    load();
  }, []);

  const handleCompletedStep = async (
    taskId: string,
    step: string
  ): Promise<boolean> => {
    try {
      switch (step) {
        case 'TODO':
          const nextTODO = await api.progressTask({ taskId });
          setTasks(tasks.map((t) => (t.id === nextTODO.id ? nextTODO : t)));
          return true;
        case 'WIP':
          const nextWIP = await api.doneTask({ taskId });
          setTasks(tasks.map((t) => (t.id === nextWIP.id ? nextWIP : t)));
          return true;
        case 'DONE':
          return true;
        default:
          return false;
      }
    } catch (e) {
      alert(e);
      console.error(e);
      return false;
    }
  };

  const handleReload = async () => {
    await load();
  };

  const handleCreate = async ({ name, description }: NewTaskForm) => {
    try {
      await api.createTask({ taskForm: { name, description } });
      await load();
    } catch (e) {
      alert(e);
    }
  };

  const handleAssign = async (taskId: string) => {
    try {
      await api.assignSignedInUserToTask({ taskId });
      await load();
    } catch (e) {
      alert(e);
    }
  };

  return (
    <AuthPage
      title={
        <Grid item xs={12}>
          <Typography variant="h3" className={classNames.title}>
            Task Dashboard
          </Typography>
          <IconButton className={classNames.reloadIcon} onClick={handleReload}>
            <CachedTwoToneIcon />
          </IconButton>
        </Grid>
      }
    >
      <Grid item>
        <Card
          className={classNames.card + ' ' + classNames.newCard}
          variant="outlined"
        >
          <NewTaskDialog onCreate={handleCreate} />
        </Card>
      </Grid>
      <TaskList
        tasks={tasks}
        onAssign={handleAssign}
        onStepCompleted={handleCompletedStep}
      />
    </AuthPage>
  );
};

export default TaskDashboard;
