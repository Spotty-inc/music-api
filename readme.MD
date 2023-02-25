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

e.g. sudo docker run -e SPOTIFY_ID="myspotifyid" -e SPOTIFY_SECRET="myspotifysecret" -d samlangley/music-api:latest

The service can now be accessed on port 5000
