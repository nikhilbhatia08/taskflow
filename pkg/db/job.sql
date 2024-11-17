
CREATE TABLE jobs (
    id UUID PRIMARY KEY,
    queuename text NOT NULL,
    payload jsonb NOT NULL,
    retries INTEGER NOT NULL,
    maxretries INTEGER NOT NULL,
    status INTEGER,
    picked_at TIMESTAMP,
    started_at TIMESTAMP, -- when the worker started executing the task.
    completed_at TIMESTAMP, -- when the task was completed (success case)
    failed_at TIMESTAMP -- when the task failed (failure case)
);
