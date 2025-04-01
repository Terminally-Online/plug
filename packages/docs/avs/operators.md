# Operators

## Roles

### Performer

Performer is an AVS Operator that executes a task, provides a Proof of Task, and sends the results to [Attesters](#attester). After successfully executing a task, the Performer publishes an event via peer-to-peer networking for [Attester](#attester) nodes to discover.

The RPC call that the Performer sends to the [Attestor](#attester) Nodes:

```json
{
    "jsonrpc": "2.0",
    "method": "sendTask",
    "params": [<proofOfTask>, <data>, <taskDefinitionId>,  <performerAddress>, <signature>]
}
```

### Attester

Attesters are AVS Operators' quorum that attests to the validity of the executed task. Each task must be attested as either "valid" or "invalid".

The Operator's voting power is proportional and calculated against the amount of re-stake assets staked on the shared security layer, referred to as “dynamic voting power.” The re-staked effective balance determines each Operator's influence in the consensus process.

If over ⅔ of the quorum's voting power attest "valid", the task is considered approved. If over ⅓ of the quorum's voting power attest "invalid", the task is rejected, and the quorum executes a slashing event to the Performer. The Attesters run the validation logic using a local HTTP request to the AVS WebAPI.

```bash
curl -X POST \
 http://localhost:4002/validate_task \
 -H 'Content-Type: application/json' \
 -d '{
    "proofOfTask": "<transactionHash>",
    "data": "<plugsHash>",
    "taskDefinitionID": "<taskDefinitionId>"
 }'
```

### Aggregator

The Aggregator listens to events from the [Attester](#attester) nodes and monitors the necessary voting power contribution to a certain task. The Aggregator aggregates the signatures of the [Attesters](#attester) into a BLS aggregated signature and submits a transaction to the AttestationCenter smart contract. After successful validation, the [Performer](#performer), [Attesters](#attester), and Aggregator are eligible to claim task rewards.

## Leader Election

## Shared Consensus
