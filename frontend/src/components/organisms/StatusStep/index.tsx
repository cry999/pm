import React, { useState } from 'react';

import {
  createStyles,
  makeStyles,
  Step,
  StepButton,
  Stepper,
  Typography,
} from '@material-ui/core';
import { green } from '@material-ui/core/colors';
import CheckCircleTwoToneIcon from '@material-ui/icons/CheckCircleTwoTone';

const useStyles = makeStyles((theme) =>
  createStyles({
    root: {
      width: '100%',
    },
    active: {
      color: theme.palette.primary.light,
    },
    notyet: {
      color: theme.palette.text.secondary,
    },
    completed: {
      color: theme.palette.primary.light,
    },
    alldone: {
      color: green.A200,
    },
  })
);

type P = {
  status: string;
  onCompleteStep?: (status: string) => any;
};

const steps = ['TODO', 'WIP', 'DONE'];

const StatusSteps: React.FC<P> = ({ status, onCompleteStep = () => true }) => {
  const classNames = useStyles();
  const activeStep = steps.indexOf(status);

  const handleStep = (completedStep: number) => {
    onCompleteStep(steps[completedStep]);
  };

  return (
    <div className={classNames.root}>
      <Stepper nonLinear activeStep={activeStep}>
        {steps.map((label, index) => (
          <Step key={label}>
            <StepButton
              onClick={() => handleStep(activeStep)}
              completed={index < activeStep}
              disabled={index !== activeStep + 1}
              className={
                activeStep === steps.indexOf('DONE')
                  ? classNames.alldone
                  : index > activeStep
                  ? classNames.notyet
                  : index === activeStep
                  ? classNames.active
                  : classNames.completed
              }
              icon={
                index >= activeStep ? (
                  <Typography variant="body2">{label}</Typography>
                ) : (
                  <CheckCircleTwoToneIcon />
                )
              }
            ></StepButton>
          </Step>
        ))}
      </Stepper>
    </div>
  );
};

export default StatusSteps;
