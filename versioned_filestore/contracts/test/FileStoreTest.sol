import "assertions/Asserter.sol";
import "FileStore.sol";
import "utils/Errors.sol";

contract FileStoreTest is Asserter, Errors {

    function testCreateNullFile(){
        FileStore fs = new FileStore();
        var errorCode = fs.createFile(0);
        assertUintsEqual(uint(errorCode), NOT_ALLOWED, "allowed to add null file");
    }

    function testCreateFileLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");

        assertBytes32Equal(fs.getFileListHead(), "witcher3.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "witcher3.exe", "tail is wrong");
        assertUintsEqual(fs.getFileListSize(), 1, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("witcher3.exe"), 0, "previous is wrong");
        assertBytes32Equal(fs.getNextFile("witcher3.exe"), 0, "next is wrong");
    }

    function testCreateTwoFilesLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.createFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 2");
        assertBytes32Equal(fs.getFileListHead(), "skyrim.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "witcher3.exe", "tail is wrong");
        assertUintsEqual(fs.getFileListSize(), 2, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("witcher3.exe"), 0, "previous is wrong for tail");
        assertBytes32Equal(fs.getNextFile("witcher3.exe"), "skyrim.exe", "next is wrong for tail");
        assertBytes32Equal(fs.getPreviousFile("skyrim.exe"), "witcher3.exe", "previous is wrong for head");
        assertBytes32Equal(fs.getNextFile("skyrim.exe"), 0, "next is wrong for head");
    }

    function testCreateThreeFilesLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.createFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 2");
        errorCode = fs.createFile("cs.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 3");

        assertBytes32Equal(fs.getFileListHead(), "cs.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "witcher3.exe", "tail is wrong");

        assertUintsEqual(fs.getFileListSize(), 3, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("witcher3.exe"), 0, "previous is wrong for tail");
        assertBytes32Equal(fs.getNextFile("witcher3.exe"), "skyrim.exe", "next is wrong for tail");
        assertBytes32Equal(fs.getPreviousFile("skyrim.exe"), "witcher3.exe", "previous is wrong for middle");
        assertBytes32Equal(fs.getNextFile("skyrim.exe"), "cs.exe", "next is wrong for middle");
        assertBytes32Equal(fs.getPreviousFile("cs.exe"), "skyrim.exe", "previous is wrong for head");
        assertBytes32Equal(fs.getNextFile("cs.exe"), 0, "next is wrong for head");
    }

    function testRemoveOneOfOneFileLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.removeFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when removing file");

        assertBytes32Equal(fs.getFileListHead(), 0, "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), 0, "tail is wrong");

        assertUintsEqual(fs.getFileListSize(), 0, "size is wrong");
    }

    function testRemoveTailWithTwoLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.createFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 2");
        errorCode = fs.removeFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when removing file");

        assertBytes32Equal(fs.getFileListHead(), "skyrim.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "skyrim.exe", "tail is wrong");
        assertUintsEqual(fs.getFileListSize(), 1, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("skyrim.exe"), 0, "previous is wrong for file");
        assertBytes32Equal(fs.getNextFile("skyrim.exe"), 0, "next is wrong for file");
    }

    function testRemoveHeadWithTwoLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 2");
        errorCode = fs.removeFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when removing file");

        assertBytes32Equal(fs.getFileListHead(), "skyrim.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "skyrim.exe", "tail is wrong");
        assertUintsEqual(fs.getFileListSize(), 1, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("skyrim.exe"), 0, "previous is wrong for file");
        assertBytes32Equal(fs.getNextFile("skyrim.exe"), 0, "next is wrong for file");
    }

    function testRemoveMiddleWithThreeLinks(){
        FileStore fs = new FileStore();

        var errorCode = fs.createFile("witcher3.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file");
        errorCode = fs.createFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 2");
        errorCode = fs.createFile("cs.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when adding file 3");
        errorCode = fs.removeFile("skyrim.exe");
        assertUintsEqual(uint(errorCode), NO_ERROR, "error when removing file");

        assertBytes32Equal(fs.getFileListHead(), "cs.exe", "head is wrong");
        assertBytes32Equal(fs.getFileListTail(), "witcher3.exe", "tail is wrong");
        assertUintsEqual(fs.getFileListSize(), 2, "size is wrong");
        assertBytes32Equal(fs.getPreviousFile("witcher3.exe"), 0, "previous is wrong for tail");
        assertBytes32Equal(fs.getNextFile("witcher3.exe"), "cs.exe", "next is wrong for tail");
        assertBytes32Equal(fs.getPreviousFile("cs.exe"), "witcher3.exe", "previous is wrong for head");
        assertBytes32Equal(fs.getNextFile("cs.exe"), 0, "next is wrong for head");
    }

}