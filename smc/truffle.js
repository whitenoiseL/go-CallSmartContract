const PrivateKeyProvider = require('truffle-privatekey-provider');
const provider = process.env.PRI && process.env.URL && new PrivateKeyProvider(process.env.PRI, process.env.URL);

module.exports = {
    networks: {
        prod: {
            host:"localhost",
            network_id: "*",
            port:8545,
            gas: 4500000, // Gas limit used for deploys
            gasPrice: 10000000000, // 10 gwei
            from:"0x50F8019250Fe0e016140fA231656D2028E161612"
        },
        myprod: {
            provider,
            network_id: "*",
            gas: 4500000, // Gas limit used for deploys
            gasPrice: 10000000000 // 10 gwei
        }
    },
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
    // https://truffle.readthedocs.io/en/beta/advanced/configuration/
    mocha: {
        bail: true
    }
};
