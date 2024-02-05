# Brainstorm ideas

## Key points

- Sales teams need to report to engineering team bugs and improvements. At the moment, everything is handled using text chats with poor requirements. The idea is to find
a way to improve the process to allow Sales teams to create explanatory report for bugs and/or for new feature requests. A possible competitor / solution is https://jam.dev/
- Allow the company to customize the template for every customer, in order to have a sort of "guided" issues to open with all information needed by support team/ help desk / engineering. For example a customer A is using your product with SDK in python and customer B is just a consumer of your app/services, those two cases should provide different information when they ask for support.
- Keep track in the same tool of support ticket + developer roadmap in order to have an entire overview even for dev and customer support team and other team not directly involved.
- Reminder of a ticket after X our using mail notifications or 3rd party integration for notifications in order to keep customer up to date and never forget open tickets 


## Userbase

1. Webapp / consultings companies. Usually there are a lot of wasted time and poor documentation from the Sales team that results in unclear requirements for the dev team.

## Stack comparison

- https://vercel.com/pricing (front-end)
- https://fly.io/docs/about/pricing/ 
- https://railway.app/pricing


## Tech stack

- BE [Go](https://go.dev/) or [Rust](https://www.rust-lang.org/)
  - I will stay with GO for SaaS boilterplate repurpose. (maybe sellable in future)
- Tracking decisions [ADR Tools](https://github.com/npryce/adr-tools) - Why would you use this? When you make a decision, whatever it is, db table, how to do a thing in code, architecture or design of feature, it is tracked for future review ad you know what is changed in terms of design.
- Integration test with agnostic tool for Rest API [Hurl](https://hurl.dev/) - Why would you use this? This is super useful for having agnostic integration tests that should work with a simple tool like curl instead of having another code that runs QA tests or integration.
- If we want use conventional commits here a tool for simplify the way of commit https://github.com/cocogitto/cocogitto
- FE [React](https://it.legacy.reactjs.org/) or (https://nextjs.org/)
  - I would stay on nextjs, seems more modern
  - State management: Redux
- DB
  - psql for local development
  - on cloud it depends what services we want to use but there https://neon.tech or other mentioned above for BE
- GraphQL
  - https://www.prisma.io/blog/top-5-reasons-to-use-graphql-b60cfa683511
  - you can play also with the playground and test manually what is the output
 
We should go with vercel and fly.io
