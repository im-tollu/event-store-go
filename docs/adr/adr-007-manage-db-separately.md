# Manage database migrations as a separate code

## Decision

We will use migration approach for managing database.
Code for doing it will be implemented separately from other application code.
It will have a separate package and a separate runner.
