# Track architecture desicions

## Decision

We well track architecurally significant decisions in a separate text files that are stored
in the same repository with source code. Desicion files are stored under 
`${PROJECT_ROOT}/docs/adr/` and named according to the following format:
 
```adr-001-pick-mq-technology.md```

## Context

Tracking architecturally significant decisions is important for better understanding why the project
has evolved in a particular way and what was the reason behing this or that part of the solution.

Having ADRs in the same codebase allows to consider code changes in context with the decisions.
Keeping them in text format provides an easy way for contributors to add and edit the decisions and
lowers entry barrier.

The idea of tracking ADRs comes from Michael Nygard and others.
