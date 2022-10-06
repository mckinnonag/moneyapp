
# Money App

A web application with a React frontend and a Go backend which allows users to connect multiple bank accounts, view transactions, and split those transactions with other users.

The app uses the [Plaid API](https://plaid.com/docs/api/) to connect with banking institutions, and [Google Firebase](https://firebase.google.com/) for auth and user management.

This is still a work in progress!

### Setup
1. Install [Docker](https://docs.docker.com/engine/install/) and [Node and npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm).
2. Create a [Plaid](https://plaid.com/) account and add the API key and secret to the .env file.
3. Create a [Google Firebase](https://firebase.google.com/) project.
4. Create a [Firebase service account](https://firebase.google.com/docs/admin/setup).
5. Export the Firebase keys as a .json file, and add the .json file to the project's gitignore.
6. Copy the .json file to the main directory.

### Test
* Come back soon - currently working on this.

### Run

1. Start docker
2. Start the server and database:

        docker compose up

3. Start the frontend:

        cd ./frontend && npm start
