import "utils/Errors.sol";

contract Committer is Errors {

    struct Commit {
        bytes32 fileHash;
        uint timeStamp;
        address committer;
    }

    bytes32[] commits_bck;

    mapping(bytes32 => Commit) commits;

    function _commit(bytes32 fileHash, address committer) internal returns (uint16 error) {
        commits[fileHash] = Commit(fileHash, now, committer);
        var index = commits_bck.length++;
        commits_bck[index] = fileHash;
    }

    function getRefByIndex(uint index) returns (bytes32 hash){
        if(commits_bck.length <= index){
            return 0;
        }
        return commits_bck[index];
    }

    function getCommit(bytes32 ref) constant returns (bytes32 fileHash, uint timeStamp, address committer){
        Commit c = commits[ref];
        if(c.fileHash == 0){
            return;
        }
        fileHash = c.fileHash;
        timeStamp = c.timeStamp;
        committer = c.committer;
    }

    function getCommitFileHash(bytes32 ref) constant returns (bytes32 fileHash){
        Commit c = commits[ref];
        return c.fileHash;
    }

    function getCommitTimeStamp(bytes32 ref) constant returns (uint timeStamp){
        Commit c = commits[ref];
        return c.timeStamp;
    }

    function getCommitCommitter(bytes32 ref) constant returns (address committer){
        Commit c = commits[ref];
        return c.committer;
    }

    function getCommitListSize() constant returns (uint size){
        return commits_bck.length;
    }
}