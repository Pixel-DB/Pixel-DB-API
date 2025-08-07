![Logo](https://i.ibb.co/sJNyCH7J/Banner.png)

# PixelDB API

> This Service is still under production!!!

Pixel-BD is an open-source online platform where anyone can upload, share, and showcase their pixel art creations with the community.

This is the Repository for the [API](https://github.com/pixel-db/pixel-db-web), below you can see the features and the Routes. The PixelDB-Web is [here](https://github.com/pixel-db/pixel-db-api)

## Documentation

To access comprehensive API documentation, navigate to the <ins>BASEURL</ins>`/swagger` endpoint. This interactive documentation provides detailed information about all available API endpoints, including required parameters, response formats, and authentication requirements. The Swagger UI allows you to test API calls directly from your browser.

## Setup PixelDB with Docker

Set up the environment variables in a `stack.env` file like in the `.env.example`:

```env
    PORT=3000 #API Port
    DB_USER=example_user
    DB_PASSWORD=example_password
    DB_HOST=db-dev #Like in the Docker-Files
    DB_PORT=5432
    DB_NAME=example_db
    JWT_SECRET=example_secret
    MINIO_USER=root
    MINIO_PASSWORD=iamroot123
    MINIO_PORT=9001
    MINIO_BUCKET_NAME=pixelarts
```

Help:

```bash
  make help
```

Start in Development Mode:

```bash
  make dev
```

Start in Production Mode:

```bash
  make prod
```

## Feedback

If you have any feedback, please reach out to us on [Discord](https://discordapp.com/users/831464905131294730).

## Authors

- [@brainlesslukas](https://www.github.com/brainlesslukas)
