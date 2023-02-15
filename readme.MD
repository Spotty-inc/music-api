# Project Title

Music API

## Description

API that authenticates and fetches data from the Spotify web API to return data on music, podcasts, and more

## Usage

The application should be run as a docker container and the following environment variables are required

| Name      | Description |
| ----------- | ----------- |
| SPOTIFY_ID    | Client ID for a registered Spotify application    |
| SPOTIFY_SECRET  | Client Secret for a registered Spotify application     |

Port 5000 is exposed within the container so you will need to port forward from a local port if you're running through a host

e.g. sudo docker run -e SPOTIFY_ID="myspotifyid" -e SPOTIFY_SECRET="myspotifysecret" -d -p 8080:5000 samlangley/music-api:latest

The service should now be running and can be accessed at localhost:5000 or localhost:8080 if port forwarding through one of it's defined endpoints

e.g GET localhost:5000/random
