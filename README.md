# Y.A.C.C
## Yet Another Communication Catalyst
A comms app that uses websockets(text, media) and webrtc(live video).

### Components:
- **Backend**: Golang
- **Frontend**: Vuejs/Typescript
- **Database**: Postgresql, S3

### Installation:
1. Clone the repo:
     ```bash
     git clone https://github.com/vishruth-thimmaiah/yacc.git
     cd yacc
     ```
2. Environment Variables:
   create a .env and fill in the following:
    - POSTGRES_URL
    - OAUTH_GOOGLE_CLIENT
    - OAUTH_GOOGLE_SECRET
    - AWS_ENDPOINT_URL
    - AWS_REGION
    - AWS_ACCESS_KEY_ID
    - AWS_SECRET_ACCESS_KEY

#### without Docker:
  3. Install golang and bun:
     - Go: https://go.dev/doc/install
     - Bun: https://bun.sh/docs/installation
  4. build the frontend:
     ```bash
     cd frontend
     bun run build
     cd ..
     ```
  5. Build the backend:
     ```bash
     go build -ldflags "-s -w" -o yacc
     ```

  6. Run with:
     ```bash
     ./yacc
     ```
#### With Docker:

  3. Building:
     ```bash
     sudo docker build --tag yacc .
     ```
  5. Running:
     ```bash
     sudo docker run -p 8000:8000 yacc
     ```
