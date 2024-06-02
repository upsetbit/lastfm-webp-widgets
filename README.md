# Last.fm WebP Widgets

## Now Playing

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://caian-org.s3.amazonaws.com/lastfm-now-playing-dark.webp" width="600px">
  <img src="https://caian-org.s3.amazonaws.com/lastfm-now-playing-light.webp" width="500px">
</picture>

### Running with Docker

```bash
docker run \
    --rm \
    -e LASTFM_USERNAME='your-lastfm-username' \
    -e LASTFM_API_KEY='your-lastfm-api-key' \
    -e S3_BUCKET_NAME='an-s3-bucket-you-own' \
    -e AWS_ACCESS_KEY_ID='access-key-to-upload-to-the-bucket' \
    -e AWS_SECRET_ACCESS_KEY='secret-key-to-upload-to-the-bucket' \
    -e AWS_REGION='us-east-1' \
    caian/lastfm-widget-now-playing
```

### Build from source and run

```bash
just run lastfm-now-playing
```

or

```bash
cd cmd/lastfm-now-playing \
    && go build -tags exec_local,save_s3 -o ../../bin/lastfm-now-playing \
    && cd ../.. \
    && ./bin/lastfm-now-playing
```
