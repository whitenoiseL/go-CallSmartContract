pragma solidity ^0.4.23;

contract Try {
  uint64 public id;

  function setId(uint64 _id) public {
    id = _id;
  }

  function getId() public view returns(uint64){
    return id;
  }
}
