import { default as fs } from 'fs-extra'
import { exec, execSync } from "child_process";

import { contracts } from '../constants';
import { Contract } from '../types';

const efficientAddressesPath = "create2crunch/efficient_addresses.txt"

const crunchSeconds = 15;
const crunchLeading = 5;
const crunchTotal = 7;

execSync(`rm -rf create2crunch`);
execSync(`sudo apt install build-essential -y; curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y; source "$HOME/.cargo/env"; git clone https://github.com/0age/create2crunch && cd create2crunch; sed -i -e 's/0x4/0x40/g' src/lib.rs`);

const factoryAddress = "0x0000000000ffe8b47b3e2130213b802212439497"
const callerAddress = "0x0000000000000000000000000000000000000000"

const efficientAddressesObject: Record<string, Array<string>> = {}

const mine = async (contract: Contract): Promise<void> => {
    return new Promise((resolve, reject) => {
        const artifactPath = `artifacts/${contract.name}/`
        const contractName = contract.name.replaceAll(".sol", "").replaceAll(".", "");
        const initCodePath = `${artifactPath}${contractName}.initcode.json`
        const initCodeJson = JSON.parse(fs.readFileSync(initCodePath).toString());
        const initCodeHash = initCodeJson["initcodeHash"];

        const interval = setInterval(() => {
            const efficientAddressesExists = fs.existsSync(efficientAddressesPath);

            if (!efficientAddressesExists) {
                console.log("Efficient addresses file does not exist yet.");
                return
            }
    
            // Read and save the efficient addresses file while clearing for the next iteration.
            const efficientAddresses = fs.readFileSync(efficientAddressesPath).toString();
            const running = efficientAddresses == "" ? true : false;

            if (running) {
                console.log(`Waiting another ${crunchSeconds} second period.`);
                return
            }

            process.kill();
            clearInterval(interval);

            efficientAddressesObject[contract.name] = efficientAddresses.split("\n");
            fs.writeFileSync(efficientAddressesPath, "");

            resolve();
        }, crunchSeconds * 1000);

        const process = exec(`cd create2crunch && export FACTORY="${factoryAddress}"; export CALLER="${callerAddress}"; export INIT_CODE_HASH="${initCodeHash}"; export LEADING=${crunchLeading}; export TOTAL=${crunchTotal}; cargo run --release $FACTORY $CALLER $INIT_CODE_HASH 0 $LEADING $TOTAL`, function(error, stdout, stderr) {
            if (error) {
                console.error(`exec error: ${error}`);
                return reject(error);
            }
            console.log(`stdout: ${stdout}`);
            console.error(`stderr: ${stderr}`);
        });
    })
}

async function processContracts() {
    for (const contract of contracts) {
        console.log("✔︎ Generating efficient addresses for contract: ", contract.name);
        await mine(contract);
    }
}

// Call the async function
processContracts().then(() => {
    console.log("✔︎ All contract addresses mined.");
    console.log(efficientAddressesObject);
});
