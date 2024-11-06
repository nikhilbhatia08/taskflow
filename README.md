# TaskFlow: Task scheduler in Go

![TaskFlow Hero](assets/taskflow.png#gh-light-mode-only)
![TaskFlow Hero](assets/taskflowdark.png#gh-dark-mode-only)

TaskFlow is an efficient task scheduler which is used to schedule jobs and tasks and multiple workers can pick those tasks and execute them

## System Components

## Life of a schedule

## Directory structure

Here's a brief overview of the project's directory structure:

- [`cmd/`](./cmd/): Contains the main entry points for the scheduler, coordinator, task queue and worker services.
- [`pkg/`](./pkg/): Contains the core logic for the scheduler, coordinator, task queue and worker services.
- [`data/`](./data/): Contains SQL scripts to initialize the db.
- [`tests/`](./tests/): Contains integration tests.
- [`*-dockerfile`](./docker-compose.yml): Dockerfiles for building the scheduler, coordinator, task queue and worker services.
- [`docker-compose.yml`](./docker-compose.yml): Docker Compose configuration file for spinning up the entire cluster.