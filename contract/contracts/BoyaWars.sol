// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract BoyaWars {
    address public owner;
    address public lastBidder;
    uint public highestBid;
    uint public endTime;
    uint public commissionRate = 5;
    bool public gameActive;
    

    event GameStarted(uint endTime);
    event NewBid(address bidder, uint amount, uint newEndTime);
    event WinnerPaid(address winner, uint amount);

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }

    function startGame() public onlyOwner {
        require(!gameActive, "Game is already active.");
        highestBid = 0;
        lastBidder = address(0);
        endTime = block.timestamp + 2 seconds;
        gameActive = true;
        emit GameStarted(endTime);
    }

    function bid() public payable {
        require(gameActive, "Game is not active.");
        require(msg.value > highestBid, "Your bid is not high enough.");
        require(block.timestamp < endTime, "The bidding time has expired.");

        uint commission = (msg.value * commissionRate)/ 100;
        uint bidAmount = msg.value - commission;
        highestBid = bidAmount;

        if (lastBidder != address(0)) {
            payable(lastBidder).transfer(highestBid);
        }

        lastBidder = msg.sender;
        endTime = block.timestamp + 2 seconds;
        emit NewBid(msg.sender, bidAmount, endTime);
    }

    function endGame() public {
        require(gameActive, "Game is not active.");
        require(block.timestamp >= endTime, "The game is still ongoing.");

        gameActive = false;
        uint payoutAmount = address(this).balance;
        if (lastBidder != address(0)) {
            payable(lastBidder).transfer(payoutAmount);
            emit WinnerPaid(lastBidder, payoutAmount);
        }
    }

    function withdrawFunds() public onlyOwner {
        require(!gameActive, "Cannot withdraw funds while game is active.");
        payable(owner).transfer(address(this).balance);
    }
}
