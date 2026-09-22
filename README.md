# Todo app

### ⚠️ Work in progress - project is being built for learning purpouses.

## Current status:
- [x] Frontend 
- [x] Backend
- [x] DB migrations using Job
- [x] CronJob for generating scheduled Todos
- [x] Communication in k8s via Ingress
- [x] Secrets using SOPS

## Architecture

- Backend - Go REST API, talks to postgres and passes values to frontend. DB code generated using sqlc
- Frontend - Go fileserver serving HTML, CSS and JS
- CronJob - Python application posting a random Wikipedia link added to todo.
- Postgres - DB, with PVC storag
- Ingres - routing via middleware / -> frontend and /api -> backend 