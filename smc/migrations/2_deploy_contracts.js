const tryC=artifacts.require('./Try.sol');

module.exports = function(deployer) {
    deployer.then(async () => {
        await deployer.deploy(tryC);
    });
};
