import { NextPage } from 'next';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';

import { Card, Divider, Typography } from '@material-ui/core';

import NewTaskDialog, { NewTaskForm } from '../../../src/components/organisms/NewTaskDialog';
import TaskList from '../../../src/components/organisms/TaskList';
import AuthPage from '../../../src/components/templates/AuthPage';
import { ProjectApi, TaskApi, useFacade } from '../../../src/lib/api';
import { ReturnTypeAsync } from '../../../src/lib/type';

type P = {};

const ProjectDetail: NextPage<P> = ({}) => {
  const router = useRouter();
  const projectID = router.query.projectID as string;
  console.log(projectID, router.query);

  const api = useFacade(ProjectApi);
  const taskAPI = useFacade(TaskApi);

  type Project = ReturnTypeAsync<typeof api.getProject>;
  const [project, setProject] = useState<Project>(null);

  const load = async () => {
    try {
      const project = await api.getProject({ projectId: projectID });
      setProject(project);
    } catch (error) {
      alert(error);
    }
  };

  const handleCreate = async ({ name, description }: NewTaskForm) => {
    console.log({ name, description });
    try {
      const task = await api.createProjectTask({
        taskForm: { name, description },
        projectId: projectID,
      });
      project.tasks.concat(task);
      setProject({ ...project, tasks: project.tasks.concat(task) });
    } catch (error) {
      alert(error);
      console.log(error);
    }
  };

  const handleAssign = async (taskID: string) => {
    const task = await taskAPI.assignSignedInUserToTask({ taskId: taskID });
    setProject({
      ...project,
      tasks: project.tasks.map((t) => (t.id === task.id ? task : t)),
    });
  };

  const handleCompletedStep = async (
    taskId: string,
    step: string
  ): Promise<boolean> => {
    try {
      switch (step) {
        case 'TODO':
          const nextTODO = await taskAPI.progressTask({ taskId });
          setProject({
            ...project,
            tasks: project.tasks.map((t) =>
              t.id === nextTODO.id ? nextTODO : t
            ),
          });
          return true;
        case 'WIP':
          const nextWIP = await taskAPI.doneTask({ taskId });
          setProject({
            ...project,
            tasks: project.tasks.map((t) =>
              t.id === nextWIP.id ? nextWIP : t
            ),
          });
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

  useEffect(() => {
    load();
  }, []);

  return (
    <AuthPage title={project?.name || 'Loading...'}>
      {!project ? (
        <Typography>Loading...</Typography>
      ) : (
        <div>
          <div>
            <Typography variant="body1">Elevator Pitch</Typography>
            <Typography variant="body2">{project.elevatorPitch}</Typography>
          </div>
          <Divider />
          <div>
            <Typography variant="body1">Tasks</Typography>
            <NewTaskDialog onCreate={handleCreate} />
            <TaskList
              tasks={project?.tasks || []}
              onAssign={handleAssign}
              onStepCompleted={handleCompletedStep}
            />
          </div>
        </div>
      )}
    </AuthPage>
  );
};

export default ProjectDetail;
