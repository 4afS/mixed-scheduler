# mixed-scheduler

[![ci](https://github.com/4afS/mixed-scheduler/actions/workflows/ci.yaml/badge.svg)](https://github.com/4afS/mixed-scheduler/actions/workflows/ci.yaml)

## Description
Mixed Schedule is an application for a person who is mixed up days and nights.

## Installation
Download binaries from the [releases page](https://github.com/4afS/mixed-scheduler/releases) and add it to your path.

## Set up schedule
Create schedule file to `$HOME/.mxs/schedule.yaml`.

- Example
```yaml
  - base: 9:00
  - plan:
    - start: 9:30
      title: eat breakfast
    - start: 12:00
      title: eat lunch
```

## Usage
- Show the schedule based on the current time:
```
  mxs -now
```

- Show the schedule based on given time:
```
  mxs -on 12:00
```
