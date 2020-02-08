# Track architecture decisions

## Decision

We will track architecturally significant decisions in a separate text files that are stored
in the same repository with source code. Store decision files under 
`${PROJECT_ROOT}/docs/adr/` and named according to the following format:
 
```adr-001-pick-mq-technology.md```

## Context

Tracking architecturally significant decisions is important for better understanding why the project
has evolved in a particular way and what was the reason behind this or that part of the solution.

Having ADRs in the same codebase allows considering code changes in context with the decisions.
Keeping them in text format provides an easy way for contributors to add and edit the decisions and
lowers entry barrier.

The idea of tracking ADRs comes from Michael Nygard and others.
