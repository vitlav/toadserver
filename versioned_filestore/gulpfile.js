/*
 *      Build- and test-automation script.
 *
 *      Tasks:
 *
 *      'build-contracts' - build the smart contracts (requires the solidity compiler 'solc')
 *
 *      'build-test-contracts' - build the smart contract unit tests (requires the solidity compiler as well)
 *
 *      'deploy-contracts' - Deploy the system onto a chain.
 */

var gulp = require('gulp');
var contracts = require('./lib/contracts')();
var fs = require('fs-extra');
var app = require('./dapp.json');

require('require-dir')('./gulp-tasks');

// Build the contracts.
gulp.task('build-contracts', ['contracts-build']);

// Build the contract tests.
gulp.task('build-test-contracts', ['contracts-post-build-tests']);


// Deploy contracts task.
gulp.task('deploy-contracts', function(cb){deploy(cb)});

// Default is to build the contracts.
gulp.task('default', ['build-contracts']);

function deploy(callback){
    var abi, compiled;
    try {
        abi = fs.readJsonSync('contracts/build/InterestCalculator.abi');
        compiled = fs.readFileSync('contracts/build/InterestCalculator.binary').toString();
    } catch (error){
        callback(error);
        return;
    }
    var calcFactory = contracts(abi);

    /*
     if(dapp.bank_address){
     var bank = bankFactory.at(dapp.bank_address);
     bank.remove(function(error){
     // Put creation code here instead.
     });
     }
     */

    calcFactory.new({data: compiled}, function(error, contract){
        if(error){
            callback(error);
            return;
        }
        app.calc_address = contract.address;
        var err;
        try {
            fs.writeJsonSync('app.json', app);
        } catch(error){
            err = error;
        }
        callback(err);
    })
}