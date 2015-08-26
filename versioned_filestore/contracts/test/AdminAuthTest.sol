import "assertions/Asserter.sol";
import "auth/AdminAuth.sol";

contract AdminAuthTest is Asserter {

    function testIsAdmin(){
        AdminAuth aa = new AdminAuth(address(this));
        assertTrue(aa.isAdmin(), "sender is not admin");
    }

    function testIsAdminWithParam(){
        AdminAuth aa = new AdminAuth(address(this));
        assertTrue(aa.isAdmin(address(this)), "addr is not admin");
    }

    function testIsAdminWithParamFail(){
        AdminAuth aa = new AdminAuth(address(this));
        assertFalse(aa.isAdmin(address(0x12345)), "wrong addr is admin");
    }

    function testSetAdmin(){
        AdminAuth aa = new AdminAuth(address(this));
        var res = aa.setAdmin(0x55);
        assertAddressesEqual(aa.admin(), 0x55, "admin was not set");
    }
}