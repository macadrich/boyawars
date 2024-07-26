const fs = require('fs');

const contractName = 'BoyaWars';
const json = require(`../contract/build/contracts/${contractName}.json`);

// Extract ABI
const abi = json.abi;

// Save ABI to a file
fs.writeFileSync(`../contract/build/contracts/${contractName}ABI.json`, JSON.stringify(abi, null, 2));

console.log('ABI extracted and saved to file.');
