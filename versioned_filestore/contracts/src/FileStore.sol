import "File.sol";
import "utils/Errors.sol";

contract FileStore is Errors {

    struct FileElement {
        bytes32 previous;
        bytes32 next;
        address fileAddress;
    }

    struct FileList {
        bytes32 head;
        bytes32 tail;
        uint size;
        mapping (bytes32 => FileElement) fileMap;
    }

    FileList files;

    function createFile(bytes32 file) returns (uint16 error) {
        if(file == 0){
            return NOT_ALLOWED;
        }
        // Add
        FileElement e = files.fileMap[file];

        if(e.fileAddress != 0){
            return RESOURCE_CONFLICT;
        }

        e.fileAddress = address(new File(msg.sender));

        // Link
        if(files.tail == 0){
            files.head = file;
            files.tail = file;
        } else {
            e.previous = files.head;
            files.fileMap[files.head].next = file;
            files.head = file;
        }
        files.size++;
    }

    function removeFile(bytes32 file) returns (uint16 error){

        FileElement e = files.fileMap[file];
        if(e.fileAddress == 0){
            return RESOURCE_NOT_FOUND;
        }

        bytes32 next;
        bytes32 previous;
        FileElement storage ne;
        FileElement storage pe;

        if(files.size == 1){
            files.head = 0;
            files.tail = 0;
        } else if (files.tail == file){
            next = e.next;
            ne = files.fileMap[next];
            ne.previous = 0;
            files.tail = next;
        } else if (files.head == file){
            previous = e.previous;
            pe = files.fileMap[previous];
            pe.next = 0;
            files.head = previous;
        } else {
            next = e.next;
            previous = e.previous;
            ne = files.fileMap[next];
            pe = files.fileMap[previous];
            ne.previous = previous;
            pe.next = next;
        }
        delete files.fileMap[file];
        files.size--;
    }

    function getFileAddress(bytes32 name) constant returns (address fileAddress){
        return files.fileMap[name].fileAddress;
    }

    function getPreviousFile(bytes32 file) constant returns (bytes32 previousFile){
        return files.fileMap[file].previous;
    }

    function getNextFile(bytes32 file) constant returns (bytes32 nextFile){
        return files.fileMap[file].next;
    }

    function getFileListHead() constant returns (bytes32 head){
        return files.head;
    }

    function getFileListTail() constant returns (bytes32 tail){
        return files.tail;
    }

    function getFileListSize() constant returns (uint size){
        return files.size;
    }

    function getFileListElement(bytes32 file) constant returns (bytes32 previous, bytes32 next, address fileAddress){
        FileElement e = files.fileMap[file];
        previous = e.previous;
        next = e.next;
        fileAddress = e.fileAddress;
    }

}