import React from 'react';

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
import MoreVertIcon from '@material-ui/icons/MoreVert';

import { TaskApi } from '../../../lib/api';
import { abbrev } from '../../../lib/text';
import { datediff } from '../../../lib/time';
import { KeyType, ResultsItemType } from '../../../lib/type';
import StatusSteps from '../StatusStep';

const useStyles = makeStyles((theme) =>
  createStyles({
    card: {},
  })
);

type Task = ResultsItemType<KeyType<TaskApi, 'listAllAssociatedWithUser'>>;

type P = {
  tasks: Task[];
  onAssign?: (taskID: string) => any;
  onStepCompleted?: (taskID: string, step: string) => any;
};

const MAX_CONTENT = 100;
const MAX_TITLE = 30;

const TaskList: React.FC<P> = ({
  tasks = [],
  onAssign = () => {},
  onStepCompleted = () => {},
}) => {
  const classNames = useStyles();

  const handleAssign = (taskID: string) => {
    onAssign(taskID);
  };

  const handleCompletedStep = (taskID: string, step: string) => {
    onStepCompleted(taskID, step);
  };

  return (
    <Grid>
      {tasks.map((t) => (
        <Grid item key={t.id} xs={12}>
          <Card className={classNames.card} variant="outlined">
            <CardHeader
              avatar={
                t.assigneeId ? (
                  <Avatar aria-label="task">{t.assigneeId.substr(0, 6)}</Avatar>
                ) : (
                  <Button onClick={() => handleAssign(t.id)}>
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
            />
            <CardContent>
              <Typography variant="body2" color="textSecondary" component="p">
                {abbrev(t.description, MAX_CONTENT)}
              </Typography>
            </CardContent>
            <CardActions>
              <StatusSteps
                status={t.status}
                onCompleteStep={(step) => handleCompletedStep(t.id, step)}
              />
            </CardActions>
          </Card>
        </Grid>
      ))}
    </Grid>
  );
};

export default TaskList;
