> Hello, you have stumbled upon my working todo for Plug. It lives in the documentation here because I bounce
> from problem to problem without a real scope defined for the territory it resides within. I acknowledge this
> project is not setup in a way that enables simple contributions. That is intentional for now. Focusing on what matters
> and that is not a nicer README. That said, below you can see my todo, active notes and things as they were completed.

x: Done
o: Started

TODO:
  Alpha:
    - [x] Get infinite canvas functional.
        Notes: It is really rather simple once you understand what an infinite canvas actually is. Especially because we do not even actually care about it being infinite for real.
    - [x] Drag components.
        Notes: We should just use react-dnd so that we do not have to worry about building our own drag implementation.
      - [x] Align items to the grid (grid snapping).
    - [x] Infinite canvas and drag and drop functional at the same time.
    - [x] Prevent a user from selecting the same pin multiple times in a plug.
    - [x] Prevent a user from adding a new pin when all available pins have been exhausted.
    - [o] Page where a user can see their canvases
    - [x] Tab functionality
      - [x] Active canvas
      - [x] Close a tab
      - [x] Add a tab
    - [ ] Make adding a new tab direct to /canvas/create/
    - [ ] Do not allow having multiple of the same tab open.

    - [ ] Figure out what 'closed connection' is originating from when signing in.
    - [ ] Run a websocket server to power trpc (+ auth?)
    - [o] Implement a basic auth stack so that we can the api flowing.
        Notes: In development we should just use a mock auth provider and ignore this piece for now.
    - [o] Sync state changes to the database.
    - [o] Update the hello world to be the base noun plug.
    - [o] Easy framework to add and remove component types.
    - [o] Store the state of the canvas into the database.
      - [o] When we sync the state maybe we shouldn't even worry about saving when a component is updated and instead operate on just a set interval to check if changes have been made since the last broadcast and stream out the changes to the database. Realisticaly, my wanting to support websockets did not arise from the need of realtime collaboration. It was simply there to solve the issues that would arise when someone has two tabs open or someone is viewing a board that is not theirs.
    - [o] Retrieve the state of the canvas from the database.
      - [o] Need to get live updates in case we have multiple windows open.
          Notes: While it may seem like this is premature optimization, after having experience the difficulties of adding live responses without having properly planned for it.
    - [o] Figure out how to store complex component state in the canvas.
          Notes: We also have superjson as an option thought it will not completely solve the problem so it is likely not the right choice.
    - [ ] Select and drag an area of components.
      - [ ] Implement ability to remove a component now that we have selection implemented.
      - [ ] Do not forget to implement revocation on deletion once the order signing has been added.
    - [ ] Store ephemeral state that is not yet valid.
    - [ ] Add zoom buttons so that you don't have to scroll.
      - [ ] Fix the zoom having adverse effects on react-dnd.
    - [ ] Add 'home' button to center the screen on the center of the board.
    - [ ] Create an index table for head blocks.
        Notes: It should not be possible to "lose" a plug.
    - [ ] Think about the relayer implementation as this is what really has my interest.
    - [ ] Add noun trait bid

  Beta:

    - [ ] Improve authentication implementation to support the use of email based accounts.
      Notes: Realistically, there is no reason to limit to hot-wallets meaning wallets that are yet to be cold-started should still be able to architecture setups as well as sign orders.
      Notes: There is also privy, but pricing is way more expensive. Imagine paying hundreds just for logging in alone. Bad.

PENDING:

- [o] Finish the documentation.
  - [x] Remove the index page.
  - [x] Make the branding confirm to the branding that has now been defined.
  - [ ] Edit all of the static documentation and write the content for pieces that are not auto-generated.

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

- [x] Final nomenclature update
  - [x] Update `authority` nomenclature to just use `pin` for increased conistency.
  - [x] "Naming of things doesn't matter!" -- Anyone that says this has never actually launched something in this industry and if they have and didn't learn this lesson then that is their own problem.

- [x] Figure out what to do about naming.

  - [x] Get the domain
  - [x] Deploy on the new domains
  - [x] My favorite idea so far has been Plug cause 1) it is plug and play 2) it makes contracts plug and play with one another. It is also really short and I can get a domain that's easy to navigate to. Also, it's so simple that I finally have logo ideas.
  - [x] Rename the Github repositories
  - [x] Rename the Discord
  - [x] Deprecate the old npm packages
  - [x] Deploy documentation at docs.onplug.io
  - [x] Week 1 report

- [x] Setup Supabase to be used for the production database on the landing site.
  - [x] How are you supposed to integrate schema generation into your CI/CD?
- [x] Deploy to Vercel

- [x] Make sure our API endpoints are protected
