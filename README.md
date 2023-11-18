This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

First, run a Postgres database:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5434:5432 -d postgres
```

If you spend a lot of time in the terminal or working with postgres database it is probably worth adding these helper alias to your `~/.zshrc`:

```bash
# log the postgres database url
alias pdb_log_url="echo 'postgres://postgres:postgres@localhost:5434/postgres'"
# start a postgres database with docker
alias pdbs="pdb_log_url && docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5434:5432 -d postgres"
# access the postgres database
alias pdbg="docker exec -it postgres psql -U postgres"
# clean up the postgres database
alias pdbc="docker stop postgres && docker rm postgres"
# start and access the postgres database
alias pdb="pdbs && pdbg"
```

With the database up and running you will need to run the migration withs:

```bash
npx prisma migrate
```

Finally, you are ready to run the development server:

```bash
pnpm dev
```
