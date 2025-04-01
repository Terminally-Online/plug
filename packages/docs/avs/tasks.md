Tasks in the Circuit AVS are defined as a signed and built transaction with the expectation that the Performer will execute the transaction on the target chain.

## Task Definition

Task are defined by their transaction data and target chain. Users express preferences and constraints using the Plug application frontend. These desires are distributed as intents to the solver network which builds the tran

## Performer Selection

Plug will bootstrap the network by operating as a Performer and executing transactions built by the solver network. Operators are encouraged to join Circuit as an Attester and contribute to ongoing decentralization efforts. 

## Proof of Task

After executing the transaction associated with a task, Performers send a proof of task demonstrating the initial signed order was executed on the target blockchain as described. 

## Validation of Task Completion

Attesters run the validation service to determine if the built transaction instructions were properly executed on the target chain. Attestation of outcome is then posted onchain to the AttestationCenter ensuring that all transaction execution activities and results are tracked on a public and immutable ledger. 

