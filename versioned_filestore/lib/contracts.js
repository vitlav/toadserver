var fs = require('fs-extra');
var erisC = require('eris-contracts');
// All that needs doing atm.
var app = require('../dapp.json');
var erisdbURL = app.server_url;
var privKey = app.priv_key;

module.exports = function(){
    return erisC.contractsDev(erisdbURL, privKey);
};