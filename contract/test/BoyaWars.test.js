const BoyaWars = artifacts.require("BoyaWars");

contract("BoyaWars", accounts => {
    it("should deploy", async () => {
        const boyaWars = await BoyaWars.deployed();
        assert(boyaWars.address !== "");
    });

    it("should have an owner", async () => {
        const boyaWars = await BoyaWars.deployed();
        const owner = await boyaWars.owner();
        assert(owner === accounts[0]);
    });

    it("should game is active", async () => {
        const boyaWars = await BoyaWars.deployed();
        await boyaWars.startGame();
        const isGameActive = await boyaWars.gameActive();
        assert(isGameActive === true);
    });

    it("should have a starting bid of 0", async () => {
        const boyaWars = await BoyaWars.deployed();
        const highestBid = await boyaWars.highestBid();
        assert(highestBid.toNumber() === 0);
    });

    it("should end the game", async () => {
        const boyaWars = await BoyaWars.deployed();
        await boyaWars.endGame();
        assert(boyaWars.gameActive() === false);
    });

})