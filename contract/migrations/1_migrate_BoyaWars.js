const BoyaWars = artifacts.require("BoyaWars");

module.exports = function (deployer) {
  deployer.deploy(BoyaWars);
}