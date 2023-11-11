> Hello, you have stumbled upon my working todo for Plug. It lives in the documentation here because I bounce
> from problem to problem without a real scope defined for the territory it resides within. I acknowledge this
> project is not setup in a way that enables simple contributions. That is intentional for now. Focusing on what matters
> and that is not a nicer README. That said, below you can see my todo, active notes and things as they were completed.

TODO:

- [x] Figure out what to do about naming.
  - [x] Get the domain
  - [x] My favorite idea so far has been Plug cause 1) it is plug and play 2) it makes contracts plug and play with one another. It is also really short and I can get a domain that's easy to navigate to. Also, it's so simple that I finally have logo ideas.
  - [ ] Rename the Github repositories
  - [x] Rename the Discord
- [ ] Make sure our API endpoints are protected
- [ ] Think about the relayer implementation as this is what really has my interest.
- [ ] Update `authority` nomenclature to just use `permission` for increased conistency.

IN PROGRESS:

- [ ]

PENDING:

- [ ] Finish the documentation.
  - [ ] Remove the index page.
  - [ ] Make the branding confirm to the branding that has now been defined.

BUGS:

- [ ]

WANTS:

- [ ] Ability to declare a set of allowed cavets and domains.
      Notes: This is effectively `localization` for an api and its contained data.

RELEASING:

- [x] Make sure the landing page is responsive.

- [x] Get the raw client implementation functional with raw trpc connections.
- [x] Package up the trpc connector into the sdk to offer a more explicit integration path.
      Notes: This is where the `process.env.API_URL` and managed functions would be exposed.
- [x] Make sure that we can encode and decode each enforcer.

DONE:

- [x] Setup Supabase to be used for the production database on the landing site.
  - [x] How are you supposed to integrate schema generation into your CI/CD?
- [x] Deploy to Vercel
