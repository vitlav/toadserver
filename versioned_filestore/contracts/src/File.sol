import "Committer.sol";
import "Tagger.sol";
import "auth/AdminAuth.sol";

contract File is Committer, Tagger, AdminAuth {

    function File(address _admin) AdminAuth(_admin) {}

    function addTag(bytes32 tag, bytes32 fileHash) returns (uint16 error){
        // Only admin of this contract may write to it.
        if(!isAdmin()){
            return WRITE_ACCESS_DENIED;
        }
        // 'latest' is reserved, and empty tags (or hashes) are not allowed.
        if(tag == "latest" || tag == 0 || fileHash == 0){
            return NOT_ALLOWED;
        }
        // Can't reference a commit that does not exist.
        if(commits[fileHash].fileHash == 0){
            return NOT_ALLOWED;
        }
        error = _addTag(tag, fileHash);
    }

    function removeTag(bytes32 tag) returns (uint16 error) {
        // Only admin of this contract may write to it.
        if(!isAdmin()){
            return WRITE_ACCESS_DENIED;
        }
        // 'latest' is reserved, and empty tags are not allowed.
        if(tag == "latest" || tag == 0){
            return NOT_ALLOWED;
        }
        error = _removeTag(tag);
    }

    function commit(bytes32 fileHash) returns (uint16 error){
        // Only admin of this contract may write to it.
        if(!isAdmin()){
            return WRITE_ACCESS_DENIED;
        }
        // empty hashes are not allowed.
        if(fileHash == 0){
            return NOT_ALLOWED;
        }
        error = _commit(fileHash, msg.sender);
        if(error == NO_ERROR){
            _updateLatest();
        }
    }

    function commitAndTag(bytes32 tag, bytes32 fileHash) returns (uint16 error){
        // Only admin of this contract may write to it.
        if(!isAdmin()){
            return WRITE_ACCESS_DENIED;
        }
        // 'latest' is reserved, and empty tags (or hashes) are not allowed.
        if(tag == "latest" || tag == 0 || fileHash == 0){
            return NOT_ALLOWED;
        }
        error = _commit(fileHash, msg.sender);
        if(error != NO_ERROR){
            return;
        }
        error = _addTag(tag, fileHash);
        if(error != NO_ERROR){
            return;
        }
        _updateLatest();
    }

    function _updateLatest() internal {
        var l = commits_bck.length;
        if(l > 0){
            var latestRef = commits_bck[l - 1];
            tags.tagMap["latest"].ref = latestRef;
        }
    }

}