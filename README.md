# build frontend
cd frontend
npm i
npm run build
cd ../

# build app
go build

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
