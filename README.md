# build frontend
cd frontend
npm i
npm run build
cd ../

# build app
go build

# refer
https://www.liip.ch/en/blog/embed-sveltekit-into-a-go-binary

# copy fighter lol
go install github.com/jmhodges/copyfighter@latest
copyfighter.exe .
