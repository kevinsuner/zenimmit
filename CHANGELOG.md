# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.1] - 2024-07-15

- chore(core): update go.mod module name (7768f37)

## [1.0.0] - 2024-07-15

### Added

- docs(core): add README.md (3cbb3df)
- chore(core): add MIT license (cf1b266)
- refactor(components): remove the option to quit program pressing <q> (dd85b5c)
- style(stepper): update steps questions (516a057)
- style(components): update placeholder, remove extra line (6ab5b10)
- style(components): update placeholder (a34d2c7)
- refactor(config): update scopes (c87157c)
- feat: add missing steps to perform a commit with or without body (4de5bc7)
- refactor: lowercased placeholder text (01b2fe6)
- feat: added a textarea component to allow the user to introduce a commit body (175292d)
- refactor: use subjectLength, remove errMsg (659c05f)
- chore: downgrade github.com/charmbracelet/bubbletea to 0.26.2 (8ec2375)
- feat: add required constants and variables (7b4e869)
- refactor: add nil check for model (aa7d416)
- feat: added component to render text to the screen (7a5007b)
- feat: added an input component which uses the stepper to fill the commit subject (42cde1c)
- refactor: removed usage of pointers (c10bf83)
- chore: add package github.com/charmbracelet/bubbles (ec182e2)
- refactor: using newStepper and newList (40df67c)
- feat: added a stepper which allows to perform an action on each step (5ba31bf)
- feat: added a list component that makes use of the stepper to advance through the program steps (576514b)
- refactor: remove pkg/list and pkg/stepper (db3a7dd)
- chore: add .gitignore (d2b445a)
- feat: add example zenimmit.yaml file (dc3651f)
- wip: start work on a stepper to perform different actions based on the step (014363c)
- wip: start work on a list component to select between multiple options (56b8411)
- feat: load zenimmit.yaml using viper, start new tea program (782966f)
- chore: init go project, add bubbletea and viper packages (782482d)