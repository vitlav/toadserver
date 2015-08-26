// Base contract for authentication. This is used to authenticate contract callers (message senders).
/// @title Auth
/// @author Andreas Olofsson (andreas@erisindustries.com)
contract Auth {
    /// @dev Check if the address 'addr' is admin.
    /// @param addr - The address to check.
    /// @return true if addr is admin, otherwise false.
    function isAdmin(address addr) constant returns (bool);

    /// @dev Check if sender is admin.
    /// @return true if sender is admin, otherwise false.
    function isAdmin() constant returns (bool);

}