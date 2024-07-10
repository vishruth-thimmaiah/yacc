# Y.A.C.C

## Run without Docker
```bash
cd frontend
bun run build

cd ..
go build -ldflags "-s -w" -o yacc
./yacc
```
## Run with Docker

```bash
sudo docker build --tag yacc .
```
```bash
sudo docker run -p 8000:8000 yacc
```
Docker now runs on port 8000.
