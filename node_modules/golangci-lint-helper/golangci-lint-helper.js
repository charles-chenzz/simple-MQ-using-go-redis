#!/usr/bin/env node
const { execSync } = require('child_process');
const path = require('path');


// Check if golangci-lint is installed
try {
    execSync('golangci-lint --version', { stdio: 'ignore' });
} catch (error) {
    console.error('golangci-lint-helper: golangci-lint is not installed or not found in PATH');
    process.exit(1);
}


// Get arguments from command line
let args = process.argv.slice(2);

console.log(`golangci-lint-helper: args ${args.join(' ')}`);

// Check if any arguments provided
if (args.length === 0) {
    console.log('golangci-lint-helper: no arguments provided');
    process.exit(1);
}

let fix = "";
if (args[0] === '--fix') {
    args = args.slice(1);
    fix = '--fix';
}

let modules = [];

// Extract unique Go modules from the args which are paths
args.forEach((arg) => {
    let modulePath = path.dirname(arg);

    // Check if modulePath is not in the modules array
    if (!modules.includes(modulePath)) {
        modules.push(modulePath);
    }
});

console.log(`golangci-lint-helper: got modules ${modules.join(' ')}`);

let exitCode = 0;

for (let module of modules) {
    try {
        execSync(`golangci-lint run ${fix} ${module}`, { stdio: 'inherit' });
    } catch (error) {
        console.log(`golangci-lint-helper: error running golangci-lint on module ${module}`);
        exitCode = 1;
    }
}

process.exit(exitCode);
