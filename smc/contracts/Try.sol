pragma solidity ^0.4.23;

contract Try {
  uint64 public id;

  event ItemSet(uint64 id);

  function setId(uint64 _id) public {
    id = _id;
    emit ItemSet(id);
  }

  function getId() public view returns(uint64){
    return id;
  }


}
