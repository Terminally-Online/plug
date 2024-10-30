# Execution + Simulation

-   Right now, we have an execution in the database that controls the running of simulations.
-   Functionally, although the Solver will always be able to build a transaction, it will not always be able to execute it.
    -   This is because the Solver is not actually simulating the transaction as it runs, instead it is simulating through the socket.
-   So, simulations should only ever show the final outcome of the transaction that is built.

-   This means we have three different outcomes:
    -   Success: The transaction is executed successfully
    -   Warning: The transaction was built and simulated, but not executed
    -   Failure: The transaction is not simulated successfully

- In practice, as far as the user is concerned, the "simulation" is less a "simulation" and more a "run" of the transaction.

- For each simulation, a user should be able to drill down into the simulation and related transaction.
    - If it was a failure, the error should be displayed relative to each action.
    - If it was a success, the transaction details should be displayed.

# Onboarding

-   We only want to onboard authenticated users. A new user logging in really shouldn't have that experience. If they just get going, they are effectively onboarding themselves.

    -   This may honestly not really be what we wanted, but if we have a stronger call to action to authenticate, attention can at least be called to it.

-   Onboarding often feels like bullshit because you don't really know what onboarding is doing, or how long it's going to take.

    -   If you can't follow a thread of logic, you start to just bullshit it and try and get through it as quickly as possible.
    -   This means that while we are going to ask them questions to curate their experience, we can quickly visualize the experience has been curated.
    -   This means that realistically a new user (identifier wise) MUST have the option to say "hey im not really a new user"

-   Initial trigger for this different direction of onboarding came from here:
    friction is ongoing burden, friction isn’t binarily “more steps”

        ie: it slows down your day to day procedures. arguably, not having names is the real long term user friction

        even acknowledging it as “friction” though, yes. software today feels unremarkable and uncontrolled because so much “friction” has been removed making everything feel impersonal. proper account management and naming is a good step to counter that

-   Realistically, one-time onboarding kind of misses the mark because you expect a user to frontload all the context they need.

    -   A tutorial does not appropriately address this.
    -   When you are starting a new game, you learn the fundamental requirements almost immediately:
        -   This is how you walk, this is how you attack, this is how you run
    -   Following that though, the rest of your understanding is microdosed in highly specific contexts.
    -   GTA is the largest world out there, yet they drip feed you the information you need to navigate the city as a whole. Today, most technical onboarding does not accomplish this.

-   When you are a new user, you do not want to learn how to do something you're not prepared to, you want to know how to do things one step at at time.
-   Before the game even begins, you choose your difficulty, this is effectively the same as asking "How experience are you with Plug?" to determine how much they are spoon-fed.

-   While it may seem logical to store onboarding information on their local device, really it belongs in the database. This means that if someone connects a new wallet, they need a way to bypass this because they are not really a new user.

-   Long form onboarding / knowledge dissementation will be most successful when it is clearly communicated as a drip. If you only reveal 1% of what they need to know, but they assume it to be 100%, they will feel confused and quit before they make it further in the process to realize 100% actually is taught, they just don't need to know it yet.

-   On mobile, onboarding is quite clear when on mobile as it's a full page takeover.
-   On desktop, there is a lot of layout that enables a significant amount of information to be delivered in one eye scan of the screen.

    -   This means that the experience of actually onboarding needs to limit distractions and immediate exposure.
        -   This does however conflict with the idea of `DEFAULT_COLUMNS` because then what is the user actually loading into?

-   I think as a whole, this concept can be referred to as `deferred user onboarding`.

## Onboarding Steps

→ (Pre-onboarding) Understand the user

1. "How experienced are you with Plug?"

    1. Not at all → Kick them into the tutorial
    2. Relatively (they've used it before aren't a poweruser) → They'll go to the real app, but they could use some tips
    3. Very (self selecting as a poweruser) → Just let em cook

2. "What things do you want to be doing most often?"
    1. Give them a list of all the categories and they can toggle the ones they care about.
        - This will end up being used for curated release -- Having a curated segmentation enables the ability to have longer timelines on what is curated and what rolls through that segment. This also means that if we have an in-category Plug that peforms really well in-category, then it suggests there may be some correlation with other out-of-category users for a little exposure therapy.
        - It would be cool, if the category selections here even informed the plug that they are sandboxed into during the onboarding session.
            - Someone in DeFi is not going to care about Nouns as much even if it is possible to get an immediate click there.

→ (Onboarding) Knowledge sharing: Enable the user (they land into a sandboxed console -- this is not the real experience)

-   Load in default columns that are My Plugs and Activity
-   The spotlight that is active is based on the `has` variables.

-   Spotlight: `hasStarted = false`: Get started by clicking the get started Plug

    -   This sets `hasStarted` to `true`

-   Spotlight: `hasAddedAction = false`: Add the "bid" action to the pre-populated starter plug.
    -   This sets `hasAddedAction` to `true`
-   Spotlight: `hasFocusedTraitType = false`: Click an action attribute to open its frame
    -   This sets `hasFocusedTraitType` to `true`.
-   Spotlight: `hasAddedTraitType = false`: Select which trait type to add
    -   They select some type. This sets `hasAddedTraitType` to `true`
-   Spotlight: `hasAddedTrait = false`: Select with trait to choose

    -   They select some trait setting `hasAddedTrait` to `true`

-   Spotlight: `hasRun = false`: Run the plug that they've created
    -They click run and set `hasRun = true`

-   Now we focus `Activity` and highlight the outcome of the intent they ran
    -   (There will be some nuance here that I have not yet answered or solved)
-   Now they can add any column -- Ideally, the discovery column
-   Now we direct them to authenticate
    -   Before this stage, do not show them anything that can be distracting such as actions that are not be called attention to.
        -   Add plug button
        -   Search button
        -   View as button
        -   We will keep all the options in the columns, since it doesn't really made which one they add first.
        -   We will however hide the "add column" section until they've reached this point of the onboarding flow

→ (Post-onboarding) Skill sharing: Empower the user to be as successful as possible
