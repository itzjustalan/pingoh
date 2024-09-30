# Pingoh - ping your websites periodically
Inspired by [Uptime Kuma](https://github.com/louislam/uptime-kuma), this is a simple uptime monitoring tool that pings your websites periodically and notifies you if they are down. But without all the bloat the binary aims to be less than 20MB thanks to golang.

# build instructions
Make sure you have the dependencies installed.
```bash
❯ node -v
v20.15.0
❯ go version
go version go1.22.5 linux/amd64
```
Once you have them deps installed, you can build the app -
```bash
cd frontend
npm ci .
npm run build
cd ../
# build app
go build
```

# docker
```bash
docker build -t pingoh .
docker run pingoh
# port forwarding where 4001 is the host port and 4002 is the container port
docker run -p 4001:4002 pingoh --port :4002 # other pingoh flags
```

# usage
pingoh --help
  --db string
        db file path (default "pingoh.db")
  --email string
        admin user email (default "admin@mail.com")
  --log string
        log file path (default "pingoh.log")
  --password string
        admin user password (default "password")
  --port string
        port number in :3000 format (default ":3000")

# for development
Install [nodemon](https://nodemon.io) with `npm install -g nodemon` and run `nodemon` in the root directory.
Nodemon should pick up the ./nodemon.json configuration file and run the app with the correct configuration.
For the frontend, run `npm run dev` in the frontend directory.

# refer
https://www.liip.ch/en/blog/embed-sveltekit-into-a-go-binary

# copy fighter lol
go install github.com/jmhodges/copyfighter@latest
copyfighter.exe .

# guidelines
Lint and format code before commit.
```bash
# backend
go fmt ./...
# frontend
cd frontend
npm run format
npm run lint
```
