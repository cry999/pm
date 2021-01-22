import { NextPage } from 'next';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';

import {
  Avatar,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  createStyles,
  Grid,
  IconButton,
  makeStyles,
  Typography,
} from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';
import CachedTwoToneIcon from '@material-ui/icons/CachedTwoTone';
import MoreVertIcon from '@material-ui/icons/MoreVert';

import NewTaskDialog, { NewTaskForm } from '../../src/components/organisms/NewTaskDialog';
import StatusSteps from '../../src/components/organisms/StatusStep';
import AuthPage from '../../src/components/templates/AuthPage';
import { facade, TaskApi } from '../../src/lib/api';
import { abbrev } from '../../src/lib/text';
import { datediff } from '../../src/lib/time';
import { KeyType, ResultsItemType } from '../../src/lib/type';

const useStyles = makeStyles((theme) =>
  createStyles({
    reloadIcon: {
      width: 50,
      height: 50,
      float: 'right',
      margin: theme.spacing(1),
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
      height: theme.spacing(45),
      width: theme.spacing(35),
      margin: theme.spacing(2),
    },
    cardHeader: {
      height: theme.spacing(14),
    },
    cardContent: {
      height: theme.spacing(12),
    },
    cardStatus: {
      height: theme.spacing(11),
    },
  })
);

type P = {};

type Task = ResultsItemType<KeyType<TaskApi, 'listAllAssociatedWithUser'>>;

const MAX_CONTENT = 100;
const MAX_TITLE = 30;

const TaskDashboard: NextPage<P> = ({}) => {
  const classNames = useStyles();
  const api = facade(TaskApi);
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
      const { results } = await api.listAllAssociatedWithUser();
      setTasks(results);
    } catch (e) {
      alert(e);
    }
  };

  return (
    <AuthPage
      title={
        <Grid item>
          <Typography variant="h3" className={classNames.title}>
            Task Dashboard
          </Typography>
          <IconButton className={classNames.reloadIcon}>
            <CachedTwoToneIcon onClick={handleReload} />
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
      {tasks.map((t) => (
        <Grid item key={t.id}>
          <Card className={classNames.card} variant="outlined">
            <CardHeader
              avatar={
                t.assigneeId ? (
                  <Avatar aria-label="task">{t.assigneeId.substr(0, 6)}</Avatar>
                ) : (
                  <Button>
                    <Avatar aria-label="task">Do!</Avatar>
                  </Button>
                )
              }
              action={
                <IconButton aria-label="settings">
                  <MoreVertIcon />
                </IconButton>
              }
              title={abbrev(t.name, MAX_TITLE)}
              subheader={
                'last updated at ' + datediff(new Date(), t.updatedAt) + ' ago'
              }
              className={classNames.cardHeader}
            />
            <CardContent className={classNames.cardContent}>
              <Typography variant="body2" color="textSecondary" component="p">
                {abbrev(t.description, MAX_CONTENT)}
              </Typography>
            </CardContent>
            <CardActions className={classNames.cardStatus}>
              <StatusSteps
                status={t.status}
                onCompleteStep={(step) => handleCompletedStep(t.id, step)}
              />
            </CardActions>
          </Card>
        </Grid>
      ))}
    </AuthPage>
  );
};

export default TaskDashboard;
