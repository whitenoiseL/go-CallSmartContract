const tryC=artifacts.require('./Try.sol');

module.exports = async function(deployer) {
   await deployer.deploy(tryC);
};
