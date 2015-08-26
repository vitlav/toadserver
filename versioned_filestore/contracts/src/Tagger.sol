import "utils/Errors.sol";

contract Tagger is Errors {

    struct TagElement {
        bytes32 previous;
        bytes32 next;
        bytes32 ref;
    }

    struct TagList {
        bytes32 head;
        bytes32 tail;
        uint size;
        mapping (bytes32 => TagElement) tagMap;
    }

    TagList tags;

    function Tagger(){
      _addTag("latest", 0);
    }

    function _addTag(bytes32 tag, bytes32 hash) internal returns (uint16 error) {

        // Add
        TagElement e = tags.tagMap[tag];
        e.ref = hash;

        // Link
        if(tags.tail == 0){
            tags.head = tag;
            tags.tail = tag;
        } else {
            e.previous = tags.head;
            tags.tagMap[tags.head].next = tag;
            tags.head = tag;
        }
        tags.size++;
    }

    function _removeTag(bytes32 tag) internal returns (uint16 error){

        TagElement e = tags.tagMap[tag];
        if(e.ref == 0){
            return RESOURCE_NOT_FOUND;
        }

        bytes32 next;
        bytes32 previous;
        TagElement storage ne;
        TagElement storage pe;

        // Note we may never remove the tail and there is always at least 1 element left due to the
        // restrictions of the 'latest' tag.
        if (tags.head == tag){
            previous = e.previous;
            pe = tags.tagMap[previous];
            pe.next = 0;
            tags.head = previous;
        } else {
            next = e.next;
            previous = e.previous;
            ne = tags.tagMap[next];
            pe = tags.tagMap[previous];
            ne.previous = previous;
            pe.next = next;
        }
        delete tags.tagMap[tag];
        tags.size--;
    }

    function getRef(bytes32 tag) constant returns (bytes32 hash){
        return tags.tagMap[tag].ref;
    }

    function getPreviousTag(bytes32 tag) constant returns (bytes32 previousBond){
        return tags.tagMap[tag].previous;
    }

    function getNextTag(bytes32 tag) constant returns (bytes32 nextBond){
        return tags.tagMap[tag].next;
    }

    function getTagListHead() constant returns (bytes32 head){
        return tags.head;
    }

    function getTagListTail() constant returns (bytes32 tail){
        return tags.tail;
    }

    function getTagListSize() constant returns (uint size){
        return tags.size;
    }

    function getTagListElement(bytes32 tag) constant returns (bytes32 previous, bytes32 next, bytes32 ref){
        TagElement e = tags.tagMap[tag];
        previous = e.previous;
        next = e.next;
        ref = e.ref;
    }

}