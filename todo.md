> Hello, you have stumbled upon my working todo for Plug. It lives in the documentation here because I bounce
> from problem to problem without a real scope defined for the territory it resides within. I acknowledge this
> project is not setup in a way that enables simple contributions. That is intentional for now. Focusing on what matters
> and that is not a nicer README. That said, below you can see my todo, active notes and things as they were completed.

- x: Done
- h: Holding
- o: Started

- TODO:

  - Alpha:

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
    - [x] Make adding a new tab direct to /canvas/create/
    - [x] Do not allow having multiple of the same tab open.
    - [x] Figure out what 'closed connection' is originating from when signing in.
    - [x] Fully functional authentication implementation
    - [x] Implement a basic auth stack so that we can the basic api flowing.
          Notes: In development we should just use a mock auth provider and ignore this piece for now.
    - [x] Get the basic authentication ux designed and boiled.
    - [x] Implement everything that is needed to actually run the websocket server and make sure everything is running properly.
    - [x] When this is in place bundle it up into `pnpm dev` to run everything at once
    - [x] Run a websocket server to power trpc
    - [x] Get user canvases from the database.

    - [O] Pick back up by finishing the conversion to page router. `layout.tsx` is the next thing to figure out since it is handled differently.

    - [o] Sync state changes to the database.
    - [o] Update the hello world to be the base noun plug.
    - [o] Easy framework to add and remove component types.
    - [o] Store the state of the canvas into the database.
      - [o] When we sync the state maybe we shouldn't even worry about saving when a component is updated and instead operate on just a set interval to check if changes have been made since the last broadcast and stream out the changes to the database. Realisticaly, my wanting to support websockets did not arise from the need of realtime collaboration. It was simply there to solve the issues that would arise when someone has two tabs open or someone is viewing a board that is not theirs.
    - [x] Retrieve the state of the canvas from the database.
      - [o] Need to get live updates in case we have multiple windows open.
        Notes: While it may seem like this is premature optimization, after having experience the difficulties of adding live responses without having properly planned for it.
    - [x] Figure out how to store complex component state in the canvas.
          Notes: We also have superjson as an option thought it will not completely solve the problem so it is likely not the right choice.
    - [ ] Update `ws.prod.ts` to actually be ready for production.
    - [ ] Add zoom buttons so that you don't have to scroll.
      - [ ] Fix the zoom having adverse effects on react-dnd.
    - [ ] Add 'home' button to center the screen on the center of the board.
    - [ ] Create an index table for head blocks.
          Notes: It should not be possible to "lose" a plug.
    - [ ] Add noun trait bid
    - [ ] Select and drag an area of components.
      - [ ] Implement ability to remove a component now that we have selection implemented.
      - [ ] Do not forget to implement revocation on deletion once the order signing has been added.

  - Beta:
    - [ ] Improve authentication implementation to support the use of email based accounts.
          Notes: Realistically, there is no reason to limit to hot-wallets meaning wallets that are yet to be cold-started should still be able to architecture setups as well as sign orders.
          Notes: There is also privy, but pricing is way more expensive. Imagine paying hundreds just for logging in alone. Bad.
    - [ ] Add the ability to be an anonymous user (until signed in)
      - [ ] Implement the page redirects as a middleware so that we do not ever even get routed there. The reason this has not just been done is because I do not want to make two calls to the database (1 in the middleware and then another on the page) to determine if it is public / the accessing user has access so just not worrying about this for now. In MVP just going to design for public, but still require that users be logged in.
        - [ ] When we do this, we will want to categorize them as a 'Anonymous Cow/Cat/Dog' and show them as active on the board when they are viewing it.
        - [ ] For anonymous users, everything will always be read-only.

PENDING:

- [o] Finish the documentation.
  - [x] Remove the index page.
  - [x] Make the branding confirm to the branding that has now been defined.
  - [ ] Edit all of the static documentation and write the content for pieces that are not auto-generated.

BUGS:

- [ ] Right now tabs will run into issues when you reach the edge of the page.
      Notes: I am pretending this isnt happening and will get to it when I get to it.
- [ ] Drag and drop tab re-ordering.
      Notes: I do not really care about this but I am sure someone will.
- [ ] If you do not have any canvases in the database yet, create one, and then go back to the home page, if you then try and create another new one it forces you back to the newly created one instead of allowing you to access the page as you should because you no longer have zero canvases.

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
