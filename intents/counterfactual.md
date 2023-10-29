---
head:
    - - meta
      - property: og:title
        content: Counterfactual
    - - meta
      - name: description
        content: A brief overview of what a counterfactual is and how it applies to Emporium.
    - - meta
      - property: og:description
        content: A brief overview of what a counterfactual is and how it applies to Emporium.
---

# Counterfactual

**This word can seem very confusing.** It's thrown all around and seems to mean something different to every person and situation. But, it's actually pretty simple.

## What is a counterfactual?

A counterfactual is a statement that is contrary to fact.

-   It's a statement that is not true.
-   It's a statement that is false.
-   It's a statement that is not real.
-   It's a statement that is not the case.

That is all a counterfactual is. It's a statement that is not true. So, if you say something that is not true, it could also be referred to as a counterfactual.

[Counterfactual statements give rise to the popular: "_If this, then that._"](https://christophm.github.io/interpretable-ml-book/counterfactual.html)

This is true within the `Emporium` as well. As a user that declares execution permissions one has the ability to not just declare conditions that must be true, but **false** as well as the action to follow if execution is allowed or denied.

## Applying this to `Emporium`

Now that you have a better understanding of what a counterfactual is, let's close the gap and bring the perspective for the use in `Emporium`.

Emporium utilizes counterfactual assertations to determine if a user has the ability to execute a function. This is done by declaring a condition that must be met. If the condition is not met, then the user does not have the ability to execute the function.

Let's look at an example that uses `Alice` and `Bob` as users:

-   **If** `Alice` doesn't have permission to `X`, **then** Alice can't execute for `Bob`.

-   **If** condition `Z` is not satisfied, **then** `Alice`can't execute for`Bob`.

You can forget about all the complexity and technical jargon and simply think of `counterfactual` as the driving pipeline of "_if this, then that_" statements that are verified and executed on the EVM blockchain of your choosing.
