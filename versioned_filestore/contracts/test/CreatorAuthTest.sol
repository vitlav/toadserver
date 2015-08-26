import "assertions/Asserter.sol";
import "auth/CreatorAuth.sol";

contract CreatorAuthTest is Asserter {

    function testIsAdmin(){
        CreatorAuth ca = new CreatorAuth();
        assertTrue(ca.isAdmin(), "sender is not admin");
    }

    function testIsAdminWithParam(){
        CreatorAuth ca = new CreatorAuth();
        assertTrue(ca.isAdmin(address(this)), "addr is not admin");
    }

    function testIsAdminWithParamFail(){
        CreatorAuth ca = new CreatorAuth();
        assertFalse(ca.isAdmin(address(0x12345)), "wrong addr is admin");
    }

    function testSetAdmin(){
        CreatorAuth ca = new CreatorAuth();
        var res = ca.setAdmin(0x55);
        assertAddressesEqual(ca.admin(), 0x55, "admin was not set");
    }
}