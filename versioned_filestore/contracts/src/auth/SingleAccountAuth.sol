import "auth/Auth.sol";
// Base contract for when only a single account is authorized.
/// @title Auth
/// @author Andreas Olofsson (andreas@erisindustries.com)
contract SingleAccountAuth {

    address public admin;

    /// @dev Check if the address 'addr' is admin.
    /// @param addr - The address to check.
    /// @return true if addr is admin, otherwise false.
    function isAdmin(address addr) constant returns (bool) {
        return addr == admin;
    }

    /// @dev Check if sender is admin.
    /// @return true if sender is admin, otherwise false.
    function isAdmin() constant returns (bool){
        return msg.sender == admin;
    }

    /// @dev Set the new admin address to 'addr'. This will fail unless sender is current admin.
    /// @param addr - The address of the new admin.
    /// @return true if admin is changed, otherwise false.
    function setAdmin(address addr) constant returns (bool) {
        if (isAdmin(msg.sender)){
            admin = addr;
            return true;
        }
        return false;
    }

}