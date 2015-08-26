import "assertions/Asserter.sol";
import "utils/Errors.sol";
import "File.sol";

contract FileTest is Asserter, Errors {

    function testCreateFileAdmin(){
        File f = new File(address(this));
        assertAddressesEqual(f.admin(), address(this), "admin not set");
    }

    // *********************** commits ***********************

    function testCommitGetRefByIndex(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        assertBytes32Equal(f.getRefByIndex(0), 0xdeadbeef, "ref does not match index");
    }

    function testCommitNotAdmin(){
        File f = new File(0xdeadbeef);
        var errorCode = f.commit(0xfeedface);
        assertUintsEqual(uint(errorCode), WRITE_ACCESS_DENIED, "commit allowed for non admin");
    }

    function testCommitEmptyHash(){
        File f = new File(address(this));
        var errorCode = f.commit(0);
        assertUintsEqual(uint(errorCode), NOT_ALLOWED, "error when committing");
    }

    function testCommit(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        var ref = f.getRefByIndex(0);
        assertBytes32Equal(f.getCommitFileHash(ref), 0xdeadbeef, "filehashes not equal");
        assertUintsEqual(f.getCommitTimeStamp(ref), now, "timestamps not equal");
        assertAddressesEqual(f.getCommitCommitter(ref), address(this), "committers not equal");
    }

    function testCommitTwice(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing first");
        errorCode = f.commit(0xfeedface);
        assertUintsEqual(uint(errorCode), 0, "error when committing second");
        assertBytes32Equal(f.getRefByIndex(1), 0xfeedface, "ref does not match index");
    }

    // *********************** tags ***********************

    function testHasLatestTag(){
        File f = new File(address(this));
        assertBytes32Equal(f.getTagListTail(), "latest", "latest tag not added");
    }

    function testLatestUpdated(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        assertBytes32Equal(f.getRef("latest"), 0xdeadbeef, "latest does not match ref");
    }

    function testAddTag(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when addTag");
        assertBytes32Equal(f.getRef("george"), 0xdeadbeef, "tag does not match ref");
    }

    function testAddNullTag(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag(0, 0xdeadbeef);
        assertUintsEqual(uint(errorCode), NOT_ALLOWED, "empty tag allowed");
    }

    function testAddNullRef(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag("george", 0);
        assertUintsEqual(uint(errorCode), NOT_ALLOWED, "empty ref allowed");
    }

    function testTagNonRef(){
        File f = new File(address(this));
        var errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), NOT_ALLOWED, "tagging null-ref allowed");
    }

    function testAddTagNotAdmin(){
        File f = new File(0xfeedface);
        var errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), WRITE_ACCESS_DENIED, "non addmin add tag allowed");
    }

    function testRemoveTagNotAdmin(){
        File f = new File(0xfeedface);
        var errorCode = f.removeTag("george");
        assertUintsEqual(uint(errorCode), WRITE_ACCESS_DENIED, "non admin remove tag allowed");
    }

    function testAddTagLinks(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when addTag");
        assertBytes32Equal(f.getTagListHead(), "george", "head is wrong");
        assertBytes32Equal(f.getTagListTail(), "latest", "tail is wrong");
        assertUintsEqual(f.getTagListSize(), 2, "size is wrong");
        assertBytes32Equal(f.getPreviousTag("george"), "latest", "previous is wrong");
        assertBytes32Equal(f.getNextTag("latest"), "george", "next is wrong");
    }

    function testRemoveTagLinks(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when addTag");
        errorCode = f.removeTag("george");
        assertUintsEqual(uint(errorCode), 0, "error when removeTag");
        assertBytes32Equal(f.getTagListHead(), "latest", "head is wrong");
        assertBytes32Equal(f.getTagListTail(), "latest", "tail is wrong");
        assertUintsEqual(f.getTagListSize(), 1, "size is wrong");
        assertBytes32Equal(f.getNextTag("latest"), 0, "next is wrong");
    }

    function testAddTwoTagsRemoveFirst(){
        File f = new File(address(this));
        var errorCode = f.commit(0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        errorCode = f.addTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when addTag");
        errorCode = f.addTag("fred", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when addTag 2");
        errorCode = f.removeTag("george");
        assertUintsEqual(uint(errorCode), 0, "error when removeTag");
        assertBytes32Equal(f.getTagListHead(), "fred", "head is wrong");
        assertBytes32Equal(f.getTagListTail(), "latest", "tail is wrong");
        assertUintsEqual(f.getTagListSize(), 2, "size is wrong");
        assertBytes32Equal(f.getNextTag("latest"), "fred", "next is wrong");
        assertBytes32Equal(f.getPreviousTag("fred"), "latest", "previous is wrong");
    }

    // *********************** commit and tag ***********************

    function testCommitAndTag(){
        File f = new File(address(this));
        var errorCode = f.commitAndTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), 0, "error when committing");
        assertBytes32Equal(f.getRef("george"), 0xdeadbeef, "tag does not match ref");
    }

    function testCommitAndTagNotAdmin(){
        File f = new File(0xfeedface);
        var errorCode = f.commitAndTag("george", 0xdeadbeef);
        assertUintsEqual(uint(errorCode), WRITE_ACCESS_DENIED, "commitAndTag allowed non-admin");
    }

}